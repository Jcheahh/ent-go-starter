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
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.String("dateUpdated").Annotations(
			entgql.OrderField("DATEUPDATED"),
		),
		field.Enum("status").Annotations(
			entgql.OrderField("STATUS"),
		).
			Values("PENDING", "COMPLETED", "CANCELLED", "REFUNDED", "DENIED", "FAILED", "EXPIRED", "VOIDED", "REVERSED", "PROCESSED", "PARTIALLY_REFUNDED", "PARTIALLY_REVERSED", "PARTIALLY_VOIDED", "PARTIALLY_PROCESSED", "PARTIALLY_COMPLETED", "PARTIALLY_CANCELLED", "PARTIALLY_DENIED", "PARTIALLY_FAILED", "PARTIALLY_EXPIRED").
			Default("PENDING"),
		field.String("paymentMethod").Annotations(
			entgql.OrderField("PAYMENTMETHOD"),
		),
		field.String("paymentStatus").Annotations(
			entgql.OrderField("PAYMENTSTATUS"),
		),
		field.String("paymentId").Annotations(
			entgql.OrderField("PAYMENTID"),
		),
		field.String("paymentAmount").Annotations(
			entgql.OrderField("PAYMENTAMOUNT"),
		),
		field.String("paymentCurrency").Annotations(
			entgql.OrderField("PAYMENTCURRENCY"),
		),
		field.String("paymentDate").Annotations(
			entgql.OrderField("PAYMENTDATE"),
		),
		field.String("paymentFee").Annotations(
			entgql.OrderField("PAYMENTFEE"),
		),
		field.String("paymentNet").Annotations(
			entgql.OrderField("PAYMENTNET"),
		),
		field.String("paymentPayerEmail").Annotations(
			entgql.OrderField("PAYMENTPAYEREMAIL"),
		),
		field.String("paymentPayerFirstName").Annotations(
			entgql.OrderField("PAYMENTPAYERFIRSTNAME"),
		),
		field.String("paymentPayerLastName").Annotations(
			entgql.OrderField("PAYMENTPAYERLASTNAME"),
		),
		field.String("paymentPayerId").Annotations(
			entgql.OrderField("PAYMENTPAYERID"),
		),
		field.String("paymentPayerStatus").Annotations(
			entgql.OrderField("PAYMENTPAYERSTATUS"),
		),
		field.String("paymentReceiverEmail").Annotations(
			entgql.OrderField("PAYMENTRECEIVEREMAIL"),
		),
		field.String("paymentReceiverId").Annotations(
			entgql.OrderField("PAYMENTRECEIVERID"),
		),
		field.String("paymentTax").Annotations(
			entgql.OrderField("PAYMENTTAX"),
		),
		field.String("paymentTransactionId").Annotations(
			entgql.OrderField("PAYMENTTRANSACTIONID"),
		),
		field.String("paymentTransactionType").Annotations(
			entgql.OrderField("PAYMENTTRANSACTIONTYPE"),
		),
		field.String("paymentPendingReason").Annotations(
			entgql.OrderField("PAYMENTPENDINGREASON"),
		),
		field.String("paymentReasonCode").Annotations(
			entgql.OrderField("PAYMENTREASONCODE"),
		),
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
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
