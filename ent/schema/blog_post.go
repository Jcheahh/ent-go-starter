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
		field.String("title").Annotations(
			entgql.OrderField("TITLE"),
		),
		field.String("content").Annotations(
			entgql.OrderField("CONTENT"),
		),
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.String("dateUpdated").Annotations(
			entgql.OrderField("DATEUPDATED"),
		),
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
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
