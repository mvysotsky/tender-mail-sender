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

var Geoip = newGeoipTable("newtendex", "geoip", "")

type geoipTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	RangeBegin  mysql.ColumnInteger
	RangeEnd    mysql.ColumnInteger
	CountryCode mysql.ColumnString
	CountryName mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type GeoipTable struct {
	geoipTable

	NEW geoipTable
}

// AS creates new GeoipTable with assigned alias
func (a GeoipTable) AS(alias string) *GeoipTable {
	return newGeoipTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GeoipTable with assigned schema name
func (a GeoipTable) FromSchema(schemaName string) *GeoipTable {
	return newGeoipTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GeoipTable with assigned table prefix
func (a GeoipTable) WithPrefix(prefix string) *GeoipTable {
	return newGeoipTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GeoipTable with assigned table suffix
func (a GeoipTable) WithSuffix(suffix string) *GeoipTable {
	return newGeoipTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGeoipTable(schemaName, tableName, alias string) *GeoipTable {
	return &GeoipTable{
		geoipTable: newGeoipTableImpl(schemaName, tableName, alias),
		NEW:        newGeoipTableImpl("", "new", ""),
	}
}

func newGeoipTableImpl(schemaName, tableName, alias string) geoipTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		RangeBeginColumn  = mysql.IntegerColumn("range_begin")
		RangeEndColumn    = mysql.IntegerColumn("range_end")
		CountryCodeColumn = mysql.StringColumn("country_code")
		CountryNameColumn = mysql.StringColumn("country_name")
		allColumns        = mysql.ColumnList{IDColumn, RangeBeginColumn, RangeEndColumn, CountryCodeColumn, CountryNameColumn}
		mutableColumns    = mysql.ColumnList{RangeBeginColumn, RangeEndColumn, CountryCodeColumn, CountryNameColumn}
	)

	return geoipTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		RangeBegin:  RangeBeginColumn,
		RangeEnd:    RangeEndColumn,
		CountryCode: CountryCodeColumn,
		CountryName: CountryNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
