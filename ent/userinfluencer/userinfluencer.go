// Code generated by ent, DO NOT EDIT.

package userinfluencer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the userinfluencer type in the database.
	Label = "user_influencer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPlaceholder holds the string denoting the placeholder field in the database.
	FieldPlaceholder = "placeholder"
	// EdgeUserProfile holds the string denoting the userprofile edge name in mutations.
	EdgeUserProfile = "userProfile"
	// EdgeReferralLinks holds the string denoting the referrallinks edge name in mutations.
	EdgeReferralLinks = "referralLinks"
	// EdgeReviews holds the string denoting the reviews edge name in mutations.
	EdgeReviews = "reviews"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// Table holds the table name of the userinfluencer in the database.
	Table = "user_influencers"
	// UserProfileTable is the table that holds the userProfile relation/edge.
	UserProfileTable = "users"
	// UserProfileInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserProfileInverseTable = "users"
	// UserProfileColumn is the table column denoting the userProfile relation/edge.
	UserProfileColumn = "user_influencer_user_profile"
	// ReferralLinksTable is the table that holds the referralLinks relation/edge.
	ReferralLinksTable = "referral_links"
	// ReferralLinksInverseTable is the table name for the ReferralLink entity.
	// It exists in this package in order to avoid circular dependency with the "referrallink" package.
	ReferralLinksInverseTable = "referral_links"
	// ReferralLinksColumn is the table column denoting the referralLinks relation/edge.
	ReferralLinksColumn = "user_influencer_referral_links"
	// ReviewsTable is the table that holds the reviews relation/edge.
	ReviewsTable = "reviews"
	// ReviewsInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewsInverseTable = "reviews"
	// ReviewsColumn is the table column denoting the reviews relation/edge.
	ReviewsColumn = "user_influencer_reviews"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "user_influencer_products"
	// TagsTable is the table that holds the tags relation/edge.
	TagsTable = "tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// TagsColumn is the table column denoting the tags relation/edge.
	TagsColumn = "user_influencer_tags"
)

// Columns holds all SQL columns for userinfluencer fields.
var Columns = []string{
	FieldID,
	FieldPlaceholder,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_influencers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"transaction_product_influencer",
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

// Order defines the ordering method for the UserInfluencer queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPlaceholder orders the results by the placeholder field.
func ByPlaceholder(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPlaceholder, opts...).ToFunc()
}

// ByUserProfileCount orders the results by userProfile count.
func ByUserProfileCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserProfileStep(), opts...)
	}
}

// ByUserProfile orders the results by userProfile terms.
func ByUserProfile(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserProfileStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReferralLinksCount orders the results by referralLinks count.
func ByReferralLinksCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReferralLinksStep(), opts...)
	}
}

// ByReferralLinks orders the results by referralLinks terms.
func ByReferralLinks(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReferralLinksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReviewsCount orders the results by reviews count.
func ByReviewsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReviewsStep(), opts...)
	}
}

// ByReviews orders the results by reviews terms.
func ByReviews(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReviewsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserProfileTable, UserProfileColumn),
	)
}
func newReferralLinksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReferralLinksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReferralLinksTable, ReferralLinksColumn),
	)
}
func newReviewsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReviewsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
	)
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TagsTable, TagsColumn),
	)
}
