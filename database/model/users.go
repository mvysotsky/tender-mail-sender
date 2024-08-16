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

type Users struct {
	ID               uint32 `sql:"primary_key"`
	AccountID        uint32
	RoleID           uint32
	Name             string
	Email            string
	Phone            string
	Login            string
	Password         string
	Enabled          uint8 // 0 - Выключен, 1 - Включен
	DateLastLogin    *time.Time
	DateSendPassword *time.Time
	UserEdit         *uint32
	DateAdd          time.Time
	DateEdit         *time.Time
	SystemRoleID     int8
	ReportsViewer    *bool
	ReportsSettings  *string
}
