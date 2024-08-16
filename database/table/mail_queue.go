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

var MailQueue = newMailQueueTable("newtendex", "mail_queue", "")

type mailQueueTable struct {
	mysql.Table

	// Columns
	ID         mysql.ColumnInteger
	TenderID   mysql.ColumnInteger
	UserID     mysql.ColumnInteger
	TemplateID mysql.ColumnInteger
	CreatedAt  mysql.ColumnTimestamp
	UpdatedAt  mysql.ColumnTimestamp
	Status     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MailQueueTable struct {
	mailQueueTable

	NEW mailQueueTable
}

// AS creates new MailQueueTable with assigned alias
func (a MailQueueTable) AS(alias string) *MailQueueTable {
	return newMailQueueTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MailQueueTable with assigned schema name
func (a MailQueueTable) FromSchema(schemaName string) *MailQueueTable {
	return newMailQueueTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MailQueueTable with assigned table prefix
func (a MailQueueTable) WithPrefix(prefix string) *MailQueueTable {
	return newMailQueueTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MailQueueTable with assigned table suffix
func (a MailQueueTable) WithSuffix(suffix string) *MailQueueTable {
	return newMailQueueTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMailQueueTable(schemaName, tableName, alias string) *MailQueueTable {
	return &MailQueueTable{
		mailQueueTable: newMailQueueTableImpl(schemaName, tableName, alias),
		NEW:            newMailQueueTableImpl("", "new", ""),
	}
}

func newMailQueueTableImpl(schemaName, tableName, alias string) mailQueueTable {
	var (
		IDColumn         = mysql.IntegerColumn("id")
		TenderIDColumn   = mysql.IntegerColumn("tender_id")
		UserIDColumn     = mysql.IntegerColumn("user_id")
		TemplateIDColumn = mysql.IntegerColumn("template_id")
		CreatedAtColumn  = mysql.TimestampColumn("created_at")
		UpdatedAtColumn  = mysql.TimestampColumn("updated_at")
		StatusColumn     = mysql.StringColumn("status")
		allColumns       = mysql.ColumnList{IDColumn, TenderIDColumn, UserIDColumn, TemplateIDColumn, CreatedAtColumn, UpdatedAtColumn, StatusColumn}
		mutableColumns   = mysql.ColumnList{TenderIDColumn, UserIDColumn, TemplateIDColumn, CreatedAtColumn, UpdatedAtColumn, StatusColumn}
	)

	return mailQueueTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		TenderID:   TenderIDColumn,
		UserID:     UserIDColumn,
		TemplateID: TemplateIDColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		Status:     StatusColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
