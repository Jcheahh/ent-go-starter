// Code generated by ent, DO NOT EDIT.

package shippingaddress

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the shippingaddress type in the database.
	Label = "shipping_address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldZip holds the string denoting the zip field in the database.
	FieldZip = "zip"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldDateUpdated holds the string denoting the dateupdated field in the database.
	FieldDateUpdated = "date_updated"
	// Table holds the table name of the shippingaddress in the database.
	Table = "shipping_addresses"
)

// Columns holds all SQL columns for shippingaddress fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAddress,
	FieldCity,
	FieldState,
	FieldZip,
	FieldCountry,
	FieldDateCreated,
	FieldDateUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "shipping_addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_shipping_addresses",
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

var (
	// DefaultDateCreated holds the default value on creation for the "dateCreated" field.
	DefaultDateCreated string
	// DefaultDateUpdated holds the default value on creation for the "dateUpdated" field.
	DefaultDateUpdated string
)

// OrderOption defines the ordering options for the ShippingAddress queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByZip orders the results by the zip field.
func ByZip(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldZip, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByDateUpdated orders the results by the dateUpdated field.
func ByDateUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateUpdated, opts...).ToFunc()
}
