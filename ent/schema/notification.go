package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Annotations(
			entgql.OrderField("TITLE"),
		),
		field.String("content").Annotations(
			entgql.OrderField("CONTENT"),
		),
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		).
			Immutable().
			Default(time.Now().Format(time.RFC3339)),
		field.String("dateUpdated").Annotations(
			entgql.OrderField("DATEUPDATED"),
		).
			Default(time.Now().Format(time.RFC3339)),
		field.Bool("read").Annotations(
			entgql.OrderField("READ"),
		),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipient", User.Type),
	}
}

func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
