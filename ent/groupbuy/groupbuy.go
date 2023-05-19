// Code generated by ent, DO NOT EDIT.

package groupbuy

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the groupbuy type in the database.
	Label = "group_buy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldProductPrice holds the string denoting the productprice field in the database.
	FieldProductPrice = "product_price"
	// FieldMoq holds the string denoting the moq field in the database.
	FieldMoq = "moq"
	// FieldStartDate holds the string denoting the startdate field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the enddate field in the database.
	FieldEndDate = "end_date"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"
	// Table holds the table name of the groupbuy in the database.
	Table = "group_buys"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "group_buy_product"
	// TransactionTable is the table that holds the transaction relation/edge.
	TransactionTable = "transactions"
	// TransactionInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionInverseTable = "transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "group_buy_transaction"
)

// Columns holds all SQL columns for groupbuy fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldProductPrice,
	FieldMoq,
	FieldStartDate,
	FieldEndDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "group_buys"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_group_buys",
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

// OrderOption defines the ordering options for the GroupBuy queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByProductPrice orders the results by the productPrice field.
func ByProductPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductPrice, opts...).ToFunc()
}

// ByMoq orders the results by the moq field.
func ByMoq(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMoq, opts...).ToFunc()
}

// ByStartDate orders the results by the startDate field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the endDate field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
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

// ByTransactionCount orders the results by transaction count.
func ByTransactionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionStep(), opts...)
	}
}

// ByTransaction orders the results by transaction terms.
func ByTransaction(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
	)
}
func newTransactionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionTable, TransactionColumn),
	)
}