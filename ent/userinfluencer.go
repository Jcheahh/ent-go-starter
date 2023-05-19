// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/userinfluencer"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserInfluencer is the model entity for the UserInfluencer schema.
type UserInfluencer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Placeholder holds the value of the "placeholder" field.
	Placeholder int `json:"placeholder,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserInfluencerQuery when eager-loading is set.
	Edges                          UserInfluencerEdges `json:"edges"`
	transaction_product_influencer *int
	selectValues                   sql.SelectValues
}

// UserInfluencerEdges holds the relations/edges for other nodes in the graph.
type UserInfluencerEdges struct {
	// UserProfile holds the value of the userProfile edge.
	UserProfile []*User `json:"userProfile,omitempty"`
	// ReferralLinks holds the value of the referralLinks edge.
	ReferralLinks []*ReferralLink `json:"referralLinks,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedUserProfile   map[string][]*User
	namedReferralLinks map[string][]*ReferralLink
	namedReviews       map[string][]*Review
	namedProducts      map[string][]*Product
	namedTags          map[string][]*Tag
}

// UserProfileOrErr returns the UserProfile value or an error if the edge
// was not loaded in eager-loading.
func (e UserInfluencerEdges) UserProfileOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.UserProfile, nil
	}
	return nil, &NotLoadedError{edge: "userProfile"}
}

// ReferralLinksOrErr returns the ReferralLinks value or an error if the edge
// was not loaded in eager-loading.
func (e UserInfluencerEdges) ReferralLinksOrErr() ([]*ReferralLink, error) {
	if e.loadedTypes[1] {
		return e.ReferralLinks, nil
	}
	return nil, &NotLoadedError{edge: "referralLinks"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e UserInfluencerEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[2] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e UserInfluencerEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[3] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e UserInfluencerEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[4] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserInfluencer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userinfluencer.FieldID, userinfluencer.FieldPlaceholder:
			values[i] = new(sql.NullInt64)
		case userinfluencer.ForeignKeys[0]: // transaction_product_influencer
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserInfluencer fields.
func (ui *UserInfluencer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userinfluencer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ui.ID = int(value.Int64)
		case userinfluencer.FieldPlaceholder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field placeholder", values[i])
			} else if value.Valid {
				ui.Placeholder = int(value.Int64)
			}
		case userinfluencer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_product_influencer", value)
			} else if value.Valid {
				ui.transaction_product_influencer = new(int)
				*ui.transaction_product_influencer = int(value.Int64)
			}
		default:
			ui.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserInfluencer.
// This includes values selected through modifiers, order, etc.
func (ui *UserInfluencer) Value(name string) (ent.Value, error) {
	return ui.selectValues.Get(name)
}

// QueryUserProfile queries the "userProfile" edge of the UserInfluencer entity.
func (ui *UserInfluencer) QueryUserProfile() *UserQuery {
	return NewUserInfluencerClient(ui.config).QueryUserProfile(ui)
}

// QueryReferralLinks queries the "referralLinks" edge of the UserInfluencer entity.
func (ui *UserInfluencer) QueryReferralLinks() *ReferralLinkQuery {
	return NewUserInfluencerClient(ui.config).QueryReferralLinks(ui)
}

// QueryReviews queries the "reviews" edge of the UserInfluencer entity.
func (ui *UserInfluencer) QueryReviews() *ReviewQuery {
	return NewUserInfluencerClient(ui.config).QueryReviews(ui)
}

// QueryProducts queries the "products" edge of the UserInfluencer entity.
func (ui *UserInfluencer) QueryProducts() *ProductQuery {
	return NewUserInfluencerClient(ui.config).QueryProducts(ui)
}

// QueryTags queries the "tags" edge of the UserInfluencer entity.
func (ui *UserInfluencer) QueryTags() *TagQuery {
	return NewUserInfluencerClient(ui.config).QueryTags(ui)
}

// Update returns a builder for updating this UserInfluencer.
// Note that you need to call UserInfluencer.Unwrap() before calling this method if this UserInfluencer
// was returned from a transaction, and the transaction was committed or rolled back.
func (ui *UserInfluencer) Update() *UserInfluencerUpdateOne {
	return NewUserInfluencerClient(ui.config).UpdateOne(ui)
}

// Unwrap unwraps the UserInfluencer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ui *UserInfluencer) Unwrap() *UserInfluencer {
	_tx, ok := ui.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserInfluencer is not a transactional entity")
	}
	ui.config.driver = _tx.drv
	return ui
}

// String implements the fmt.Stringer.
func (ui *UserInfluencer) String() string {
	var builder strings.Builder
	builder.WriteString("UserInfluencer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ui.ID))
	builder.WriteString("placeholder=")
	builder.WriteString(fmt.Sprintf("%v", ui.Placeholder))
	builder.WriteByte(')')
	return builder.String()
}

// NamedUserProfile returns the UserProfile named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ui *UserInfluencer) NamedUserProfile(name string) ([]*User, error) {
	if ui.Edges.namedUserProfile == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ui.Edges.namedUserProfile[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ui *UserInfluencer) appendNamedUserProfile(name string, edges ...*User) {
	if ui.Edges.namedUserProfile == nil {
		ui.Edges.namedUserProfile = make(map[string][]*User)
	}
	if len(edges) == 0 {
		ui.Edges.namedUserProfile[name] = []*User{}
	} else {
		ui.Edges.namedUserProfile[name] = append(ui.Edges.namedUserProfile[name], edges...)
	}
}

// NamedReferralLinks returns the ReferralLinks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ui *UserInfluencer) NamedReferralLinks(name string) ([]*ReferralLink, error) {
	if ui.Edges.namedReferralLinks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ui.Edges.namedReferralLinks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ui *UserInfluencer) appendNamedReferralLinks(name string, edges ...*ReferralLink) {
	if ui.Edges.namedReferralLinks == nil {
		ui.Edges.namedReferralLinks = make(map[string][]*ReferralLink)
	}
	if len(edges) == 0 {
		ui.Edges.namedReferralLinks[name] = []*ReferralLink{}
	} else {
		ui.Edges.namedReferralLinks[name] = append(ui.Edges.namedReferralLinks[name], edges...)
	}
}

// NamedReviews returns the Reviews named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ui *UserInfluencer) NamedReviews(name string) ([]*Review, error) {
	if ui.Edges.namedReviews == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ui.Edges.namedReviews[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ui *UserInfluencer) appendNamedReviews(name string, edges ...*Review) {
	if ui.Edges.namedReviews == nil {
		ui.Edges.namedReviews = make(map[string][]*Review)
	}
	if len(edges) == 0 {
		ui.Edges.namedReviews[name] = []*Review{}
	} else {
		ui.Edges.namedReviews[name] = append(ui.Edges.namedReviews[name], edges...)
	}
}

// NamedProducts returns the Products named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ui *UserInfluencer) NamedProducts(name string) ([]*Product, error) {
	if ui.Edges.namedProducts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ui.Edges.namedProducts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ui *UserInfluencer) appendNamedProducts(name string, edges ...*Product) {
	if ui.Edges.namedProducts == nil {
		ui.Edges.namedProducts = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		ui.Edges.namedProducts[name] = []*Product{}
	} else {
		ui.Edges.namedProducts[name] = append(ui.Edges.namedProducts[name], edges...)
	}
}

// NamedTags returns the Tags named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ui *UserInfluencer) NamedTags(name string) ([]*Tag, error) {
	if ui.Edges.namedTags == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ui.Edges.namedTags[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ui *UserInfluencer) appendNamedTags(name string, edges ...*Tag) {
	if ui.Edges.namedTags == nil {
		ui.Edges.namedTags = make(map[string][]*Tag)
	}
	if len(edges) == 0 {
		ui.Edges.namedTags[name] = []*Tag{}
	} else {
		ui.Edges.namedTags[name] = append(ui.Edges.namedTags[name], edges...)
	}
}

// UserInfluencers is a parsable slice of UserInfluencer.
type UserInfluencers []*UserInfluencer