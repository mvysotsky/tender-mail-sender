//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type TendersDocuments struct {
	ID         uint32 `sql:"primary_key"`
	TenderID   uint32
	DocumentID uint32
}
