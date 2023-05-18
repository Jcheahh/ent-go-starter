package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ViewAnalytics struct {
	ent.Schema
}

func (ViewAnalytics) Fields() []ent.Field {
	return []ent.Field{
		field.Int("views"),
		field.Int("scrolls"),
		field.Int("exits"),
		field.String("dateCreated"),
	}
}

func (ViewAnalytics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
	}
}

func (ViewAnalytics) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
