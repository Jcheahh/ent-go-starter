// Code generated by ent, DO NOT EDIT.

package primarycontent

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the primarycontent type in the database.
	Label = "primary_content"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPlaceholder holds the string denoting the placeholder field in the database.
	FieldPlaceholder = "placeholder"
	// EdgeContentBlock holds the string denoting the contentblock edge name in mutations.
	EdgeContentBlock = "contentBlock"
	// Table holds the table name of the primarycontent in the database.
	Table = "primary_contents"
	// ContentBlockTable is the table that holds the contentBlock relation/edge.
	ContentBlockTable = "content_blocks"
	// ContentBlockInverseTable is the table name for the ContentBlock entity.
	// It exists in this package in order to avoid circular dependency with the "contentblock" package.
	ContentBlockInverseTable = "content_blocks"
	// ContentBlockColumn is the table column denoting the contentBlock relation/edge.
	ContentBlockColumn = "primary_content_content_block"
)

// Columns holds all SQL columns for primarycontent fields.
var Columns = []string{
	FieldID,
	FieldPlaceholder,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "primary_contents"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_page_view_primary_content",
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

// OrderOption defines the ordering options for the PrimaryContent queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPlaceholder orders the results by the placeholder field.
func ByPlaceholder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlaceholder, opts...).ToFunc()
}

// ByContentBlockCount orders the results by contentBlock count.
func ByContentBlockCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContentBlockStep(), opts...)
	}
}

// ByContentBlock orders the results by contentBlock terms.
func ByContentBlock(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContentBlockStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newContentBlockStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContentBlockInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ContentBlockTable, ContentBlockColumn),
	)
}
