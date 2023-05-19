package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Chat struct {
	ent.Schema
}

func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("xid").Annotations(
			entgql.OrderField("XID"),
		),
	}
}

func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
