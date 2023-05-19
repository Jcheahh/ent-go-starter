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
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("address").Annotations(
			entgql.OrderField("ADDRESS"),
		),
		field.String("city").Annotations(
			entgql.OrderField("CITY"),
		),
		field.String("state").Annotations(
			entgql.OrderField("STATE"),
		),
		field.String("zip").Annotations(
			entgql.OrderField("ZIP"),
		),
		field.String("country").Annotations(
			entgql.OrderField("COUNTRY"),
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
	}
}

func (ShippingAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
