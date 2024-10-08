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

var Mailing = newMailingTable("newtendex", "mailing", "")

type mailingTable struct {
	mysql.Table

	// Columns
	ID               mysql.ColumnInteger
	SubserviceID     mysql.ColumnInteger
	GroupRecipientID mysql.ColumnInteger
	Name             mysql.ColumnInteger
	Subject          mysql.ColumnString
	Content          mysql.ColumnString
	Status           mysql.ColumnInteger // 1 - Новая, 2- В процессе, 3- Отправлена
	UserAdd          mysql.ColumnInteger
	UserEdit         mysql.ColumnInteger
	DateAdd          mysql.ColumnTimestamp
	DateEdit         mysql.ColumnTimestamp
	SenderName       mysql.ColumnString
	SenderEmail      mysql.ColumnString
	DateSent         mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MailingTable struct {
	mailingTable

	NEW mailingTable
}

// AS creates new MailingTable with assigned alias
func (a MailingTable) AS(alias string) *MailingTable {
	return newMailingTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MailingTable with assigned schema name
func (a MailingTable) FromSchema(schemaName string) *MailingTable {
	return newMailingTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MailingTable with assigned table prefix
func (a MailingTable) WithPrefix(prefix string) *MailingTable {
	return newMailingTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MailingTable with assigned table suffix
func (a MailingTable) WithSuffix(suffix string) *MailingTable {
	return newMailingTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMailingTable(schemaName, tableName, alias string) *MailingTable {
	return &MailingTable{
		mailingTable: newMailingTableImpl(schemaName, tableName, alias),
		NEW:          newMailingTableImpl("", "new", ""),
	}
}

func newMailingTableImpl(schemaName, tableName, alias string) mailingTable {
	var (
		IDColumn               = mysql.IntegerColumn("id")
		SubserviceIDColumn     = mysql.IntegerColumn("subservice_id")
		GroupRecipientIDColumn = mysql.IntegerColumn("group_recipient_id")
		NameColumn             = mysql.IntegerColumn("name")
		SubjectColumn          = mysql.StringColumn("subject")
		ContentColumn          = mysql.StringColumn("content")
		StatusColumn           = mysql.IntegerColumn("status")
		UserAddColumn          = mysql.IntegerColumn("user_add")
		UserEditColumn         = mysql.IntegerColumn("user_edit")
		DateAddColumn          = mysql.TimestampColumn("date_add")
		DateEditColumn         = mysql.TimestampColumn("date_edit")
		SenderNameColumn       = mysql.StringColumn("sender_name")
		SenderEmailColumn      = mysql.StringColumn("sender_email")
		DateSentColumn         = mysql.TimestampColumn("date_sent")
		allColumns             = mysql.ColumnList{IDColumn, SubserviceIDColumn, GroupRecipientIDColumn, NameColumn, SubjectColumn, ContentColumn, StatusColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn, SenderNameColumn, SenderEmailColumn, DateSentColumn}
		mutableColumns         = mysql.ColumnList{SubserviceIDColumn, GroupRecipientIDColumn, NameColumn, SubjectColumn, ContentColumn, StatusColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn, SenderNameColumn, SenderEmailColumn, DateSentColumn}
	)

	return mailingTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		SubserviceID:     SubserviceIDColumn,
		GroupRecipientID: GroupRecipientIDColumn,
		Name:             NameColumn,
		Subject:          SubjectColumn,
		Content:          ContentColumn,
		Status:           StatusColumn,
		UserAdd:          UserAddColumn,
		UserEdit:         UserEditColumn,
		DateAdd:          DateAddColumn,
		DateEdit:         DateEditColumn,
		SenderName:       SenderNameColumn,
		SenderEmail:      SenderEmailColumn,
		DateSent:         DateSentColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
