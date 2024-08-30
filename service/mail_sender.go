package service

import (
	"fmt"
	"log"
	"mail-sender/db/model"
	"mail-sender/db/repos"
	"mail-sender/tools"
	"strconv"

	"github.com/cbroglie/mustache"
)

type MailSender interface {
	SendEmailNotifications() error
}

type mailSender struct {
	mailQueue       repos.DBMailQueue
	templates       repos.DBTemplates
	tenders         repos.DBTenders
	users           repos.DBUsers
	senderName      string
	subserviceAlias string
	domain          string
}

func NewMailSender() MailSender {
	return &mailSender{
		mailQueue:       repos.MailQueue(),
		templates:       repos.Templates(),
		tenders:         repos.Tenders(),
		users:           repos.Users(),
		senderName:      tools.GetEnv("OWNER_NAME", "ТОВ АТБ-МАРКЕТ"),
		subserviceAlias: tools.GetEnv("SUBSERVICE_ALIAS", "ua"),
		domain:          tools.GetEnv("DOMAIN", ""),
	}
}

func (s *mailSender) SendEmailNotifications() error {
	queue, err := s.mailQueue.GetPendingMail()
	if err != nil {
		return err
	}

	for _, entry := range queue {
		err = s.send(entry)

		if err != nil {
			if err = s.mailQueue.SetEntryError(entry.ID, err.Error()); err != nil {
				log.Printf("Error setting entry error: %v", err)
				continue
			}

			if err = s.mailQueue.SetEntryStatus(entry.ID, model.MailQueueStatus_Error); err != nil {
				log.Printf("Error setting entry status: %v", err)
			}

			continue
		}

		if err = s.mailQueue.SetEntryStatus(entry.ID, model.MailQueueStatus_Sent); err != nil {
			log.Printf("Error setting entry status: %v", err)
		}
	}

	return nil
}

func (s *mailSender) send(entry repos.MailQueueEntry) (err error) {
	var (
		template string
		tender   repos.Tender
		user     repos.User
		message  string
	)

	if template, err = s.templates.GetTemplateByID(entry.TemplateID); err != nil {
		return err
	}

	if tender, err = s.tenders.GetTender(entry.TenderID); err != nil {
		return err
	}

	if user, err = s.users.GetUser(entry.UserID); err != nil {
		return err
	}

	tenderURL := fmt.Sprintf("%s%s/tenders/open_tender/%d", s.domain, s.subserviceAlias, entry.TenderID)

	if message, err = mustache.Render(template, map[string]string{
		"user_fio":           user.Users.Name,
		"account_name":       user.Accounts.Name,
		"account_owner_name": s.senderName,
		"category":           tender.Categories.Name,
		"tender_id":          strconv.Itoa(int(entry.TenderID)),
		"date_start":         tender.DateStart.Format("02.01.2006 15:04"),
		"date_end":           tender.DateEnd.Format("02.01.2006 15:04"),
		"tender_link":        tenderURL,
	}); err != nil {
		return err
	}

	// Send the email
	return nil
}
