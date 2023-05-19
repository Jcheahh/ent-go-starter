// Code generated by ent, DO NOT EDIT.

package blogpost

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the blogpost type in the database.
	Label = "blog_post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldDateUpdated holds the string denoting the dateupdated field in the database.
	FieldDateUpdated = "date_updated"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the blogpost in the database.
	Table = "blog_posts"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "user_sellers"
	// AuthorInverseTable is the table name for the UserSeller entity.
	// It exists in this package in order to avoid circular dependency with the "userseller" package.
	AuthorInverseTable = "user_sellers"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "blog_post_author"
)

// Columns holds all SQL columns for blogpost fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldDateCreated,
	FieldDateUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "blog_posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_blog_posts",
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

// OrderOption defines the ordering options for the BlogPost queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByDateUpdated orders the results by the dateUpdated field.
func ByDateUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateUpdated, opts...).ToFunc()
}

// ByAuthorCount orders the results by author count.
func ByAuthorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorStep(), opts...)
	}
}

// ByAuthor orders the results by author terms.
func ByAuthor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorTable, AuthorColumn),
	)
}