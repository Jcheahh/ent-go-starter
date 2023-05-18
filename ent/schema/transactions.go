package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("dateCreated"),
		field.String("dateUpdated"),
		field.Enum("status").
			Values("PENDING", "COMPLETED", "CANCELLED", "REFUNDED", "DENIED", "FAILED", "EXPIRED", "VOIDED", "REVERSED", "PROCESSED", "PARTIALLY_REFUNDED", "PARTIALLY_REVERSED", "PARTIALLY_VOIDED", "PARTIALLY_PROCESSED", "PARTIALLY_COMPLETED", "PARTIALLY_CANCELLED", "PARTIALLY_DENIED", "PARTIALLY_FAILED", "PARTIALLY_EXPIRED").
			Default("PENDING"),
		field.String("paymentMethod"),
		field.String("paymentStatus"),
		field.String("paymentId"),
		field.String("paymentAmount"),
		field.String("paymentCurrency"),
		field.String("paymentDate"),
		field.String("paymentFee"),
		field.String("paymentNet"),
		field.String("paymentPayerEmail"),
		field.String("paymentPayerFirstName"),
		field.String("paymentPayerLastName"),
		field.String("paymentPayerId"),
		field.String("paymentPayerStatus"),
		field.String("paymentReceiverEmail"),
		field.String("paymentReceiverId"),
		field.String("paymentTax"),
		field.String("paymentTransactionId"),
		field.String("paymentTransactionType"),
		field.String("paymentPendingReason"),
		field.String("paymentReasonCode"),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
		edge.To("originLink", LinkVisit.Type),
		edge.To("productCustomer", UserBuyer.Type),
		edge.To("shop", Shop.Type),
		edge.To("productInfluencer", UserInfluencer.Type),
	}
}

func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
