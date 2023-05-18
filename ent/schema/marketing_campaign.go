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
		field.String("name"),
		field.String("description"),
		field.String("consumerPurchaseValue"),
		// TODO put logic
		// (age > 25 & age < 30) & address == "Kuala Lumpur" else can be niil. This logic is applied to the customer
		field.String("customerApplicationLogic"),
		// # productSales > 20 & productReviews > 100, then start the campaign
		field.String("initialisationLogic"),
		field.String("startDate"),
		field.String("endDate"),
		field.String("dateCreated"),
		field.String("dateUpdated"),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
