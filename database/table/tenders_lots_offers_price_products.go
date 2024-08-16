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

var TendersLotsOffersPriceProducts = newTendersLotsOffersPriceProductsTable("newtendex", "tenders_lots_offers_price_products", "")

type tendersLotsOffersPriceProductsTable struct {
	mysql.Table

	// Columns
	ID               mysql.ColumnInteger
	TenderLotOfferID mysql.ColumnInteger
	ProductID        mysql.ColumnInteger
	Value            mysql.ColumnFloat

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TendersLotsOffersPriceProductsTable struct {
	tendersLotsOffersPriceProductsTable

	NEW tendersLotsOffersPriceProductsTable
}

// AS creates new TendersLotsOffersPriceProductsTable with assigned alias
func (a TendersLotsOffersPriceProductsTable) AS(alias string) *TendersLotsOffersPriceProductsTable {
	return newTendersLotsOffersPriceProductsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TendersLotsOffersPriceProductsTable with assigned schema name
func (a TendersLotsOffersPriceProductsTable) FromSchema(schemaName string) *TendersLotsOffersPriceProductsTable {
	return newTendersLotsOffersPriceProductsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TendersLotsOffersPriceProductsTable with assigned table prefix
func (a TendersLotsOffersPriceProductsTable) WithPrefix(prefix string) *TendersLotsOffersPriceProductsTable {
	return newTendersLotsOffersPriceProductsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TendersLotsOffersPriceProductsTable with assigned table suffix
func (a TendersLotsOffersPriceProductsTable) WithSuffix(suffix string) *TendersLotsOffersPriceProductsTable {
	return newTendersLotsOffersPriceProductsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTendersLotsOffersPriceProductsTable(schemaName, tableName, alias string) *TendersLotsOffersPriceProductsTable {
	return &TendersLotsOffersPriceProductsTable{
		tendersLotsOffersPriceProductsTable: newTendersLotsOffersPriceProductsTableImpl(schemaName, tableName, alias),
		NEW:                                 newTendersLotsOffersPriceProductsTableImpl("", "new", ""),
	}
}

func newTendersLotsOffersPriceProductsTableImpl(schemaName, tableName, alias string) tendersLotsOffersPriceProductsTable {
	var (
		IDColumn               = mysql.IntegerColumn("id")
		TenderLotOfferIDColumn = mysql.IntegerColumn("tender_lot_offer_id")
		ProductIDColumn        = mysql.IntegerColumn("product_id")
		ValueColumn            = mysql.FloatColumn("value")
		allColumns             = mysql.ColumnList{IDColumn, TenderLotOfferIDColumn, ProductIDColumn, ValueColumn}
		mutableColumns         = mysql.ColumnList{TenderLotOfferIDColumn, ProductIDColumn, ValueColumn}
	)

	return tendersLotsOffersPriceProductsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		TenderLotOfferID: TenderLotOfferIDColumn,
		ProductID:        ProductIDColumn,
		Value:            ValueColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
