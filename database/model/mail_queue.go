//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type MailQueue struct {
	ID         uint32 `sql:"primary_key"`
	TenderID   uint32
	UserID     uint32
	TemplateID uint32
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     *MailQueueStatus
}
