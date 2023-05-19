package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Review struct {
	ent.Schema
}

func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Annotations(
			entgql.OrderField("TYPE"),
		),
		field.String("content").Annotations(
			entgql.OrderField("CONTENT"),
		),
		field.String("rating").Annotations(
			entgql.OrderField("RATING"),
		),
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
	}
}

func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
		edge.To("productCustomer", UserBuyer.Type),
	}
}

func (Review) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
