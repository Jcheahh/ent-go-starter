package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ContentBlock struct {
	ent.Schema
}

func (ContentBlock) Fields() []ent.Field {
	return []ent.Field{
		field.String("primaryMessage"),
		field.String("secondaryMessage"),
	}
}

func (ContentBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type),
	}
}

func (ContentBlock) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
