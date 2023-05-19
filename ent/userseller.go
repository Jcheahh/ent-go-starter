// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/userseller"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserSeller is the model entity for the UserSeller schema.
type UserSeller struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BrandName holds the value of the "brandName" field.
	BrandName string `json:"brandName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSellerQuery when eager-loading is set.
	Edges                                      UserSellerEdges `json:"edges"`
	blog_post_author                           *int
	commission_structure_schema_product_seller *int
	product_product_seller                     *int
	selectValues                               sql.SelectValues
}

// UserSellerEdges holds the relations/edges for other nodes in the graph.
type UserSellerEdges struct {
	// UserProfile holds the value of the userProfile edge.
	UserProfile []*User `json:"userProfile,omitempty"`
	// Shops holds the value of the shops edge.
	Shops []*Shop `json:"shops,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedUserProfile map[string][]*User
	namedShops       map[string][]*Shop
}

// UserProfileOrErr returns the UserProfile value or an error if the edge
// was not loaded in eager-loading.
func (e UserSellerEdges) UserProfileOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.UserProfile, nil
	}
	return nil, &NotLoadedError{edge: "userProfile"}
}

// ShopsOrErr returns the Shops value or an error if the edge
// was not loaded in eager-loading.
func (e UserSellerEdges) ShopsOrErr() ([]*Shop, error) {
	if e.loadedTypes[1] {
		return e.Shops, nil
	}
	return nil, &NotLoadedError{edge: "shops"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSeller) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userseller.FieldID:
			values[i] = new(sql.NullInt64)
		case userseller.FieldBrandName:
			values[i] = new(sql.NullString)
		case userseller.ForeignKeys[0]: // blog_post_author
			values[i] = new(sql.NullInt64)
		case userseller.ForeignKeys[1]: // commission_structure_schema_product_seller
			values[i] = new(sql.NullInt64)
		case userseller.ForeignKeys[2]: // product_product_seller
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSeller fields.
func (us *UserSeller) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userseller.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int(value.Int64)
		case userseller.FieldBrandName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field brandName", values[i])
			} else if value.Valid {
				us.BrandName = value.String
			}
		case userseller.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field blog_post_author", value)
			} else if value.Valid {
				us.blog_post_author = new(int)
				*us.blog_post_author = int(value.Int64)
			}
		case userseller.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field commission_structure_schema_product_seller", value)
			} else if value.Valid {
				us.commission_structure_schema_product_seller = new(int)
				*us.commission_structure_schema_product_seller = int(value.Int64)
			}
		case userseller.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_product_seller", value)
			} else if value.Valid {
				us.product_product_seller = new(int)
				*us.product_product_seller = int(value.Int64)
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSeller.
// This includes values selected through modifiers, order, etc.
func (us *UserSeller) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUserProfile queries the "userProfile" edge of the UserSeller entity.
func (us *UserSeller) QueryUserProfile() *UserQuery {
	return NewUserSellerClient(us.config).QueryUserProfile(us)
}

// QueryShops queries the "shops" edge of the UserSeller entity.
func (us *UserSeller) QueryShops() *ShopQuery {
	return NewUserSellerClient(us.config).QueryShops(us)
}

// Update returns a builder for updating this UserSeller.
// Note that you need to call UserSeller.Unwrap() before calling this method if this UserSeller
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSeller) Update() *UserSellerUpdateOne {
	return NewUserSellerClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSeller entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSeller) Unwrap() *UserSeller {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserSeller is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSeller) String() string {
	var builder strings.Builder
	builder.WriteString("UserSeller(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("brandName=")
	builder.WriteString(us.BrandName)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUserProfile returns the UserProfile named value or an error if the edge was not
// loaded in eager-loading with this name.
func (us *UserSeller) NamedUserProfile(name string) ([]*User, error) {
	if us.Edges.namedUserProfile == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := us.Edges.namedUserProfile[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (us *UserSeller) appendNamedUserProfile(name string, edges ...*User) {
	if us.Edges.namedUserProfile == nil {
		us.Edges.namedUserProfile = make(map[string][]*User)
	}
	if len(edges) == 0 {
		us.Edges.namedUserProfile[name] = []*User{}
	} else {
		us.Edges.namedUserProfile[name] = append(us.Edges.namedUserProfile[name], edges...)
	}
}

// NamedShops returns the Shops named value or an error if the edge was not
// loaded in eager-loading with this name.
func (us *UserSeller) NamedShops(name string) ([]*Shop, error) {
	if us.Edges.namedShops == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := us.Edges.namedShops[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (us *UserSeller) appendNamedShops(name string, edges ...*Shop) {
	if us.Edges.namedShops == nil {
		us.Edges.namedShops = make(map[string][]*Shop)
	}
	if len(edges) == 0 {
		us.Edges.namedShops[name] = []*Shop{}
	} else {
		us.Edges.namedShops[name] = append(us.Edges.namedShops[name], edges...)
	}
}

// UserSellers is a parsable slice of UserSeller.
type UserSellers []*UserSeller