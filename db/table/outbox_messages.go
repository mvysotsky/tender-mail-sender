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

var OutboxMessages = newOutboxMessagesTable("newtendex", "outbox_messages", "")

type outboxMessagesTable struct {
	mysql.Table

	// Columns
	ID            mysql.ColumnInteger
	MessageBodyID mysql.ColumnInteger
	UserSender    mysql.ColumnInteger
	UserRecipient mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type OutboxMessagesTable struct {
	outboxMessagesTable

	NEW outboxMessagesTable
}

// AS creates new OutboxMessagesTable with assigned alias
func (a OutboxMessagesTable) AS(alias string) *OutboxMessagesTable {
	return newOutboxMessagesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OutboxMessagesTable with assigned schema name
func (a OutboxMessagesTable) FromSchema(schemaName string) *OutboxMessagesTable {
	return newOutboxMessagesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OutboxMessagesTable with assigned table prefix
func (a OutboxMessagesTable) WithPrefix(prefix string) *OutboxMessagesTable {
	return newOutboxMessagesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OutboxMessagesTable with assigned table suffix
func (a OutboxMessagesTable) WithSuffix(suffix string) *OutboxMessagesTable {
	return newOutboxMessagesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOutboxMessagesTable(schemaName, tableName, alias string) *OutboxMessagesTable {
	return &OutboxMessagesTable{
		outboxMessagesTable: newOutboxMessagesTableImpl(schemaName, tableName, alias),
		NEW:                 newOutboxMessagesTableImpl("", "new", ""),
	}
}

func newOutboxMessagesTableImpl(schemaName, tableName, alias string) outboxMessagesTable {
	var (
		IDColumn            = mysql.IntegerColumn("id")
		MessageBodyIDColumn = mysql.IntegerColumn("message_body_id")
		UserSenderColumn    = mysql.IntegerColumn("user_sender")
		UserRecipientColumn = mysql.IntegerColumn("user_recipient")
		allColumns          = mysql.ColumnList{IDColumn, MessageBodyIDColumn, UserSenderColumn, UserRecipientColumn}
		mutableColumns      = mysql.ColumnList{MessageBodyIDColumn, UserSenderColumn, UserRecipientColumn}
	)

	return outboxMessagesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		MessageBodyID: MessageBodyIDColumn,
		UserSender:    UserSenderColumn,
		UserRecipient: UserRecipientColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
