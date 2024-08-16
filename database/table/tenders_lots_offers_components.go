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

var TendersLotsOffersComponents = newTendersLotsOffersComponentsTable("newtendex", "tenders_lots_offers_components", "")

type tendersLotsOffersComponentsTable struct {
	mysql.Table

	// Columns
	ID                          mysql.ColumnInteger
	TenderLotOfferID            mysql.ColumnInteger
	TenderLotFormulaParameterID mysql.ColumnInteger
	Value                       mysql.ColumnFloat

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TendersLotsOffersComponentsTable struct {
	tendersLotsOffersComponentsTable

	NEW tendersLotsOffersComponentsTable
}

// AS creates new TendersLotsOffersComponentsTable with assigned alias
func (a TendersLotsOffersComponentsTable) AS(alias string) *TendersLotsOffersComponentsTable {
	return newTendersLotsOffersComponentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TendersLotsOffersComponentsTable with assigned schema name
func (a TendersLotsOffersComponentsTable) FromSchema(schemaName string) *TendersLotsOffersComponentsTable {
	return newTendersLotsOffersComponentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TendersLotsOffersComponentsTable with assigned table prefix
func (a TendersLotsOffersComponentsTable) WithPrefix(prefix string) *TendersLotsOffersComponentsTable {
	return newTendersLotsOffersComponentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TendersLotsOffersComponentsTable with assigned table suffix
func (a TendersLotsOffersComponentsTable) WithSuffix(suffix string) *TendersLotsOffersComponentsTable {
	return newTendersLotsOffersComponentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTendersLotsOffersComponentsTable(schemaName, tableName, alias string) *TendersLotsOffersComponentsTable {
	return &TendersLotsOffersComponentsTable{
		tendersLotsOffersComponentsTable: newTendersLotsOffersComponentsTableImpl(schemaName, tableName, alias),
		NEW:                              newTendersLotsOffersComponentsTableImpl("", "new", ""),
	}
}

func newTendersLotsOffersComponentsTableImpl(schemaName, tableName, alias string) tendersLotsOffersComponentsTable {
	var (
		IDColumn                          = mysql.IntegerColumn("id")
		TenderLotOfferIDColumn            = mysql.IntegerColumn("tender_lot_offer_id")
		TenderLotFormulaParameterIDColumn = mysql.IntegerColumn("tender_lot_formula_parameter_id")
		ValueColumn                       = mysql.FloatColumn("value")
		allColumns                        = mysql.ColumnList{IDColumn, TenderLotOfferIDColumn, TenderLotFormulaParameterIDColumn, ValueColumn}
		mutableColumns                    = mysql.ColumnList{TenderLotOfferIDColumn, TenderLotFormulaParameterIDColumn, ValueColumn}
	)

	return tendersLotsOffersComponentsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                          IDColumn,
		TenderLotOfferID:            TenderLotOfferIDColumn,
		TenderLotFormulaParameterID: TenderLotFormulaParameterIDColumn,
		Value:                       ValueColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
