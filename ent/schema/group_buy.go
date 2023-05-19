package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type GroupBuy struct {
	ent.Schema
}

func (GroupBuy) Fields() []ent.Field {
	return []ent.Field{
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.Int("productPrice").Annotations(
			entgql.OrderField("PRODUCTPRICE"),
		),
		field.Int("moq").Annotations(
			entgql.OrderField("MOQ"),
		),
		field.String("startDate").Annotations(
			entgql.OrderField("STARTDATE"),
		),
		field.String("endDate").Annotations(
			entgql.OrderField("ENDDATE"),
		),
	}
}

func (GroupBuy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
		edge.To("transaction", Transaction.Type),
	}
}

func (GroupBuy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
