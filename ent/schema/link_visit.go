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
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.String("ipAddress").Annotations(
			entgql.OrderField("IPADDRESS"),
		),
		field.Int("saleValue").Annotations(
			entgql.OrderField("SALEVALUE"),
		),
		field.Int("commissionEarned").Annotations(
			entgql.OrderField("COMMISSIONEARNED"),
		),
	}
}

func (LinkVisit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
