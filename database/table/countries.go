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

var Countries = newCountriesTable("newtendex", "countries", "")

type countriesTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	SubserviceID mysql.ColumnInteger
	Name         mysql.ColumnString
	UserAdd      mysql.ColumnInteger
	UserEdit     mysql.ColumnInteger
	DateAdd      mysql.ColumnTimestamp
	DateEdit     mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type CountriesTable struct {
	countriesTable

	NEW countriesTable
}

// AS creates new CountriesTable with assigned alias
func (a CountriesTable) AS(alias string) *CountriesTable {
	return newCountriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CountriesTable with assigned schema name
func (a CountriesTable) FromSchema(schemaName string) *CountriesTable {
	return newCountriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CountriesTable with assigned table prefix
func (a CountriesTable) WithPrefix(prefix string) *CountriesTable {
	return newCountriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CountriesTable with assigned table suffix
func (a CountriesTable) WithSuffix(suffix string) *CountriesTable {
	return newCountriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCountriesTable(schemaName, tableName, alias string) *CountriesTable {
	return &CountriesTable{
		countriesTable: newCountriesTableImpl(schemaName, tableName, alias),
		NEW:            newCountriesTableImpl("", "new", ""),
	}
}

func newCountriesTableImpl(schemaName, tableName, alias string) countriesTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		SubserviceIDColumn = mysql.IntegerColumn("subservice_id")
		NameColumn         = mysql.StringColumn("name")
		UserAddColumn      = mysql.IntegerColumn("user_add")
		UserEditColumn     = mysql.IntegerColumn("user_edit")
		DateAddColumn      = mysql.TimestampColumn("date_add")
		DateEditColumn     = mysql.TimestampColumn("date_edit")
		allColumns         = mysql.ColumnList{IDColumn, SubserviceIDColumn, NameColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn}
		mutableColumns     = mysql.ColumnList{SubserviceIDColumn, NameColumn, UserAddColumn, UserEditColumn, DateAddColumn, DateEditColumn}
	)

	return countriesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		SubserviceID: SubserviceIDColumn,
		Name:         NameColumn,
		UserAdd:      UserAddColumn,
		UserEdit:     UserEditColumn,
		DateAdd:      DateAddColumn,
		DateEdit:     DateEditColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
