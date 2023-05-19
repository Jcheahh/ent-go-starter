package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type ProductAttribute struct {
	ent.Schema
}

func (ProductAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.Int("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.Int("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
		field.Int("value").Annotations(
			entgql.OrderField("VALUE"),
		),
	}
}

func (ProductAttribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
