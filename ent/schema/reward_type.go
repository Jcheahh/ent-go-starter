package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RewardType holds the schema definition for the RewardType entity.
type RewardType struct {
	ent.Schema
}

// Fields of the RewardType.
func (RewardType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values(
				"DISCOUNT",
				"FREE_PRODUCT",
				"SHIPPING_DISCOUNT",
			),
		field.Int("val"),
	}
}

// Edges of the RewardType.
func (RewardType) Edges() []ent.Edge {
	return nil
}

func (RewardType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
