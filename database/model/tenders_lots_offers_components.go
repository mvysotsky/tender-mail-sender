//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type TendersLotsOffersComponents struct {
	ID                          uint32 `sql:"primary_key"`
	TenderLotOfferID            uint32
	TenderLotFormulaParameterID uint32
	Value                       float64
}
