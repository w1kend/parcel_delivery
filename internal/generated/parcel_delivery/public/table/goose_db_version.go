//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var GooseDbVersion = newGooseDbVersionTable("public", "goose_db_version", "")

type gooseDbVersionTable struct {
	postgres.Table

	//Columns
	ID        postgres.ColumnInteger
	VersionID postgres.ColumnInteger
	IsApplied postgres.ColumnBool
	Tstamp    postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type GooseDbVersionTable struct {
	gooseDbVersionTable

	EXCLUDED gooseDbVersionTable
}

// AS creates new GooseDbVersionTable with assigned alias
func (a GooseDbVersionTable) AS(alias string) *GooseDbVersionTable {
	return newGooseDbVersionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GooseDbVersionTable with assigned schema name
func (a GooseDbVersionTable) FromSchema(schemaName string) *GooseDbVersionTable {
	return newGooseDbVersionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GooseDbVersionTable with assigned table prefix
func (a GooseDbVersionTable) WithPrefix(prefix string) *GooseDbVersionTable {
	return newGooseDbVersionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GooseDbVersionTable with assigned table suffix
func (a GooseDbVersionTable) WithSuffix(suffix string) *GooseDbVersionTable {
	return newGooseDbVersionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGooseDbVersionTable(schemaName, tableName, alias string) *GooseDbVersionTable {
	return &GooseDbVersionTable{
		gooseDbVersionTable: newGooseDbVersionTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newGooseDbVersionTableImpl("", "excluded", ""),
	}
}

func newGooseDbVersionTableImpl(schemaName, tableName, alias string) gooseDbVersionTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		VersionIDColumn = postgres.IntegerColumn("version_id")
		IsAppliedColumn = postgres.BoolColumn("is_applied")
		TstampColumn    = postgres.TimestampColumn("tstamp")
		allColumns      = postgres.ColumnList{IDColumn, VersionIDColumn, IsAppliedColumn, TstampColumn}
		mutableColumns  = postgres.ColumnList{VersionIDColumn, IsAppliedColumn, TstampColumn}
	)

	return gooseDbVersionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		VersionID: VersionIDColumn,
		IsApplied: IsAppliedColumn,
		Tstamp:    TstampColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
