package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserSeller holds the schema definition for the UserSeller entity.
type UserSeller struct {
	ent.Schema
}

// Fields of the UserSeller.
func (UserSeller) Fields() []ent.Field {
	return []ent.Field{
		field.String("brandName").Annotations(
			entgql.OrderField("BRANDNAME"),
		),
	}
}

// Edges of the UserSeller.
func (UserSeller) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userProfile", User.Type),
		edge.To("shops", Shop.Type),
	}
}

func (UserSeller) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
