package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserBuyer holds the schema definition for the UserBuyer entity.
type UserBuyer struct {
	ent.Schema
}

// TODO FIX
func (UserBuyer) Fields() []ent.Field {
	// return []ent.Field{}
	return []ent.Field{
		field.Int("placeholder").Optional(),
	}
}

// Edges of the UserBuyer.
func (UserBuyer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userProfile", User.Type),
		edge.To("reviews", Review.Type),
		edge.To("transactions", Transaction.Type),
		edge.To("linksClicked", LinkVisit.Type),
	}
}

func (UserBuyer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
