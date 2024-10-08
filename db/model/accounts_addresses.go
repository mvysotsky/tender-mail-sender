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

type AccountsAddresses struct {
	ID                uint32 `sql:"primary_key"`
	Type              uint8  // 1 - Юр, 2 - Физ. 3 - Склад
	AccountID         uint32
	CountryID         *uint32
	CountryRegionID   uint32
	CountryLocationID uint32
	Street            *string
	House             *string
	RoomApartment     *string
	Address           *string
	UserAdd           *uint32
	UserEdit          *uint32
	DateAdd           time.Time
	DateEdit          *time.Time
}
