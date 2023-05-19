// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/shop"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Shop is the model entity for the Shop schema.
type Shop struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShopQuery when eager-loading is set.
	Edges             ShopEdges `json:"edges"`
	product_shop      *int
	transaction_shop  *int
	user_seller_shops *int
	selectValues      sql.SelectValues
}

// ShopEdges holds the relations/edges for other nodes in the graph.
type ShopEdges struct {
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// BankAccounts holds the value of the bankAccounts edge.
	BankAccounts []*BankAccount `json:"bankAccounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedProducts     map[string][]*Product
	namedTransactions map[string][]*Transaction
	namedBankAccounts map[string][]*BankAccount
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e ShopEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e ShopEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[1] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// BankAccountsOrErr returns the BankAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e ShopEdges) BankAccountsOrErr() ([]*BankAccount, error) {
	if e.loadedTypes[2] {
		return e.BankAccounts, nil
	}
	return nil, &NotLoadedError{edge: "bankAccounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Shop) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shop.FieldID:
			values[i] = new(sql.NullInt64)
		case shop.FieldName, shop.FieldDescription:
			values[i] = new(sql.NullString)
		case shop.ForeignKeys[0]: // product_shop
			values[i] = new(sql.NullInt64)
		case shop.ForeignKeys[1]: // transaction_shop
			values[i] = new(sql.NullInt64)
		case shop.ForeignKeys[2]: // user_seller_shops
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Shop fields.
func (s *Shop) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shop.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case shop.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case shop.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case shop.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_shop", value)
			} else if value.Valid {
				s.product_shop = new(int)
				*s.product_shop = int(value.Int64)
			}
		case shop.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_shop", value)
			} else if value.Valid {
				s.transaction_shop = new(int)
				*s.transaction_shop = int(value.Int64)
			}
		case shop.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_seller_shops", value)
			} else if value.Valid {
				s.user_seller_shops = new(int)
				*s.user_seller_shops = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Shop.
// This includes values selected through modifiers, order, etc.
func (s *Shop) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryProducts queries the "products" edge of the Shop entity.
func (s *Shop) QueryProducts() *ProductQuery {
	return NewShopClient(s.config).QueryProducts(s)
}

// QueryTransactions queries the "transactions" edge of the Shop entity.
func (s *Shop) QueryTransactions() *TransactionQuery {
	return NewShopClient(s.config).QueryTransactions(s)
}

// QueryBankAccounts queries the "bankAccounts" edge of the Shop entity.
func (s *Shop) QueryBankAccounts() *BankAccountQuery {
	return NewShopClient(s.config).QueryBankAccounts(s)
}

// Update returns a builder for updating this Shop.
// Note that you need to call Shop.Unwrap() before calling this method if this Shop
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Shop) Update() *ShopUpdateOne {
	return NewShopClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Shop entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Shop) Unwrap() *Shop {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Shop is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Shop) String() string {
	var builder strings.Builder
	builder.WriteString("Shop(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedProducts returns the Products named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Shop) NamedProducts(name string) ([]*Product, error) {
	if s.Edges.namedProducts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedProducts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Shop) appendNamedProducts(name string, edges ...*Product) {
	if s.Edges.namedProducts == nil {
		s.Edges.namedProducts = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		s.Edges.namedProducts[name] = []*Product{}
	} else {
		s.Edges.namedProducts[name] = append(s.Edges.namedProducts[name], edges...)
	}
}

// NamedTransactions returns the Transactions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Shop) NamedTransactions(name string) ([]*Transaction, error) {
	if s.Edges.namedTransactions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedTransactions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Shop) appendNamedTransactions(name string, edges ...*Transaction) {
	if s.Edges.namedTransactions == nil {
		s.Edges.namedTransactions = make(map[string][]*Transaction)
	}
	if len(edges) == 0 {
		s.Edges.namedTransactions[name] = []*Transaction{}
	} else {
		s.Edges.namedTransactions[name] = append(s.Edges.namedTransactions[name], edges...)
	}
}

// NamedBankAccounts returns the BankAccounts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Shop) NamedBankAccounts(name string) ([]*BankAccount, error) {
	if s.Edges.namedBankAccounts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedBankAccounts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Shop) appendNamedBankAccounts(name string, edges ...*BankAccount) {
	if s.Edges.namedBankAccounts == nil {
		s.Edges.namedBankAccounts = make(map[string][]*BankAccount)
	}
	if len(edges) == 0 {
		s.Edges.namedBankAccounts[name] = []*BankAccount{}
	} else {
		s.Edges.namedBankAccounts[name] = append(s.Edges.namedBankAccounts[name], edges...)
	}
}

// Shops is a parsable slice of Shop.
type Shops []*Shop