// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/transaction"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateCreated holds the value of the "dateCreated" field.
	DateCreated string `json:"dateCreated,omitempty"`
	// DateUpdated holds the value of the "dateUpdated" field.
	DateUpdated string `json:"dateUpdated,omitempty"`
	// Status holds the value of the "status" field.
	Status transaction.Status `json:"status,omitempty"`
	// PaymentMethod holds the value of the "paymentMethod" field.
	PaymentMethod string `json:"paymentMethod,omitempty"`
	// PaymentStatus holds the value of the "paymentStatus" field.
	PaymentStatus string `json:"paymentStatus,omitempty"`
	// PaymentId holds the value of the "paymentId" field.
	PaymentId string `json:"paymentId,omitempty"`
	// PaymentAmount holds the value of the "paymentAmount" field.
	PaymentAmount string `json:"paymentAmount,omitempty"`
	// PaymentCurrency holds the value of the "paymentCurrency" field.
	PaymentCurrency string `json:"paymentCurrency,omitempty"`
	// PaymentDate holds the value of the "paymentDate" field.
	PaymentDate string `json:"paymentDate,omitempty"`
	// PaymentFee holds the value of the "paymentFee" field.
	PaymentFee string `json:"paymentFee,omitempty"`
	// PaymentNet holds the value of the "paymentNet" field.
	PaymentNet string `json:"paymentNet,omitempty"`
	// PaymentPayerEmail holds the value of the "paymentPayerEmail" field.
	PaymentPayerEmail string `json:"paymentPayerEmail,omitempty"`
	// PaymentPayerFirstName holds the value of the "paymentPayerFirstName" field.
	PaymentPayerFirstName string `json:"paymentPayerFirstName,omitempty"`
	// PaymentPayerLastName holds the value of the "paymentPayerLastName" field.
	PaymentPayerLastName string `json:"paymentPayerLastName,omitempty"`
	// PaymentPayerId holds the value of the "paymentPayerId" field.
	PaymentPayerId string `json:"paymentPayerId,omitempty"`
	// PaymentPayerStatus holds the value of the "paymentPayerStatus" field.
	PaymentPayerStatus string `json:"paymentPayerStatus,omitempty"`
	// PaymentReceiverEmail holds the value of the "paymentReceiverEmail" field.
	PaymentReceiverEmail string `json:"paymentReceiverEmail,omitempty"`
	// PaymentReceiverId holds the value of the "paymentReceiverId" field.
	PaymentReceiverId string `json:"paymentReceiverId,omitempty"`
	// PaymentTax holds the value of the "paymentTax" field.
	PaymentTax string `json:"paymentTax,omitempty"`
	// PaymentTransactionId holds the value of the "paymentTransactionId" field.
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	// PaymentTransactionType holds the value of the "paymentTransactionType" field.
	PaymentTransactionType string `json:"paymentTransactionType,omitempty"`
	// PaymentPendingReason holds the value of the "paymentPendingReason" field.
	PaymentPendingReason string `json:"paymentPendingReason,omitempty"`
	// PaymentReasonCode holds the value of the "paymentReasonCode" field.
	PaymentReasonCode string `json:"paymentReasonCode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                           TransactionEdges `json:"edges"`
	group_buy_transaction           *int
	refund_transactions_transaction *int
	shop_transactions               *int
	user_buyer_transactions         *int
	selectValues                    sql.SelectValues
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// OriginLink holds the value of the originLink edge.
	OriginLink []*LinkVisit `json:"originLink,omitempty"`
	// ProductCustomer holds the value of the productCustomer edge.
	ProductCustomer []*UserBuyer `json:"productCustomer,omitempty"`
	// Shop holds the value of the shop edge.
	Shop []*Shop `json:"shop,omitempty"`
	// ProductInfluencer holds the value of the productInfluencer edge.
	ProductInfluencer []*UserInfluencer `json:"productInfluencer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedProduct           map[string][]*Product
	namedOriginLink        map[string][]*LinkVisit
	namedProductCustomer   map[string][]*UserBuyer
	namedShop              map[string][]*Shop
	namedProductInfluencer map[string][]*UserInfluencer
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// OriginLinkOrErr returns the OriginLink value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) OriginLinkOrErr() ([]*LinkVisit, error) {
	if e.loadedTypes[1] {
		return e.OriginLink, nil
	}
	return nil, &NotLoadedError{edge: "originLink"}
}

// ProductCustomerOrErr returns the ProductCustomer value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) ProductCustomerOrErr() ([]*UserBuyer, error) {
	if e.loadedTypes[2] {
		return e.ProductCustomer, nil
	}
	return nil, &NotLoadedError{edge: "productCustomer"}
}

// ShopOrErr returns the Shop value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) ShopOrErr() ([]*Shop, error) {
	if e.loadedTypes[3] {
		return e.Shop, nil
	}
	return nil, &NotLoadedError{edge: "shop"}
}

// ProductInfluencerOrErr returns the ProductInfluencer value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) ProductInfluencerOrErr() ([]*UserInfluencer, error) {
	if e.loadedTypes[4] {
		return e.ProductInfluencer, nil
	}
	return nil, &NotLoadedError{edge: "productInfluencer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			values[i] = new(sql.NullInt64)
		case transaction.FieldDateCreated, transaction.FieldDateUpdated, transaction.FieldStatus, transaction.FieldPaymentMethod, transaction.FieldPaymentStatus, transaction.FieldPaymentId, transaction.FieldPaymentAmount, transaction.FieldPaymentCurrency, transaction.FieldPaymentDate, transaction.FieldPaymentFee, transaction.FieldPaymentNet, transaction.FieldPaymentPayerEmail, transaction.FieldPaymentPayerFirstName, transaction.FieldPaymentPayerLastName, transaction.FieldPaymentPayerId, transaction.FieldPaymentPayerStatus, transaction.FieldPaymentReceiverEmail, transaction.FieldPaymentReceiverId, transaction.FieldPaymentTax, transaction.FieldPaymentTransactionId, transaction.FieldPaymentTransactionType, transaction.FieldPaymentPendingReason, transaction.FieldPaymentReasonCode:
			values[i] = new(sql.NullString)
		case transaction.ForeignKeys[0]: // group_buy_transaction
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[1]: // refund_transactions_transaction
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[2]: // shop_transactions
			values[i] = new(sql.NullInt64)
		case transaction.ForeignKeys[3]: // user_buyer_transactions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transaction.FieldDateCreated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dateCreated", values[i])
			} else if value.Valid {
				t.DateCreated = value.String
			}
		case transaction.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dateUpdated", values[i])
			} else if value.Valid {
				t.DateUpdated = value.String
			}
		case transaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = transaction.Status(value.String)
			}
		case transaction.FieldPaymentMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentMethod", values[i])
			} else if value.Valid {
				t.PaymentMethod = value.String
			}
		case transaction.FieldPaymentStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentStatus", values[i])
			} else if value.Valid {
				t.PaymentStatus = value.String
			}
		case transaction.FieldPaymentId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentId", values[i])
			} else if value.Valid {
				t.PaymentId = value.String
			}
		case transaction.FieldPaymentAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentAmount", values[i])
			} else if value.Valid {
				t.PaymentAmount = value.String
			}
		case transaction.FieldPaymentCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentCurrency", values[i])
			} else if value.Valid {
				t.PaymentCurrency = value.String
			}
		case transaction.FieldPaymentDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentDate", values[i])
			} else if value.Valid {
				t.PaymentDate = value.String
			}
		case transaction.FieldPaymentFee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentFee", values[i])
			} else if value.Valid {
				t.PaymentFee = value.String
			}
		case transaction.FieldPaymentNet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentNet", values[i])
			} else if value.Valid {
				t.PaymentNet = value.String
			}
		case transaction.FieldPaymentPayerEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentPayerEmail", values[i])
			} else if value.Valid {
				t.PaymentPayerEmail = value.String
			}
		case transaction.FieldPaymentPayerFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentPayerFirstName", values[i])
			} else if value.Valid {
				t.PaymentPayerFirstName = value.String
			}
		case transaction.FieldPaymentPayerLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentPayerLastName", values[i])
			} else if value.Valid {
				t.PaymentPayerLastName = value.String
			}
		case transaction.FieldPaymentPayerId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentPayerId", values[i])
			} else if value.Valid {
				t.PaymentPayerId = value.String
			}
		case transaction.FieldPaymentPayerStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentPayerStatus", values[i])
			} else if value.Valid {
				t.PaymentPayerStatus = value.String
			}
		case transaction.FieldPaymentReceiverEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentReceiverEmail", values[i])
			} else if value.Valid {
				t.PaymentReceiverEmail = value.String
			}
		case transaction.FieldPaymentReceiverId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentReceiverId", values[i])
			} else if value.Valid {
				t.PaymentReceiverId = value.String
			}
		case transaction.FieldPaymentTax:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentTax", values[i])
			} else if value.Valid {
				t.PaymentTax = value.String
			}
		case transaction.FieldPaymentTransactionId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentTransactionId", values[i])
			} else if value.Valid {
				t.PaymentTransactionId = value.String
			}
		case transaction.FieldPaymentTransactionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentTransactionType", values[i])
			} else if value.Valid {
				t.PaymentTransactionType = value.String
			}
		case transaction.FieldPaymentPendingReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentPendingReason", values[i])
			} else if value.Valid {
				t.PaymentPendingReason = value.String
			}
		case transaction.FieldPaymentReasonCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentReasonCode", values[i])
			} else if value.Valid {
				t.PaymentReasonCode = value.String
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_buy_transaction", value)
			} else if value.Valid {
				t.group_buy_transaction = new(int)
				*t.group_buy_transaction = int(value.Int64)
			}
		case transaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field refund_transactions_transaction", value)
			} else if value.Valid {
				t.refund_transactions_transaction = new(int)
				*t.refund_transactions_transaction = int(value.Int64)
			}
		case transaction.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shop_transactions", value)
			} else if value.Valid {
				t.shop_transactions = new(int)
				*t.shop_transactions = int(value.Int64)
			}
		case transaction.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_buyer_transactions", value)
			} else if value.Valid {
				t.user_buyer_transactions = new(int)
				*t.user_buyer_transactions = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Transaction.
// This includes values selected through modifiers, order, etc.
func (t *Transaction) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the Transaction entity.
func (t *Transaction) QueryProduct() *ProductQuery {
	return NewTransactionClient(t.config).QueryProduct(t)
}

// QueryOriginLink queries the "originLink" edge of the Transaction entity.
func (t *Transaction) QueryOriginLink() *LinkVisitQuery {
	return NewTransactionClient(t.config).QueryOriginLink(t)
}

// QueryProductCustomer queries the "productCustomer" edge of the Transaction entity.
func (t *Transaction) QueryProductCustomer() *UserBuyerQuery {
	return NewTransactionClient(t.config).QueryProductCustomer(t)
}

// QueryShop queries the "shop" edge of the Transaction entity.
func (t *Transaction) QueryShop() *ShopQuery {
	return NewTransactionClient(t.config).QueryShop(t)
}

// QueryProductInfluencer queries the "productInfluencer" edge of the Transaction entity.
func (t *Transaction) QueryProductInfluencer() *UserInfluencerQuery {
	return NewTransactionClient(t.config).QueryProductInfluencer(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return NewTransactionClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("dateCreated=")
	builder.WriteString(t.DateCreated)
	builder.WriteString(", ")
	builder.WriteString("dateUpdated=")
	builder.WriteString(t.DateUpdated)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("paymentMethod=")
	builder.WriteString(t.PaymentMethod)
	builder.WriteString(", ")
	builder.WriteString("paymentStatus=")
	builder.WriteString(t.PaymentStatus)
	builder.WriteString(", ")
	builder.WriteString("paymentId=")
	builder.WriteString(t.PaymentId)
	builder.WriteString(", ")
	builder.WriteString("paymentAmount=")
	builder.WriteString(t.PaymentAmount)
	builder.WriteString(", ")
	builder.WriteString("paymentCurrency=")
	builder.WriteString(t.PaymentCurrency)
	builder.WriteString(", ")
	builder.WriteString("paymentDate=")
	builder.WriteString(t.PaymentDate)
	builder.WriteString(", ")
	builder.WriteString("paymentFee=")
	builder.WriteString(t.PaymentFee)
	builder.WriteString(", ")
	builder.WriteString("paymentNet=")
	builder.WriteString(t.PaymentNet)
	builder.WriteString(", ")
	builder.WriteString("paymentPayerEmail=")
	builder.WriteString(t.PaymentPayerEmail)
	builder.WriteString(", ")
	builder.WriteString("paymentPayerFirstName=")
	builder.WriteString(t.PaymentPayerFirstName)
	builder.WriteString(", ")
	builder.WriteString("paymentPayerLastName=")
	builder.WriteString(t.PaymentPayerLastName)
	builder.WriteString(", ")
	builder.WriteString("paymentPayerId=")
	builder.WriteString(t.PaymentPayerId)
	builder.WriteString(", ")
	builder.WriteString("paymentPayerStatus=")
	builder.WriteString(t.PaymentPayerStatus)
	builder.WriteString(", ")
	builder.WriteString("paymentReceiverEmail=")
	builder.WriteString(t.PaymentReceiverEmail)
	builder.WriteString(", ")
	builder.WriteString("paymentReceiverId=")
	builder.WriteString(t.PaymentReceiverId)
	builder.WriteString(", ")
	builder.WriteString("paymentTax=")
	builder.WriteString(t.PaymentTax)
	builder.WriteString(", ")
	builder.WriteString("paymentTransactionId=")
	builder.WriteString(t.PaymentTransactionId)
	builder.WriteString(", ")
	builder.WriteString("paymentTransactionType=")
	builder.WriteString(t.PaymentTransactionType)
	builder.WriteString(", ")
	builder.WriteString("paymentPendingReason=")
	builder.WriteString(t.PaymentPendingReason)
	builder.WriteString(", ")
	builder.WriteString("paymentReasonCode=")
	builder.WriteString(t.PaymentReasonCode)
	builder.WriteByte(')')
	return builder.String()
}

// NamedProduct returns the Product named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Transaction) NamedProduct(name string) ([]*Product, error) {
	if t.Edges.namedProduct == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedProduct[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Transaction) appendNamedProduct(name string, edges ...*Product) {
	if t.Edges.namedProduct == nil {
		t.Edges.namedProduct = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		t.Edges.namedProduct[name] = []*Product{}
	} else {
		t.Edges.namedProduct[name] = append(t.Edges.namedProduct[name], edges...)
	}
}

// NamedOriginLink returns the OriginLink named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Transaction) NamedOriginLink(name string) ([]*LinkVisit, error) {
	if t.Edges.namedOriginLink == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedOriginLink[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Transaction) appendNamedOriginLink(name string, edges ...*LinkVisit) {
	if t.Edges.namedOriginLink == nil {
		t.Edges.namedOriginLink = make(map[string][]*LinkVisit)
	}
	if len(edges) == 0 {
		t.Edges.namedOriginLink[name] = []*LinkVisit{}
	} else {
		t.Edges.namedOriginLink[name] = append(t.Edges.namedOriginLink[name], edges...)
	}
}

// NamedProductCustomer returns the ProductCustomer named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Transaction) NamedProductCustomer(name string) ([]*UserBuyer, error) {
	if t.Edges.namedProductCustomer == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedProductCustomer[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Transaction) appendNamedProductCustomer(name string, edges ...*UserBuyer) {
	if t.Edges.namedProductCustomer == nil {
		t.Edges.namedProductCustomer = make(map[string][]*UserBuyer)
	}
	if len(edges) == 0 {
		t.Edges.namedProductCustomer[name] = []*UserBuyer{}
	} else {
		t.Edges.namedProductCustomer[name] = append(t.Edges.namedProductCustomer[name], edges...)
	}
}

// NamedShop returns the Shop named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Transaction) NamedShop(name string) ([]*Shop, error) {
	if t.Edges.namedShop == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedShop[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Transaction) appendNamedShop(name string, edges ...*Shop) {
	if t.Edges.namedShop == nil {
		t.Edges.namedShop = make(map[string][]*Shop)
	}
	if len(edges) == 0 {
		t.Edges.namedShop[name] = []*Shop{}
	} else {
		t.Edges.namedShop[name] = append(t.Edges.namedShop[name], edges...)
	}
}

// NamedProductInfluencer returns the ProductInfluencer named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Transaction) NamedProductInfluencer(name string) ([]*UserInfluencer, error) {
	if t.Edges.namedProductInfluencer == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedProductInfluencer[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Transaction) appendNamedProductInfluencer(name string, edges ...*UserInfluencer) {
	if t.Edges.namedProductInfluencer == nil {
		t.Edges.namedProductInfluencer = make(map[string][]*UserInfluencer)
	}
	if len(edges) == 0 {
		t.Edges.namedProductInfluencer[name] = []*UserInfluencer{}
	} else {
		t.Edges.namedProductInfluencer[name] = append(t.Edges.namedProductInfluencer[name], edges...)
	}
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction
