// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldDateUpdated holds the string denoting the dateupdated field in the database.
	FieldDateUpdated = "date_updated"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPaymentMethod holds the string denoting the paymentmethod field in the database.
	FieldPaymentMethod = "payment_method"
	// FieldPaymentStatus holds the string denoting the paymentstatus field in the database.
	FieldPaymentStatus = "payment_status"
	// FieldPaymentId holds the string denoting the paymentid field in the database.
	FieldPaymentId = "payment_id"
	// FieldPaymentAmount holds the string denoting the paymentamount field in the database.
	FieldPaymentAmount = "payment_amount"
	// FieldPaymentCurrency holds the string denoting the paymentcurrency field in the database.
	FieldPaymentCurrency = "payment_currency"
	// FieldPaymentDate holds the string denoting the paymentdate field in the database.
	FieldPaymentDate = "payment_date"
	// FieldPaymentFee holds the string denoting the paymentfee field in the database.
	FieldPaymentFee = "payment_fee"
	// FieldPaymentNet holds the string denoting the paymentnet field in the database.
	FieldPaymentNet = "payment_net"
	// FieldPaymentPayerEmail holds the string denoting the paymentpayeremail field in the database.
	FieldPaymentPayerEmail = "payment_payer_email"
	// FieldPaymentPayerFirstName holds the string denoting the paymentpayerfirstname field in the database.
	FieldPaymentPayerFirstName = "payment_payer_first_name"
	// FieldPaymentPayerLastName holds the string denoting the paymentpayerlastname field in the database.
	FieldPaymentPayerLastName = "payment_payer_last_name"
	// FieldPaymentPayerId holds the string denoting the paymentpayerid field in the database.
	FieldPaymentPayerId = "payment_payer_id"
	// FieldPaymentPayerStatus holds the string denoting the paymentpayerstatus field in the database.
	FieldPaymentPayerStatus = "payment_payer_status"
	// FieldPaymentReceiverEmail holds the string denoting the paymentreceiveremail field in the database.
	FieldPaymentReceiverEmail = "payment_receiver_email"
	// FieldPaymentReceiverId holds the string denoting the paymentreceiverid field in the database.
	FieldPaymentReceiverId = "payment_receiver_id"
	// FieldPaymentTax holds the string denoting the paymenttax field in the database.
	FieldPaymentTax = "payment_tax"
	// FieldPaymentTransactionId holds the string denoting the paymenttransactionid field in the database.
	FieldPaymentTransactionId = "payment_transaction_id"
	// FieldPaymentTransactionType holds the string denoting the paymenttransactiontype field in the database.
	FieldPaymentTransactionType = "payment_transaction_type"
	// FieldPaymentPendingReason holds the string denoting the paymentpendingreason field in the database.
	FieldPaymentPendingReason = "payment_pending_reason"
	// FieldPaymentReasonCode holds the string denoting the paymentreasoncode field in the database.
	FieldPaymentReasonCode = "payment_reason_code"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeOriginLink holds the string denoting the originlink edge name in mutations.
	EdgeOriginLink = "originLink"
	// EdgeProductCustomer holds the string denoting the productcustomer edge name in mutations.
	EdgeProductCustomer = "productCustomer"
	// EdgeShop holds the string denoting the shop edge name in mutations.
	EdgeShop = "shop"
	// EdgeProductInfluencer holds the string denoting the productinfluencer edge name in mutations.
	EdgeProductInfluencer = "productInfluencer"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "transaction_product"
	// OriginLinkTable is the table that holds the originLink relation/edge.
	OriginLinkTable = "link_visits"
	// OriginLinkInverseTable is the table name for the LinkVisit entity.
	// It exists in this package in order to avoid circular dependency with the "linkvisit" package.
	OriginLinkInverseTable = "link_visits"
	// OriginLinkColumn is the table column denoting the originLink relation/edge.
	OriginLinkColumn = "transaction_origin_link"
	// ProductCustomerTable is the table that holds the productCustomer relation/edge.
	ProductCustomerTable = "user_buyers"
	// ProductCustomerInverseTable is the table name for the UserBuyer entity.
	// It exists in this package in order to avoid circular dependency with the "userbuyer" package.
	ProductCustomerInverseTable = "user_buyers"
	// ProductCustomerColumn is the table column denoting the productCustomer relation/edge.
	ProductCustomerColumn = "transaction_product_customer"
	// ShopTable is the table that holds the shop relation/edge.
	ShopTable = "shops"
	// ShopInverseTable is the table name for the Shop entity.
	// It exists in this package in order to avoid circular dependency with the "shop" package.
	ShopInverseTable = "shops"
	// ShopColumn is the table column denoting the shop relation/edge.
	ShopColumn = "transaction_shop"
	// ProductInfluencerTable is the table that holds the productInfluencer relation/edge.
	ProductInfluencerTable = "user_influencers"
	// ProductInfluencerInverseTable is the table name for the UserInfluencer entity.
	// It exists in this package in order to avoid circular dependency with the "userinfluencer" package.
	ProductInfluencerInverseTable = "user_influencers"
	// ProductInfluencerColumn is the table column denoting the productInfluencer relation/edge.
	ProductInfluencerColumn = "transaction_product_influencer"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldDateUpdated,
	FieldStatus,
	FieldPaymentMethod,
	FieldPaymentStatus,
	FieldPaymentId,
	FieldPaymentAmount,
	FieldPaymentCurrency,
	FieldPaymentDate,
	FieldPaymentFee,
	FieldPaymentNet,
	FieldPaymentPayerEmail,
	FieldPaymentPayerFirstName,
	FieldPaymentPayerLastName,
	FieldPaymentPayerId,
	FieldPaymentPayerStatus,
	FieldPaymentReceiverEmail,
	FieldPaymentReceiverId,
	FieldPaymentTax,
	FieldPaymentTransactionId,
	FieldPaymentTransactionType,
	FieldPaymentPendingReason,
	FieldPaymentReasonCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_buy_transaction",
	"refund_transactions_transaction",
	"shop_transactions",
	"user_buyer_transactions",
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

// Status defines the type for the "status" enum field.
type Status string

// StatusPENDING is the default value of the Status enum.
const DefaultStatus = StatusPENDING

// Status values.
const (
	StatusPENDING             Status = "PENDING"
	StatusCOMPLETED           Status = "COMPLETED"
	StatusCANCELLED           Status = "CANCELLED"
	StatusREFUNDED            Status = "REFUNDED"
	StatusDENIED              Status = "DENIED"
	StatusFAILED              Status = "FAILED"
	StatusEXPIRED             Status = "EXPIRED"
	StatusVOIDED              Status = "VOIDED"
	StatusREVERSED            Status = "REVERSED"
	StatusPROCESSED           Status = "PROCESSED"
	StatusPARTIALLY_REFUNDED  Status = "PARTIALLY_REFUNDED"
	StatusPARTIALLY_REVERSED  Status = "PARTIALLY_REVERSED"
	StatusPARTIALLY_VOIDED    Status = "PARTIALLY_VOIDED"
	StatusPARTIALLY_PROCESSED Status = "PARTIALLY_PROCESSED"
	StatusPARTIALLY_COMPLETED Status = "PARTIALLY_COMPLETED"
	StatusPARTIALLY_CANCELLED Status = "PARTIALLY_CANCELLED"
	StatusPARTIALLY_DENIED    Status = "PARTIALLY_DENIED"
	StatusPARTIALLY_FAILED    Status = "PARTIALLY_FAILED"
	StatusPARTIALLY_EXPIRED   Status = "PARTIALLY_EXPIRED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPENDING, StatusCOMPLETED, StatusCANCELLED, StatusREFUNDED, StatusDENIED, StatusFAILED, StatusEXPIRED, StatusVOIDED, StatusREVERSED, StatusPROCESSED, StatusPARTIALLY_REFUNDED, StatusPARTIALLY_REVERSED, StatusPARTIALLY_VOIDED, StatusPARTIALLY_PROCESSED, StatusPARTIALLY_COMPLETED, StatusPARTIALLY_CANCELLED, StatusPARTIALLY_DENIED, StatusPARTIALLY_FAILED, StatusPARTIALLY_EXPIRED:
		return nil
	default:
		return fmt.Errorf("transaction: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Transaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByDateUpdated orders the results by the dateUpdated field.
func ByDateUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateUpdated, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPaymentMethod orders the results by the paymentMethod field.
func ByPaymentMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentMethod, opts...).ToFunc()
}

// ByPaymentStatus orders the results by the paymentStatus field.
func ByPaymentStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentStatus, opts...).ToFunc()
}

// ByPaymentId orders the results by the paymentId field.
func ByPaymentId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentId, opts...).ToFunc()
}

// ByPaymentAmount orders the results by the paymentAmount field.
func ByPaymentAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentAmount, opts...).ToFunc()
}

// ByPaymentCurrency orders the results by the paymentCurrency field.
func ByPaymentCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentCurrency, opts...).ToFunc()
}

// ByPaymentDate orders the results by the paymentDate field.
func ByPaymentDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentDate, opts...).ToFunc()
}

// ByPaymentFee orders the results by the paymentFee field.
func ByPaymentFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentFee, opts...).ToFunc()
}

// ByPaymentNet orders the results by the paymentNet field.
func ByPaymentNet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentNet, opts...).ToFunc()
}

// ByPaymentPayerEmail orders the results by the paymentPayerEmail field.
func ByPaymentPayerEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentPayerEmail, opts...).ToFunc()
}

// ByPaymentPayerFirstName orders the results by the paymentPayerFirstName field.
func ByPaymentPayerFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentPayerFirstName, opts...).ToFunc()
}

// ByPaymentPayerLastName orders the results by the paymentPayerLastName field.
func ByPaymentPayerLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentPayerLastName, opts...).ToFunc()
}

// ByPaymentPayerId orders the results by the paymentPayerId field.
func ByPaymentPayerId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentPayerId, opts...).ToFunc()
}

// ByPaymentPayerStatus orders the results by the paymentPayerStatus field.
func ByPaymentPayerStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentPayerStatus, opts...).ToFunc()
}

// ByPaymentReceiverEmail orders the results by the paymentReceiverEmail field.
func ByPaymentReceiverEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentReceiverEmail, opts...).ToFunc()
}

// ByPaymentReceiverId orders the results by the paymentReceiverId field.
func ByPaymentReceiverId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentReceiverId, opts...).ToFunc()
}

// ByPaymentTax orders the results by the paymentTax field.
func ByPaymentTax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentTax, opts...).ToFunc()
}

// ByPaymentTransactionId orders the results by the paymentTransactionId field.
func ByPaymentTransactionId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentTransactionId, opts...).ToFunc()
}

// ByPaymentTransactionType orders the results by the paymentTransactionType field.
func ByPaymentTransactionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentTransactionType, opts...).ToFunc()
}

// ByPaymentPendingReason orders the results by the paymentPendingReason field.
func ByPaymentPendingReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentPendingReason, opts...).ToFunc()
}

// ByPaymentReasonCode orders the results by the paymentReasonCode field.
func ByPaymentReasonCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentReasonCode, opts...).ToFunc()
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

// ByOriginLinkCount orders the results by originLink count.
func ByOriginLinkCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOriginLinkStep(), opts...)
	}
}

// ByOriginLink orders the results by originLink terms.
func ByOriginLink(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOriginLinkStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByShopCount orders the results by shop count.
func ByShopCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newShopStep(), opts...)
	}
}

// ByShop orders the results by shop terms.
func ByShop(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShopStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductInfluencerCount orders the results by productInfluencer count.
func ByProductInfluencerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductInfluencerStep(), opts...)
	}
}

// ByProductInfluencer orders the results by productInfluencer terms.
func ByProductInfluencer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductInfluencerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
	)
}
func newOriginLinkStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OriginLinkInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OriginLinkTable, OriginLinkColumn),
	)
}
func newProductCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductCustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductCustomerTable, ProductCustomerColumn),
	)
}
func newShopStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShopInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ShopTable, ShopColumn),
	)
}
func newProductInfluencerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInfluencerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductInfluencerTable, ProductInfluencerColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}