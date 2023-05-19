package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("email").Annotations(
			entgql.OrderField("EMAIL"),
		),
		field.String("phone").Annotations(
			entgql.OrderField("PHONE"),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notifications", Notification.Type),
		edge.To("bankAccounts", BankAccount.Type),
		edge.To("shippingAddresses", ShippingAddress.Type),
		edge.To("paymentMethods", PaymentMethod.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
