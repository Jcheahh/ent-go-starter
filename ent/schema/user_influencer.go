package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserInfluencer holds the schema definition for the UserInfluencer entity.
type UserInfluencer struct {
	ent.Schema
}

// TODO FIX
func (UserInfluencer) Fields() []ent.Field {
	// return []ent.Field{}
	return []ent.Field{
		field.Int("placeholder").Optional(),
	}
}

// Edges of the UserInfluencer.
func (UserInfluencer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userProfile", User.Type),
		edge.To("referralLinks", ReferralLink.Type),
		edge.To("reviews", Review.Type),
		edge.To("products", Product.Type),
		edge.To("tags", Tag.Type),
	}
}

func (UserInfluencer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
