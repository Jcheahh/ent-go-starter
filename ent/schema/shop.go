package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Shop holds the schema definition for the Shop entity.
type Shop struct {
	ent.Schema
}

// Fields of the Shop.
func (Shop) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
	}
}

// Edges of the Shop.
func (Shop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("transactions", Transaction.Type),
		edge.To("bankAccounts", BankAccount.Type),
	}
}

func (Shop) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
