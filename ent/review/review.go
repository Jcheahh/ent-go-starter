// Code generated by ent, DO NOT EDIT.

package review

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the review type in the database.
	Label = "review"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeProductCustomer holds the string denoting the productcustomer edge name in mutations.
	EdgeProductCustomer = "productCustomer"
	// Table holds the table name of the review in the database.
	Table = "reviews"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "review_product"
	// ProductCustomerTable is the table that holds the productCustomer relation/edge.
	ProductCustomerTable = "user_buyers"
	// ProductCustomerInverseTable is the table name for the UserBuyer entity.
	// It exists in this package in order to avoid circular dependency with the "userbuyer" package.
	ProductCustomerInverseTable = "user_buyers"
	// ProductCustomerColumn is the table column denoting the productCustomer relation/edge.
	ProductCustomerColumn = "review_product_customer"
)

// Columns holds all SQL columns for review fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldContent,
	FieldRating,
	FieldDateCreated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reviews"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_reviews",
	"user_buyer_reviews",
	"user_influencer_reviews",
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

// OrderOption defines the ordering options for the Review queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByRating orders the results by the rating field.
func ByRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRating, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByProductCount orders the results by product count.
func ByProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductStep(), opts...)
	}
}

// ByProduct orders the results by product terms.
func ByProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductCustomerCount orders the results by productCustomer count.
func ByProductCustomerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductCustomerStep(), opts...)
	}
}

// ByProductCustomer orders the results by productCustomer terms.
func ByProductCustomer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductCustomerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
	)
}
func newProductCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductCustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductCustomerTable, ProductCustomerColumn),
	)
}