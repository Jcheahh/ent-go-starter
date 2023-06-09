// Code generated by ent, DO NOT EDIT.

package emailcampaign

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the emailcampaign type in the database.
	Label = "email_campaign"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldXid holds the string denoting the xid field in the database.
	FieldXid = "xid"
	// Table holds the table name of the emailcampaign in the database.
	Table = "email_campaigns"
)

// Columns holds all SQL columns for emailcampaign fields.
var Columns = []string{
	FieldID,
	FieldXid,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "email_campaigns"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_email_campaign",
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

// OrderOption defines the ordering options for the EmailCampaign queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByXid orders the results by the xid field.
func ByXid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldXid, opts...).ToFunc()
}
