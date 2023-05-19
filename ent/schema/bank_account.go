package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type BankAccount struct {
	ent.Schema
}

// Fields of the BankAccount.
func (BankAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("xid").Annotations(
			entgql.OrderField("XID"),
		),
	}
}

func (BankAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
