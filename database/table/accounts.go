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

var Accounts = newAccountsTable("newtendex", "accounts", "")

type accountsTable struct {
	mysql.Table

	// Columns
	ID                 mysql.ColumnInteger
	SubserviceID       mysql.ColumnInteger
	Name               mysql.ColumnString
	Img                mysql.ColumnString
	Description        mysql.ColumnString
	Status             mysql.ColumnInteger // 1 - Активный, 2 - Неактивный
	IsActiveContractor mysql.ColumnInteger // 1- Активный, 0 - Неактивный
	DateAdd            mysql.ColumnTimestamp
	DateEdit           mysql.ColumnTimestamp
	AllowIps           mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AccountsTable struct {
	accountsTable

	NEW accountsTable
}

// AS creates new AccountsTable with assigned alias
func (a AccountsTable) AS(alias string) *AccountsTable {
	return newAccountsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountsTable with assigned schema name
func (a AccountsTable) FromSchema(schemaName string) *AccountsTable {
	return newAccountsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountsTable with assigned table prefix
func (a AccountsTable) WithPrefix(prefix string) *AccountsTable {
	return newAccountsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountsTable with assigned table suffix
func (a AccountsTable) WithSuffix(suffix string) *AccountsTable {
	return newAccountsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountsTable(schemaName, tableName, alias string) *AccountsTable {
	return &AccountsTable{
		accountsTable: newAccountsTableImpl(schemaName, tableName, alias),
		NEW:           newAccountsTableImpl("", "new", ""),
	}
}

func newAccountsTableImpl(schemaName, tableName, alias string) accountsTable {
	var (
		IDColumn                 = mysql.IntegerColumn("id")
		SubserviceIDColumn       = mysql.IntegerColumn("subservice_id")
		NameColumn               = mysql.StringColumn("name")
		ImgColumn                = mysql.StringColumn("img")
		DescriptionColumn        = mysql.StringColumn("description")
		StatusColumn             = mysql.IntegerColumn("status")
		IsActiveContractorColumn = mysql.IntegerColumn("is_active_contractor")
		DateAddColumn            = mysql.TimestampColumn("date_add")
		DateEditColumn           = mysql.TimestampColumn("date_edit")
		AllowIpsColumn           = mysql.StringColumn("allow_ips")
		allColumns               = mysql.ColumnList{IDColumn, SubserviceIDColumn, NameColumn, ImgColumn, DescriptionColumn, StatusColumn, IsActiveContractorColumn, DateAddColumn, DateEditColumn, AllowIpsColumn}
		mutableColumns           = mysql.ColumnList{SubserviceIDColumn, NameColumn, ImgColumn, DescriptionColumn, StatusColumn, IsActiveContractorColumn, DateAddColumn, DateEditColumn, AllowIpsColumn}
	)

	return accountsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                 IDColumn,
		SubserviceID:       SubserviceIDColumn,
		Name:               NameColumn,
		Img:                ImgColumn,
		Description:        DescriptionColumn,
		Status:             StatusColumn,
		IsActiveContractor: IsActiveContractorColumn,
		DateAdd:            DateAddColumn,
		DateEdit:           DateEditColumn,
		AllowIps:           AllowIpsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
