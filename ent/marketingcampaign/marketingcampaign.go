// Code generated by ent, DO NOT EDIT.

package marketingcampaign

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the marketingcampaign type in the database.
	Label = "marketing_campaign"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldConsumerPurchaseValue holds the string denoting the consumerpurchasevalue field in the database.
	FieldConsumerPurchaseValue = "consumer_purchase_value"
	// FieldCustomerApplicationLogic holds the string denoting the customerapplicationlogic field in the database.
	FieldCustomerApplicationLogic = "customer_application_logic"
	// FieldInitialisationLogic holds the string denoting the initialisationlogic field in the database.
	FieldInitialisationLogic = "initialisation_logic"
	// FieldStartDate holds the string denoting the startdate field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the enddate field in the database.
	FieldEndDate = "end_date"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldDateUpdated holds the string denoting the dateupdated field in the database.
	FieldDateUpdated = "date_updated"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeConsumerReward holds the string denoting the consumerreward edge name in mutations.
	EdgeConsumerReward = "consumerReward"
	// Table holds the table name of the marketingcampaign in the database.
	Table = "marketing_campaigns"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "marketing_campaign_product"
	// ConsumerRewardTable is the table that holds the consumerReward relation/edge.
	ConsumerRewardTable = "reward_types"
	// ConsumerRewardInverseTable is the table name for the RewardType entity.
	// It exists in this package in order to avoid circular dependency with the "rewardtype" package.
	ConsumerRewardInverseTable = "reward_types"
	// ConsumerRewardColumn is the table column denoting the consumerReward relation/edge.
	ConsumerRewardColumn = "marketing_campaign_consumer_reward"
)

// Columns holds all SQL columns for marketingcampaign fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldConsumerPurchaseValue,
	FieldCustomerApplicationLogic,
	FieldInitialisationLogic,
	FieldStartDate,
	FieldEndDate,
	FieldDateCreated,
	FieldDateUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "marketing_campaigns"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_marketing_campaigns",
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

// OrderOption defines the ordering options for the MarketingCampaign queries.
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

// ByConsumerPurchaseValue orders the results by the consumerPurchaseValue field.
func ByConsumerPurchaseValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConsumerPurchaseValue, opts...).ToFunc()
}

// ByCustomerApplicationLogic orders the results by the customerApplicationLogic field.
func ByCustomerApplicationLogic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerApplicationLogic, opts...).ToFunc()
}

// ByInitialisationLogic orders the results by the initialisationLogic field.
func ByInitialisationLogic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInitialisationLogic, opts...).ToFunc()
}

// ByStartDate orders the results by the startDate field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the endDate field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByDateCreated orders the results by the dateCreated field.
func ByDateCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateCreated, opts...).ToFunc()
}

// ByDateUpdated orders the results by the dateUpdated field.
func ByDateUpdated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateUpdated, opts...).ToFunc()
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

// ByConsumerRewardCount orders the results by consumerReward count.
func ByConsumerRewardCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newConsumerRewardStep(), opts...)
	}
}

// ByConsumerReward orders the results by consumerReward terms.
func ByConsumerReward(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConsumerRewardStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
	)
}
func newConsumerRewardStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConsumerRewardInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ConsumerRewardTable, ConsumerRewardColumn),
	)
}
