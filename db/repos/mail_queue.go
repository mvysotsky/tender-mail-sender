package repos

import (
	"database/sql"
	"mail-sender/db"
	"mail-sender/db/enum"
	"mail-sender/db/model"
	"mail-sender/db/table"

	"github.com/go-jet/jet/v2/mysql"
)

type MailQueueEntry model.MailQueue
type MailQueueData []MailQueueEntry

type DBMailQueue interface {
	GetPendingMail() (MailQueueData, error)
	SetEntryStatus(id uint32, status model.MailQueueStatus) error
	SetEntryError(id uint32, error string) error
}

type dbMailQueue struct {
	dbHandle *sql.DB
}

func MailQueue() DBMailQueue {
	return dbMailQueue{
		dbHandle: db.Handle,
	}
}

func (r dbMailQueue) GetPendingMail() (mailQueue MailQueueData, err error) {
	err = mysql.
		SELECT(table.MailQueue.AllColumns).
		FROM(table.MailQueue).
		WHERE(table.MailQueue.Status.EQ(enum.MailQueueStatus.Pending)).
		Query(r.dbHandle, &mailQueue)

	return mailQueue, err
}

func (r dbMailQueue) SetEntryStatus(id uint32, status model.MailQueueStatus) error {
	_, err := table.MailQueue.
		UPDATE(table.MailQueue.Status).
		MODEL(model.MailQueue{Status: &status}).
		WHERE(table.MailQueue.ID.EQ(mysql.Int(int64(id)))).
		Exec(db.Handle)

	return err
}

func (r dbMailQueue) SetEntryError(id uint32, error string) error {
	_, err := table.MailQueueErrors.
		INSERT(table.MailQueueErrors.MailQueueID, table.MailQueueErrors.Error).
		VALUES(id, error).
		ON_DUPLICATE_KEY_UPDATE(table.MailQueueErrors.Error.SET(mysql.String(error))).
		Exec(db.Handle)

	return err
}
