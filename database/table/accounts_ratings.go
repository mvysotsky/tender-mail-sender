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

var AccountsRatings = newAccountsRatingsTable("newtendex", "accounts_ratings", "")

type accountsRatingsTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnInteger
	AccountID mysql.ColumnInteger
	Value     mysql.ColumnInteger
	Cause     mysql.ColumnString
	DateAdd   mysql.ColumnTimestamp
	UserAdd   mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AccountsRatingsTable struct {
	accountsRatingsTable

	NEW accountsRatingsTable
}

// AS creates new AccountsRatingsTable with assigned alias
func (a AccountsRatingsTable) AS(alias string) *AccountsRatingsTable {
	return newAccountsRatingsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountsRatingsTable with assigned schema name
func (a AccountsRatingsTable) FromSchema(schemaName string) *AccountsRatingsTable {
	return newAccountsRatingsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountsRatingsTable with assigned table prefix
func (a AccountsRatingsTable) WithPrefix(prefix string) *AccountsRatingsTable {
	return newAccountsRatingsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountsRatingsTable with assigned table suffix
func (a AccountsRatingsTable) WithSuffix(suffix string) *AccountsRatingsTable {
	return newAccountsRatingsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountsRatingsTable(schemaName, tableName, alias string) *AccountsRatingsTable {
	return &AccountsRatingsTable{
		accountsRatingsTable: newAccountsRatingsTableImpl(schemaName, tableName, alias),
		NEW:                  newAccountsRatingsTableImpl("", "new", ""),
	}
}

func newAccountsRatingsTableImpl(schemaName, tableName, alias string) accountsRatingsTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		AccountIDColumn = mysql.IntegerColumn("account_id")
		ValueColumn     = mysql.IntegerColumn("value")
		CauseColumn     = mysql.StringColumn("cause")
		DateAddColumn   = mysql.TimestampColumn("date_add")
		UserAddColumn   = mysql.IntegerColumn("user_add")
		allColumns      = mysql.ColumnList{IDColumn, AccountIDColumn, ValueColumn, CauseColumn, DateAddColumn, UserAddColumn}
		mutableColumns  = mysql.ColumnList{AccountIDColumn, ValueColumn, CauseColumn, DateAddColumn, UserAddColumn}
	)

	return accountsRatingsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		AccountID: AccountIDColumn,
		Value:     ValueColumn,
		Cause:     CauseColumn,
		DateAdd:   DateAddColumn,
		UserAdd:   UserAddColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
