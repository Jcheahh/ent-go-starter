package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ReferralLink struct {
	ent.Schema
}

func (ReferralLink) Fields() []ent.Field {
	return []ent.Field{
		field.Int("xid").Annotations(
			entgql.OrderField("XID"),
		),
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
		field.String("link").Annotations(
			entgql.OrderField("LINK"),
		),
	}
}

func (ReferralLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("visits", LinkVisit.Type),
	}
}

func (ReferralLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
