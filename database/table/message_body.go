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

var MessageBody = newMessageBodyTable("newtendex", "message_body", "")

type messageBodyTable struct {
	mysql.Table

	// Columns
	ID       mysql.ColumnInteger
	TenderID mysql.ColumnInteger
	Subject  mysql.ColumnString
	Message  mysql.ColumnString
	DateSent mysql.ColumnTimestamp
	IsSystem mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MessageBodyTable struct {
	messageBodyTable

	NEW messageBodyTable
}

// AS creates new MessageBodyTable with assigned alias
func (a MessageBodyTable) AS(alias string) *MessageBodyTable {
	return newMessageBodyTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MessageBodyTable with assigned schema name
func (a MessageBodyTable) FromSchema(schemaName string) *MessageBodyTable {
	return newMessageBodyTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MessageBodyTable with assigned table prefix
func (a MessageBodyTable) WithPrefix(prefix string) *MessageBodyTable {
	return newMessageBodyTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MessageBodyTable with assigned table suffix
func (a MessageBodyTable) WithSuffix(suffix string) *MessageBodyTable {
	return newMessageBodyTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMessageBodyTable(schemaName, tableName, alias string) *MessageBodyTable {
	return &MessageBodyTable{
		messageBodyTable: newMessageBodyTableImpl(schemaName, tableName, alias),
		NEW:              newMessageBodyTableImpl("", "new", ""),
	}
}

func newMessageBodyTableImpl(schemaName, tableName, alias string) messageBodyTable {
	var (
		IDColumn       = mysql.IntegerColumn("id")
		TenderIDColumn = mysql.IntegerColumn("tender_id")
		SubjectColumn  = mysql.StringColumn("subject")
		MessageColumn  = mysql.StringColumn("message")
		DateSentColumn = mysql.TimestampColumn("date_sent")
		IsSystemColumn = mysql.IntegerColumn("is_system")
		allColumns     = mysql.ColumnList{IDColumn, TenderIDColumn, SubjectColumn, MessageColumn, DateSentColumn, IsSystemColumn}
		mutableColumns = mysql.ColumnList{TenderIDColumn, SubjectColumn, MessageColumn, DateSentColumn, IsSystemColumn}
	)

	return messageBodyTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		TenderID: TenderIDColumn,
		Subject:  SubjectColumn,
		Message:  MessageColumn,
		DateSent: DateSentColumn,
		IsSystem: IsSystemColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
