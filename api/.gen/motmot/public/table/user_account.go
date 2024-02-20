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

var UserAccount = newUserAccountTable("public", "user_account", "")

type userAccountTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	DisplayName postgres.ColumnString
	Hash        postgres.ColumnString
	Email       postgres.ColumnString
	Bio         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserAccountTable struct {
	userAccountTable

	EXCLUDED userAccountTable
}

// AS creates new UserAccountTable with assigned alias
func (a UserAccountTable) AS(alias string) *UserAccountTable {
	return newUserAccountTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserAccountTable with assigned schema name
func (a UserAccountTable) FromSchema(schemaName string) *UserAccountTable {
	return newUserAccountTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserAccountTable with assigned table prefix
func (a UserAccountTable) WithPrefix(prefix string) *UserAccountTable {
	return newUserAccountTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserAccountTable with assigned table suffix
func (a UserAccountTable) WithSuffix(suffix string) *UserAccountTable {
	return newUserAccountTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserAccountTable(schemaName, tableName, alias string) *UserAccountTable {
	return &UserAccountTable{
		userAccountTable: newUserAccountTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newUserAccountTableImpl("", "excluded", ""),
	}
}

func newUserAccountTableImpl(schemaName, tableName, alias string) userAccountTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		DisplayNameColumn = postgres.StringColumn("display_name")
		HashColumn        = postgres.StringColumn("hash")
		EmailColumn       = postgres.StringColumn("email")
		BioColumn         = postgres.StringColumn("bio")
		allColumns        = postgres.ColumnList{IDColumn, DisplayNameColumn, HashColumn, EmailColumn, BioColumn}
		mutableColumns    = postgres.ColumnList{DisplayNameColumn, HashColumn, EmailColumn, BioColumn}
	)

	return userAccountTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		DisplayName: DisplayNameColumn,
		Hash:        HashColumn,
		Email:       EmailColumn,
		Bio:         BioColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
