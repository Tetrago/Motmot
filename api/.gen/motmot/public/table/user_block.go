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

var UserBlock = newUserBlockTable("public", "user_block", "")

type userBlockTable struct {
	postgres.Table

	// Columns
	UserID      postgres.ColumnInteger
	BlockUserID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserBlockTable struct {
	userBlockTable

	EXCLUDED userBlockTable
}

// AS creates new UserBlockTable with assigned alias
func (a UserBlockTable) AS(alias string) *UserBlockTable {
	return newUserBlockTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserBlockTable with assigned schema name
func (a UserBlockTable) FromSchema(schemaName string) *UserBlockTable {
	return newUserBlockTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserBlockTable with assigned table prefix
func (a UserBlockTable) WithPrefix(prefix string) *UserBlockTable {
	return newUserBlockTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserBlockTable with assigned table suffix
func (a UserBlockTable) WithSuffix(suffix string) *UserBlockTable {
	return newUserBlockTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserBlockTable(schemaName, tableName, alias string) *UserBlockTable {
	return &UserBlockTable{
		userBlockTable: newUserBlockTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newUserBlockTableImpl("", "excluded", ""),
	}
}

func newUserBlockTableImpl(schemaName, tableName, alias string) userBlockTable {
	var (
		UserIDColumn      = postgres.IntegerColumn("user_id")
		BlockUserIDColumn = postgres.IntegerColumn("block_user_id")
		allColumns        = postgres.ColumnList{UserIDColumn, BlockUserIDColumn}
		mutableColumns    = postgres.ColumnList{}
	)

	return userBlockTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:      UserIDColumn,
		BlockUserID: BlockUserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
