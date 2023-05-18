package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PrimaryContent struct {
	ent.Schema
}

// TODO FIX
func (PrimaryContent) Fields() []ent.Field {
	// return nil
	return []ent.Field{
		field.Int("placeholder").Optional(),
	}
}

func (PrimaryContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contentBlock", ContentBlock.Type),
	}
}

func (PrimaryContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
