package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ProductPageView struct {
	ent.Schema
}

func (ProductPageView) Fields() []ent.Field {
	return []ent.Field{
		field.Int("version"),
	}
}

func (ProductPageView) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("heroContent", HeroContent.Type),
		edge.To("primaryContent", PrimaryContent.Type),
		edge.To("viewAnalytics", ViewAnalytics.Type),
	}
}

func (ProductPageView) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
