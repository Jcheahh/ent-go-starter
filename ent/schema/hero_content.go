package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HeroContent struct {
	ent.Schema
}

func (HeroContent) Fields() []ent.Field {
	return []ent.Field{
		field.String("primaryMessage").Annotations(
			entgql.OrderField("PRIMARYMESSAGE"),
		),
		field.String("secondaryMessage").Annotations(
			entgql.OrderField("SECONDARYMESSAGE"),
		),
	}
}

func (HeroContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type),
	}
}

func (HeroContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
