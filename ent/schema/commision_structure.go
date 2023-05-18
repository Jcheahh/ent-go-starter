package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CommissionStructure struct {
	ent.Schema
}

func (CommissionStructure) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("commissionValue"),
		field.String("commissionPercentage"),
	}
}

func (CommissionStructure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("productSeller", UserSeller.Type),
	}
}

func (CommissionStructure) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
