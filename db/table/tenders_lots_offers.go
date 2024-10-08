//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var TendersLotsOffers = newTendersLotsOffersTable("newtendex", "tenders_lots_offers", "")

type tendersLotsOffersTable struct {
	mysql.Table

	// Columns
	ID                  mysql.ColumnInteger
	TenderLotID         mysql.ColumnInteger
	Offer               mysql.ColumnFloat
	Description         mysql.ColumnString
	DateAdd             mysql.ColumnTimestamp
	DateWin             mysql.ColumnTimestamp
	UserAdd             mysql.ColumnInteger
	IsCanceled          mysql.ColumnInteger // 0 - Не отменена, 1 - Отменена
	IsWinner            mysql.ColumnInteger // 0 - Не победитель, 1 - Победитель
	NotBid              mysql.ColumnInteger
	InternalCause       mysql.ColumnString
	ExternalCause       mysql.ColumnString
	WinnerInternalCause mysql.ColumnString
	WinnerExternalCause mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TendersLotsOffersTable struct {
	tendersLotsOffersTable

	NEW tendersLotsOffersTable
}

// AS creates new TendersLotsOffersTable with assigned alias
func (a TendersLotsOffersTable) AS(alias string) *TendersLotsOffersTable {
	return newTendersLotsOffersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TendersLotsOffersTable with assigned schema name
func (a TendersLotsOffersTable) FromSchema(schemaName string) *TendersLotsOffersTable {
	return newTendersLotsOffersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TendersLotsOffersTable with assigned table prefix
func (a TendersLotsOffersTable) WithPrefix(prefix string) *TendersLotsOffersTable {
	return newTendersLotsOffersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TendersLotsOffersTable with assigned table suffix
func (a TendersLotsOffersTable) WithSuffix(suffix string) *TendersLotsOffersTable {
	return newTendersLotsOffersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTendersLotsOffersTable(schemaName, tableName, alias string) *TendersLotsOffersTable {
	return &TendersLotsOffersTable{
		tendersLotsOffersTable: newTendersLotsOffersTableImpl(schemaName, tableName, alias),
		NEW:                    newTendersLotsOffersTableImpl("", "new", ""),
	}
}

func newTendersLotsOffersTableImpl(schemaName, tableName, alias string) tendersLotsOffersTable {
	var (
		IDColumn                  = mysql.IntegerColumn("id")
		TenderLotIDColumn         = mysql.IntegerColumn("tender_lot_id")
		OfferColumn               = mysql.FloatColumn("offer")
		DescriptionColumn         = mysql.StringColumn("description")
		DateAddColumn             = mysql.TimestampColumn("date_add")
		DateWinColumn             = mysql.TimestampColumn("date_win")
		UserAddColumn             = mysql.IntegerColumn("user_add")
		IsCanceledColumn          = mysql.IntegerColumn("is_canceled")
		IsWinnerColumn            = mysql.IntegerColumn("is_winner")
		NotBidColumn              = mysql.IntegerColumn("notBid")
		InternalCauseColumn       = mysql.StringColumn("internal_cause")
		ExternalCauseColumn       = mysql.StringColumn("external_cause")
		WinnerInternalCauseColumn = mysql.StringColumn("winner_internal_cause")
		WinnerExternalCauseColumn = mysql.StringColumn("winner_external_cause")
		allColumns                = mysql.ColumnList{IDColumn, TenderLotIDColumn, OfferColumn, DescriptionColumn, DateAddColumn, DateWinColumn, UserAddColumn, IsCanceledColumn, IsWinnerColumn, NotBidColumn, InternalCauseColumn, ExternalCauseColumn, WinnerInternalCauseColumn, WinnerExternalCauseColumn}
		mutableColumns            = mysql.ColumnList{TenderLotIDColumn, OfferColumn, DescriptionColumn, DateAddColumn, DateWinColumn, UserAddColumn, IsCanceledColumn, IsWinnerColumn, NotBidColumn, InternalCauseColumn, ExternalCauseColumn, WinnerInternalCauseColumn, WinnerExternalCauseColumn}
	)

	return tendersLotsOffersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                  IDColumn,
		TenderLotID:         TenderLotIDColumn,
		Offer:               OfferColumn,
		Description:         DescriptionColumn,
		DateAdd:             DateAddColumn,
		DateWin:             DateWinColumn,
		UserAdd:             UserAddColumn,
		IsCanceled:          IsCanceledColumn,
		IsWinner:            IsWinnerColumn,
		NotBid:              NotBidColumn,
		InternalCause:       InternalCauseColumn,
		ExternalCause:       ExternalCauseColumn,
		WinnerInternalCause: WinnerInternalCauseColumn,
		WinnerExternalCause: WinnerExternalCauseColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
