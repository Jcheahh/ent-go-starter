// Code generated by ent, DO NOT EDIT.

package chat

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the chat type in the database.
	Label = "chat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldXid holds the string denoting the xid field in the database.
	FieldXid = "xid"
	// Table holds the table name of the chat in the database.
	Table = "chats"
)

// Columns holds all SQL columns for chat fields.
var Columns = []string{
	FieldID,
	FieldXid,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "chats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_chats",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Chat queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByXid orders the results by the xid field.
func ByXid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXid, opts...).ToFunc()
}
