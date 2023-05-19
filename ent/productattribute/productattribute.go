// Code generated by ent, DO NOT EDIT.

package productattribute

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the productattribute type in the database.
	Label = "product_attribute"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the productattribute in the database.
	Table = "product_attributes"
)

// Columns holds all SQL columns for productattribute fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "product_attributes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_product_attributes",
	"product_variation_product_attributes",
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

// OrderOption defines the ordering options for the ProductAttribute queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}