package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ProductVariation struct {
	ent.Schema
}

func (ProductVariation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
		field.String("price").Annotations(
			entgql.OrderField("PRICE"),
		),
	}
}

func (ProductVariation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("productAttributes", ProductAttribute.Type),
	}
}

func (ProductVariation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
