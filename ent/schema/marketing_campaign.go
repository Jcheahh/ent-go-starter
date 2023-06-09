package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MarketingCampaign struct {
	ent.Schema
}

func (MarketingCampaign) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
		field.String("consumerPurchaseValue").Annotations(
			entgql.OrderField("CONSUMERPURCHASEVALUE"),
		),
		// TODO put logic
		// (age > 25 & age < 30) & address == "Kuala Lumpur" else can be niil. This logic is applied to the customer
		field.String("customerApplicationLogic").Annotations(
			entgql.OrderField("CUSTOMERAPPLICATIONLOGIC"),
		),
		// # productSales > 20 & productReviews > 100, then start the campaign
		field.String("initialisationLogic").Annotations(
			entgql.OrderField("INITIALISATIONLOGIC"),
		),
		field.String("startDate").Annotations(
			entgql.OrderField("STARTDATE"),
		),
		field.String("endDate").Annotations(
			entgql.OrderField("ENDDATE"),
		),
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.String("dateUpdated").Annotations(
			entgql.OrderField("DATEUPDATED"),
		),
	}
}

func (MarketingCampaign) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).Required(),
		edge.To("consumerReward", RewardType.Type),
	}
}

func (MarketingCampaign) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
