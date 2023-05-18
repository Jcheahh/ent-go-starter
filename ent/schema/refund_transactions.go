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
		field.String("refundAmount"),
		field.String("refundCurrency"),
		field.String("refundReason"),
		field.String("refundStatus"),
		field.String("dateCreated"),
		field.String("dateUpdated"),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
