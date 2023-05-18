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
		field.Int("name"),
		field.Int("description"),
		field.Int("value"),
	}
}

func (ProductAttribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
