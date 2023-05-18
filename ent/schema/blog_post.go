package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type BlogPost struct {
	ent.Schema
}

func (BlogPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.String("dateCreated"),
		field.String("dateUpdated"),
	}
}

func (BlogPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", UserSeller.Type),
	}
}

func (BlogPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
