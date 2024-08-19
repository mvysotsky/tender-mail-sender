package queries

import (
	"mail-sender/db"
	"mail-sender/db/enum"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

type MailQueueEntry model.MailQueue

type MailQueue []MailQueueEntry

func GetPendingMail() (mailQueue MailQueue, err error) {
	err = mysql.SELECT(table.MailQueue.AllColumns).FROM(
		table.MailQueue,
	).WHERE(
		table.MailQueue.Status.EQ(enum.MailQueueStatus.Pending),
	).Query(db.Handle, &mailQueue)

	return mailQueue, err
}
