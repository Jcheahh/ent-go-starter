package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type RefundTransactions struct {
	ent.Schema
}

func (RefundTransactions) Fields() []ent.Field {
	return []ent.Field{
		field.String("refundAmount").Annotations(
			entgql.OrderField("REFUNDAMOUNT"),
		),
		field.String("refundCurrency").Annotations(
			entgql.OrderField("REFUNDCURRENCY"),
		),
		field.String("refundReason").Annotations(
			entgql.OrderField("REFUNDREASON"),
		),
		field.String("refundStatus").Annotations(
			entgql.OrderField("REFUNDSTATUS"),
		),
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.String("dateUpdated").Annotations(
			entgql.OrderField("DATEUPDATED"),
		),
	}
}

func (RefundTransactions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transaction", Transaction.Type),
	}
}

func (RefundTransactions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
