package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CommissionStructureSchema struct {
	ent.Schema
}

func (CommissionStructureSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
		field.String("commissionValue").Annotations(
			entgql.OrderField("COMMISSIONVALUE"),
		),
		field.String("commissionPercentage").Annotations(
			entgql.OrderField("COMMISSIONPERCENTAGE"),
		),
	}
}

func (CommissionStructureSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("productSeller", UserSeller.Type),
	}
}

func (CommissionStructureSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
