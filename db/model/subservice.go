//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Subservice struct {
	ID           uint32 `sql:"primary_key"`
	LanguageID   uint32
	CurrenciesID *uint32
	IsDefault    uint8
	Name         string
	URL          string
	TimeZone     string
	Requisites   *string
}
