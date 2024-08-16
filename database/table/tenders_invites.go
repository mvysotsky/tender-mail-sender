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

var TendersInvites = newTendersInvitesTable("newtendex", "tenders_invites", "")

type tendersInvitesTable struct {
	mysql.Table

	// Columns
	ID              mysql.ColumnInteger
	TenderID        mysql.ColumnInteger
	InviteAccountID mysql.ColumnInteger
	Status          mysql.ColumnInteger // 1 - новая, 2 - принята, 3 - отклонена
	UserAdd         mysql.ColumnInteger
	UserEdit        mysql.ColumnInteger
	DateAdd         mysql.ColumnTimestamp
	DateEdit        mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TendersInvitesTable struct {
	tendersInvitesTable

	NEW tendersInvitesTable
}

// AS creates new TendersInvitesTable with assigned alias
func (a TendersInvitesTable) AS(alias string) *TendersInvitesTable {
	return newTendersInvitesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TendersInvitesTable with assigned schema name
func (a TendersInvitesTable) FromSchema(schemaName string) *TendersInvitesTable {
	return newTendersInvitesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TendersInvitesTable with assigned table prefix
func (a TendersInvitesTable) WithPrefix(prefix string) *TendersInvitesTable {
	return newTendersInvitesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TendersInvitesTable with assigned table suffix
func (a TendersInvitesTable) WithSuffix(suffix string) *TendersInvitesTable {
	return newTendersInvitesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTendersInvitesTable(schemaName, tableName, alias string) *TendersInvitesTable {
	return &TendersInvitesTable{
		tendersInvitesTable: newTendersInvitesTableImpl(schemaName, tableName, alias),
		NEW:                 newTendersInvitesTableImpl("", "new", ""),
	}
}

func newTendersInvitesTableImpl(schemaName, tableName, alias string) tendersInvitesTable {
	var (
		IDColumn              = mysql.IntegerColumn("id")
		TenderIDColumn        = mysql.IntegerColumn("tender_id")
		InviteAccountIDColumn = mysql.IntegerColumn("invite_account_id")
		StatusColumn          = mysql.IntegerColumn("status")
		UserAddColumn         = mysql.IntegerColumn("user_add")
		UserEditColumn        = mysql.IntegerColumn("user_edit")
		DateAddColumn         = mysql.TimestampColumn("date_add")
		DateEditColumn        = mysql.TimestampColumn("date_edit")
		allColumns            = mysql.ColumnList{IDColumn, TenderIDColumn, InviteAccountIDColumn, StatusColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn}
		mutableColumns        = mysql.ColumnList{TenderIDColumn, InviteAccountIDColumn, StatusColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn}
	)

	return tendersInvitesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		TenderID:        TenderIDColumn,
		InviteAccountID: InviteAccountIDColumn,
		Status:          StatusColumn,
		UserAdd:         UserAddColumn,
		UserEdit:        UserEditColumn,
		DateAdd:         DateAddColumn,
		DateEdit:        DateEditColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
