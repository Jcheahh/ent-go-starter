package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ShippingAddress holds the schema definition for the ShippingAddress entity.
type ShippingAddress struct {
	ent.Schema
}

// Fields of the ShippingAddress.
func (ShippingAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("address"),
		field.String("city"),
		field.String("state"),
		field.String("zip"),
		field.String("country"),
		field.String("dateCreated").
			Immutable().
			Default(time.Now().Format(time.RFC3339)),
		field.String("dateUpdated").
			Default(time.Now().Format(time.RFC3339)),
	}
}

func (ShippingAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
