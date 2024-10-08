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

var MailQueueErrors = newMailQueueErrorsTable("newtendex", "mail_queue_errors", "")

type mailQueueErrorsTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	MailQueueID mysql.ColumnInteger
	Error       mysql.ColumnString
	CreatedAt   mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MailQueueErrorsTable struct {
	mailQueueErrorsTable

	NEW mailQueueErrorsTable
}

// AS creates new MailQueueErrorsTable with assigned alias
func (a MailQueueErrorsTable) AS(alias string) *MailQueueErrorsTable {
	return newMailQueueErrorsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MailQueueErrorsTable with assigned schema name
func (a MailQueueErrorsTable) FromSchema(schemaName string) *MailQueueErrorsTable {
	return newMailQueueErrorsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MailQueueErrorsTable with assigned table prefix
func (a MailQueueErrorsTable) WithPrefix(prefix string) *MailQueueErrorsTable {
	return newMailQueueErrorsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MailQueueErrorsTable with assigned table suffix
func (a MailQueueErrorsTable) WithSuffix(suffix string) *MailQueueErrorsTable {
	return newMailQueueErrorsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMailQueueErrorsTable(schemaName, tableName, alias string) *MailQueueErrorsTable {
	return &MailQueueErrorsTable{
		mailQueueErrorsTable: newMailQueueErrorsTableImpl(schemaName, tableName, alias),
		NEW:                  newMailQueueErrorsTableImpl("", "new", ""),
	}
}

func newMailQueueErrorsTableImpl(schemaName, tableName, alias string) mailQueueErrorsTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		MailQueueIDColumn = mysql.IntegerColumn("mail_queue_id")
		ErrorColumn       = mysql.StringColumn("error")
		CreatedAtColumn   = mysql.TimestampColumn("created_at")
		allColumns        = mysql.ColumnList{IDColumn, MailQueueIDColumn, ErrorColumn, CreatedAtColumn}
		mutableColumns    = mysql.ColumnList{MailQueueIDColumn, ErrorColumn, CreatedAtColumn}
	)

	return mailQueueErrorsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		MailQueueID: MailQueueIDColumn,
		Error:       ErrorColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
