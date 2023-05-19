// Code generated by ent, DO NOT EDIT.

package shop

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the shop type in the database.
	Label = "shop"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// EdgeBankAccounts holds the string denoting the bankaccounts edge name in mutations.
	EdgeBankAccounts = "bankAccounts"
	// Table holds the table name of the shop in the database.
	Table = "shops"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "shop_products"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transactions"
	// TransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionsInverseTable = "transactions"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "shop_transactions"
	// BankAccountsTable is the table that holds the bankAccounts relation/edge.
	BankAccountsTable = "bank_accounts"
	// BankAccountsInverseTable is the table name for the BankAccount entity.
	// It exists in this package in order to avoid circular dependency with the "bankaccount" package.
	BankAccountsInverseTable = "bank_accounts"
	// BankAccountsColumn is the table column denoting the bankAccounts relation/edge.
	BankAccountsColumn = "shop_bank_accounts"
)

// Columns holds all SQL columns for shop fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "shops"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_shop",
	"transaction_shop",
	"user_seller_shops",
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

// OrderOption defines the ordering options for the Shop queries.
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

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTransactionsCount orders the results by transactions count.
func ByTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionsStep(), opts...)
	}
}

// ByTransactions orders the results by transactions terms.
func ByTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBankAccountsCount orders the results by bankAccounts count.
func ByBankAccountsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBankAccountsStep(), opts...)
	}
}

// ByBankAccounts orders the results by bankAccounts terms.
func ByBankAccounts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBankAccountsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
	)
}
func newTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
	)
}
func newBankAccountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BankAccountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BankAccountsTable, BankAccountsColumn),
	)
}