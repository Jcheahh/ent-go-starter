package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
