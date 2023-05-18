// Code generated by ent, DO NOT EDIT.

package linkvisit

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the linkvisit type in the database.
	Label = "link_visit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldIpAddress holds the string denoting the ipaddress field in the database.
	FieldIpAddress = "ip_address"
	// FieldSaleValue holds the string denoting the salevalue field in the database.
	FieldSaleValue = "sale_value"
	// FieldCommissionEarned holds the string denoting the commissionearned field in the database.
	FieldCommissionEarned = "commission_earned"
	// Table holds the table name of the linkvisit in the database.
	Table = "link_visits"
)

// Columns holds all SQL columns for linkvisit fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldIpAddress,
	FieldSaleValue,
	FieldCommissionEarned,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "link_visits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"referral_link_visits",
	"transaction_origin_link",
	"user_buyer_links_clicked",
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

// Order defines the ordering method for the LinkVisit queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByIpAddress orders the results by the ipAddress field.
func ByIpAddress(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldIpAddress, opts...).ToFunc()
}

// BySaleValue orders the results by the saleValue field.
func BySaleValue(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldSaleValue, opts...).ToFunc()
}

// ByCommissionEarned orders the results by the commissionEarned field.
func ByCommissionEarned(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCommissionEarned, opts...).ToFunc()
}
