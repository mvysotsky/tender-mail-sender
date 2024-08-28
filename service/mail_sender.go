package service

import (
	"log"
	"mail-sender/db/model"
	"mail-sender/db/repos"
)

type MailSender interface {
	SendEmailNotifications() error
}

type mailSender struct {
	mailQueue repos.DBMailQueue
	templates repos.DBTemplates
	tenders   repos.DBTenders
	users     repos.DBUsers
}

func NewMailSender() MailSender {
	return mailSender{
		mailQueue: repos.MailQueue(),
		templates: repos.Templates(),
		tenders:   repos.Tenders(),
		users:     repos.Users(),
	}
}

func (s mailSender) SendEmailNotifications() error {
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

func (s mailSender) send(entry repos.MailQueueEntry) error {
	template, err := s.templates.GetTemplateByID(entry.TemplateID)
	if err != nil {
		return err
	}

	tender, err := s.tenders.GetTender(entry.TenderID)
	if err != nil {
		return err
	}

	user, err := s.users.GetUser(entry.UserID)
	if err != nil {
		return err
	}

	// Send the email
	return nil
}
