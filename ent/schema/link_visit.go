package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type LinkVisit struct {
	ent.Schema
}

func (LinkVisit) Fields() []ent.Field {
	return []ent.Field{
		field.String("dateCreated"),
		field.String("ipAddress"),
		field.Int("saleValue"),
		field.Int("commissionEarned"),
	}
}

func (LinkVisit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
