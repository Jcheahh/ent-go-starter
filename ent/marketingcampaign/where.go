// Code generated by ent, DO NOT EDIT.

package marketingcampaign

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldDescription, v))
}

// ConsumerPurchaseValue applies equality check predicate on the "consumerPurchaseValue" field. It's identical to ConsumerPurchaseValueEQ.
func ConsumerPurchaseValue(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldConsumerPurchaseValue, v))
}

// CustomerApplicationLogic applies equality check predicate on the "customerApplicationLogic" field. It's identical to CustomerApplicationLogicEQ.
func CustomerApplicationLogic(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldCustomerApplicationLogic, v))
}

// InitialisationLogic applies equality check predicate on the "initialisationLogic" field. It's identical to InitialisationLogicEQ.
func InitialisationLogic(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldInitialisationLogic, v))
}

// StartDate applies equality check predicate on the "startDate" field. It's identical to StartDateEQ.
func StartDate(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "endDate" field. It's identical to EndDateEQ.
func EndDate(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldEndDate, v))
}

// DateCreated applies equality check predicate on the "dateCreated" field. It's identical to DateCreatedEQ.
func DateCreated(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldDateCreated, v))
}

// DateUpdated applies equality check predicate on the "dateUpdated" field. It's identical to DateUpdatedEQ.
func DateUpdated(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldDateUpdated, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldDescription, v))
}

// ConsumerPurchaseValueEQ applies the EQ predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueNEQ applies the NEQ predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueIn applies the In predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldConsumerPurchaseValue, vs...))
}

// ConsumerPurchaseValueNotIn applies the NotIn predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldConsumerPurchaseValue, vs...))
}

// ConsumerPurchaseValueGT applies the GT predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueGTE applies the GTE predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueLT applies the LT predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueLTE applies the LTE predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueContains applies the Contains predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueHasPrefix applies the HasPrefix predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueHasSuffix applies the HasSuffix predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueEqualFold applies the EqualFold predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldConsumerPurchaseValue, v))
}

// ConsumerPurchaseValueContainsFold applies the ContainsFold predicate on the "consumerPurchaseValue" field.
func ConsumerPurchaseValueContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldConsumerPurchaseValue, v))
}

// CustomerApplicationLogicEQ applies the EQ predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicNEQ applies the NEQ predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicIn applies the In predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldCustomerApplicationLogic, vs...))
}

// CustomerApplicationLogicNotIn applies the NotIn predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldCustomerApplicationLogic, vs...))
}

// CustomerApplicationLogicGT applies the GT predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicGTE applies the GTE predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicLT applies the LT predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicLTE applies the LTE predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicContains applies the Contains predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicHasPrefix applies the HasPrefix predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicHasSuffix applies the HasSuffix predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicEqualFold applies the EqualFold predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldCustomerApplicationLogic, v))
}

// CustomerApplicationLogicContainsFold applies the ContainsFold predicate on the "customerApplicationLogic" field.
func CustomerApplicationLogicContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldCustomerApplicationLogic, v))
}

// InitialisationLogicEQ applies the EQ predicate on the "initialisationLogic" field.
func InitialisationLogicEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldInitialisationLogic, v))
}

// InitialisationLogicNEQ applies the NEQ predicate on the "initialisationLogic" field.
func InitialisationLogicNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldInitialisationLogic, v))
}

// InitialisationLogicIn applies the In predicate on the "initialisationLogic" field.
func InitialisationLogicIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldInitialisationLogic, vs...))
}

// InitialisationLogicNotIn applies the NotIn predicate on the "initialisationLogic" field.
func InitialisationLogicNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldInitialisationLogic, vs...))
}

// InitialisationLogicGT applies the GT predicate on the "initialisationLogic" field.
func InitialisationLogicGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldInitialisationLogic, v))
}

// InitialisationLogicGTE applies the GTE predicate on the "initialisationLogic" field.
func InitialisationLogicGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldInitialisationLogic, v))
}

// InitialisationLogicLT applies the LT predicate on the "initialisationLogic" field.
func InitialisationLogicLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldInitialisationLogic, v))
}

// InitialisationLogicLTE applies the LTE predicate on the "initialisationLogic" field.
func InitialisationLogicLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldInitialisationLogic, v))
}

// InitialisationLogicContains applies the Contains predicate on the "initialisationLogic" field.
func InitialisationLogicContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldInitialisationLogic, v))
}

// InitialisationLogicHasPrefix applies the HasPrefix predicate on the "initialisationLogic" field.
func InitialisationLogicHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldInitialisationLogic, v))
}

// InitialisationLogicHasSuffix applies the HasSuffix predicate on the "initialisationLogic" field.
func InitialisationLogicHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldInitialisationLogic, v))
}

// InitialisationLogicEqualFold applies the EqualFold predicate on the "initialisationLogic" field.
func InitialisationLogicEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldInitialisationLogic, v))
}

// InitialisationLogicContainsFold applies the ContainsFold predicate on the "initialisationLogic" field.
func InitialisationLogicContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldInitialisationLogic, v))
}

// StartDateEQ applies the EQ predicate on the "startDate" field.
func StartDateEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "startDate" field.
func StartDateNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "startDate" field.
func StartDateIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "startDate" field.
func StartDateNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "startDate" field.
func StartDateGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "startDate" field.
func StartDateGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "startDate" field.
func StartDateLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "startDate" field.
func StartDateLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldStartDate, v))
}

// StartDateContains applies the Contains predicate on the "startDate" field.
func StartDateContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldStartDate, v))
}

// StartDateHasPrefix applies the HasPrefix predicate on the "startDate" field.
func StartDateHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldStartDate, v))
}

// StartDateHasSuffix applies the HasSuffix predicate on the "startDate" field.
func StartDateHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldStartDate, v))
}

// StartDateEqualFold applies the EqualFold predicate on the "startDate" field.
func StartDateEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldStartDate, v))
}

// StartDateContainsFold applies the ContainsFold predicate on the "startDate" field.
func StartDateContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldStartDate, v))
}

// EndDateEQ applies the EQ predicate on the "endDate" field.
func EndDateEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "endDate" field.
func EndDateNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "endDate" field.
func EndDateIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "endDate" field.
func EndDateNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "endDate" field.
func EndDateGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "endDate" field.
func EndDateGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "endDate" field.
func EndDateLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "endDate" field.
func EndDateLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldEndDate, v))
}

// EndDateContains applies the Contains predicate on the "endDate" field.
func EndDateContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldEndDate, v))
}

// EndDateHasPrefix applies the HasPrefix predicate on the "endDate" field.
func EndDateHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldEndDate, v))
}

// EndDateHasSuffix applies the HasSuffix predicate on the "endDate" field.
func EndDateHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldEndDate, v))
}

// EndDateEqualFold applies the EqualFold predicate on the "endDate" field.
func EndDateEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldEndDate, v))
}

// EndDateContainsFold applies the ContainsFold predicate on the "endDate" field.
func EndDateContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldEndDate, v))
}

// DateCreatedEQ applies the EQ predicate on the "dateCreated" field.
func DateCreatedEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "dateCreated" field.
func DateCreatedNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "dateCreated" field.
func DateCreatedIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "dateCreated" field.
func DateCreatedNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "dateCreated" field.
func DateCreatedGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "dateCreated" field.
func DateCreatedGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "dateCreated" field.
func DateCreatedLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "dateCreated" field.
func DateCreatedLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldDateCreated, v))
}

// DateCreatedContains applies the Contains predicate on the "dateCreated" field.
func DateCreatedContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldDateCreated, v))
}

// DateCreatedHasPrefix applies the HasPrefix predicate on the "dateCreated" field.
func DateCreatedHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldDateCreated, v))
}

// DateCreatedHasSuffix applies the HasSuffix predicate on the "dateCreated" field.
func DateCreatedHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldDateCreated, v))
}

// DateCreatedEqualFold applies the EqualFold predicate on the "dateCreated" field.
func DateCreatedEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldDateCreated, v))
}

// DateCreatedContainsFold applies the ContainsFold predicate on the "dateCreated" field.
func DateCreatedContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldDateCreated, v))
}

// DateUpdatedEQ applies the EQ predicate on the "dateUpdated" field.
func DateUpdatedEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEQ(FieldDateUpdated, v))
}

// DateUpdatedNEQ applies the NEQ predicate on the "dateUpdated" field.
func DateUpdatedNEQ(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNEQ(FieldDateUpdated, v))
}

// DateUpdatedIn applies the In predicate on the "dateUpdated" field.
func DateUpdatedIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldIn(FieldDateUpdated, vs...))
}

// DateUpdatedNotIn applies the NotIn predicate on the "dateUpdated" field.
func DateUpdatedNotIn(vs ...string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldNotIn(FieldDateUpdated, vs...))
}

// DateUpdatedGT applies the GT predicate on the "dateUpdated" field.
func DateUpdatedGT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGT(FieldDateUpdated, v))
}

// DateUpdatedGTE applies the GTE predicate on the "dateUpdated" field.
func DateUpdatedGTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldGTE(FieldDateUpdated, v))
}

// DateUpdatedLT applies the LT predicate on the "dateUpdated" field.
func DateUpdatedLT(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLT(FieldDateUpdated, v))
}

// DateUpdatedLTE applies the LTE predicate on the "dateUpdated" field.
func DateUpdatedLTE(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldLTE(FieldDateUpdated, v))
}

// DateUpdatedContains applies the Contains predicate on the "dateUpdated" field.
func DateUpdatedContains(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContains(FieldDateUpdated, v))
}

// DateUpdatedHasPrefix applies the HasPrefix predicate on the "dateUpdated" field.
func DateUpdatedHasPrefix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasPrefix(FieldDateUpdated, v))
}

// DateUpdatedHasSuffix applies the HasSuffix predicate on the "dateUpdated" field.
func DateUpdatedHasSuffix(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldHasSuffix(FieldDateUpdated, v))
}

// DateUpdatedEqualFold applies the EqualFold predicate on the "dateUpdated" field.
func DateUpdatedEqualFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldEqualFold(FieldDateUpdated, v))
}

// DateUpdatedContainsFold applies the ContainsFold predicate on the "dateUpdated" field.
func DateUpdatedContainsFold(v string) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(sql.FieldContainsFold(FieldDateUpdated, v))
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConsumerReward applies the HasEdge predicate on the "consumerReward" edge.
func HasConsumerReward() predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConsumerRewardTable, ConsumerRewardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConsumerRewardWith applies the HasEdge predicate on the "consumerReward" edge with a given conditions (other predicates).
func HasConsumerRewardWith(preds ...predicate.RewardType) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		step := newConsumerRewardStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MarketingCampaign) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MarketingCampaign) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MarketingCampaign) predicate.MarketingCampaign {
	return predicate.MarketingCampaign(func(s *sql.Selector) {
		p(s.Not())
	})
}
