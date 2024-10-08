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

var GroupsRecipientsUsers = newGroupsRecipientsUsersTable("newtendex", "groups_recipients_users", "")

type groupsRecipientsUsersTable struct {
	mysql.Table

	// Columns
	ID               mysql.ColumnInteger
	GroupRecipientID mysql.ColumnInteger
	Email            mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type GroupsRecipientsUsersTable struct {
	groupsRecipientsUsersTable

	NEW groupsRecipientsUsersTable
}

// AS creates new GroupsRecipientsUsersTable with assigned alias
func (a GroupsRecipientsUsersTable) AS(alias string) *GroupsRecipientsUsersTable {
	return newGroupsRecipientsUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GroupsRecipientsUsersTable with assigned schema name
func (a GroupsRecipientsUsersTable) FromSchema(schemaName string) *GroupsRecipientsUsersTable {
	return newGroupsRecipientsUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GroupsRecipientsUsersTable with assigned table prefix
func (a GroupsRecipientsUsersTable) WithPrefix(prefix string) *GroupsRecipientsUsersTable {
	return newGroupsRecipientsUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GroupsRecipientsUsersTable with assigned table suffix
func (a GroupsRecipientsUsersTable) WithSuffix(suffix string) *GroupsRecipientsUsersTable {
	return newGroupsRecipientsUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGroupsRecipientsUsersTable(schemaName, tableName, alias string) *GroupsRecipientsUsersTable {
	return &GroupsRecipientsUsersTable{
		groupsRecipientsUsersTable: newGroupsRecipientsUsersTableImpl(schemaName, tableName, alias),
		NEW:                        newGroupsRecipientsUsersTableImpl("", "new", ""),
	}
}

func newGroupsRecipientsUsersTableImpl(schemaName, tableName, alias string) groupsRecipientsUsersTable {
	var (
		IDColumn               = mysql.IntegerColumn("id")
		GroupRecipientIDColumn = mysql.IntegerColumn("group_recipient_id")
		EmailColumn            = mysql.StringColumn("email")
		allColumns             = mysql.ColumnList{IDColumn, GroupRecipientIDColumn, EmailColumn}
		mutableColumns         = mysql.ColumnList{GroupRecipientIDColumn, EmailColumn}
	)

	return groupsRecipientsUsersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		GroupRecipientID: GroupRecipientIDColumn,
		Email:            EmailColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
