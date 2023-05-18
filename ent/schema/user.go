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
		field.String("name"),
		field.String("email"),
		field.String("phone"),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
