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
		field.String("dateCreated"),
		field.Int("productPrice"),
		field.Int("moq"),
		field.String("startDate"),
		field.String("endDate"),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
