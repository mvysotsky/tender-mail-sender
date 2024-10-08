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

var Currencies = newCurrenciesTable("newtendex", "currencies", "")

type currenciesTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	SubserviceID mysql.ColumnInteger
	Name         mysql.ColumnString
	Value        mysql.ColumnString
	UserAdd      mysql.ColumnInteger
	UserEdit     mysql.ColumnInteger
	DateAdd      mysql.ColumnTimestamp
	DateEdit     mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type CurrenciesTable struct {
	currenciesTable

	NEW currenciesTable
}

// AS creates new CurrenciesTable with assigned alias
func (a CurrenciesTable) AS(alias string) *CurrenciesTable {
	return newCurrenciesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CurrenciesTable with assigned schema name
func (a CurrenciesTable) FromSchema(schemaName string) *CurrenciesTable {
	return newCurrenciesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CurrenciesTable with assigned table prefix
func (a CurrenciesTable) WithPrefix(prefix string) *CurrenciesTable {
	return newCurrenciesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CurrenciesTable with assigned table suffix
func (a CurrenciesTable) WithSuffix(suffix string) *CurrenciesTable {
	return newCurrenciesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCurrenciesTable(schemaName, tableName, alias string) *CurrenciesTable {
	return &CurrenciesTable{
		currenciesTable: newCurrenciesTableImpl(schemaName, tableName, alias),
		NEW:             newCurrenciesTableImpl("", "new", ""),
	}
}

func newCurrenciesTableImpl(schemaName, tableName, alias string) currenciesTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		SubserviceIDColumn = mysql.IntegerColumn("subservice_id")
		NameColumn         = mysql.StringColumn("name")
		ValueColumn        = mysql.StringColumn("value")
		UserAddColumn      = mysql.IntegerColumn("user_add")
		UserEditColumn     = mysql.IntegerColumn("user_edit")
		DateAddColumn      = mysql.TimestampColumn("date_add")
		DateEditColumn     = mysql.TimestampColumn("date_edit")
		allColumns         = mysql.ColumnList{IDColumn, SubserviceIDColumn, NameColumn, ValueColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn}
		mutableColumns     = mysql.ColumnList{SubserviceIDColumn, NameColumn, ValueColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn}
	)

	return currenciesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		SubserviceID: SubserviceIDColumn,
		Name:         NameColumn,
		Value:        ValueColumn,
		UserAdd:      UserAddColumn,
		UserEdit:     UserEditColumn,
		DateAdd:      DateAddColumn,
		DateEdit:     DateEditColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
