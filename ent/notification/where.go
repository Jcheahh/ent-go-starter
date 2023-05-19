// Code generated by ent, DO NOT EDIT.

package notification

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldContent, v))
}

// DateCreated applies equality check predicate on the "dateCreated" field. It's identical to DateCreatedEQ.
func DateCreated(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldDateCreated, v))
}

// DateUpdated applies equality check predicate on the "dateUpdated" field. It's identical to DateUpdatedEQ.
func DateUpdated(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldDateUpdated, v))
}

// Read applies equality check predicate on the "read" field. It's identical to ReadEQ.
func Read(v bool) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldRead, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldContent, v))
}

// DateCreatedEQ applies the EQ predicate on the "dateCreated" field.
func DateCreatedEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "dateCreated" field.
func DateCreatedNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "dateCreated" field.
func DateCreatedIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "dateCreated" field.
func DateCreatedNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "dateCreated" field.
func DateCreatedGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "dateCreated" field.
func DateCreatedGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "dateCreated" field.
func DateCreatedLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "dateCreated" field.
func DateCreatedLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldDateCreated, v))
}

// DateCreatedContains applies the Contains predicate on the "dateCreated" field.
func DateCreatedContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldDateCreated, v))
}

// DateCreatedHasPrefix applies the HasPrefix predicate on the "dateCreated" field.
func DateCreatedHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldDateCreated, v))
}

// DateCreatedHasSuffix applies the HasSuffix predicate on the "dateCreated" field.
func DateCreatedHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldDateCreated, v))
}

// DateCreatedEqualFold applies the EqualFold predicate on the "dateCreated" field.
func DateCreatedEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldDateCreated, v))
}

// DateCreatedContainsFold applies the ContainsFold predicate on the "dateCreated" field.
func DateCreatedContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldDateCreated, v))
}

// DateUpdatedEQ applies the EQ predicate on the "dateUpdated" field.
func DateUpdatedEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldDateUpdated, v))
}

// DateUpdatedNEQ applies the NEQ predicate on the "dateUpdated" field.
func DateUpdatedNEQ(v string) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldDateUpdated, v))
}

// DateUpdatedIn applies the In predicate on the "dateUpdated" field.
func DateUpdatedIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldIn(FieldDateUpdated, vs...))
}

// DateUpdatedNotIn applies the NotIn predicate on the "dateUpdated" field.
func DateUpdatedNotIn(vs ...string) predicate.Notification {
	return predicate.Notification(sql.FieldNotIn(FieldDateUpdated, vs...))
}

// DateUpdatedGT applies the GT predicate on the "dateUpdated" field.
func DateUpdatedGT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGT(FieldDateUpdated, v))
}

// DateUpdatedGTE applies the GTE predicate on the "dateUpdated" field.
func DateUpdatedGTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldGTE(FieldDateUpdated, v))
}

// DateUpdatedLT applies the LT predicate on the "dateUpdated" field.
func DateUpdatedLT(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLT(FieldDateUpdated, v))
}

// DateUpdatedLTE applies the LTE predicate on the "dateUpdated" field.
func DateUpdatedLTE(v string) predicate.Notification {
	return predicate.Notification(sql.FieldLTE(FieldDateUpdated, v))
}

// DateUpdatedContains applies the Contains predicate on the "dateUpdated" field.
func DateUpdatedContains(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContains(FieldDateUpdated, v))
}

// DateUpdatedHasPrefix applies the HasPrefix predicate on the "dateUpdated" field.
func DateUpdatedHasPrefix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasPrefix(FieldDateUpdated, v))
}

// DateUpdatedHasSuffix applies the HasSuffix predicate on the "dateUpdated" field.
func DateUpdatedHasSuffix(v string) predicate.Notification {
	return predicate.Notification(sql.FieldHasSuffix(FieldDateUpdated, v))
}

// DateUpdatedEqualFold applies the EqualFold predicate on the "dateUpdated" field.
func DateUpdatedEqualFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldEqualFold(FieldDateUpdated, v))
}

// DateUpdatedContainsFold applies the ContainsFold predicate on the "dateUpdated" field.
func DateUpdatedContainsFold(v string) predicate.Notification {
	return predicate.Notification(sql.FieldContainsFold(FieldDateUpdated, v))
}

// ReadEQ applies the EQ predicate on the "read" field.
func ReadEQ(v bool) predicate.Notification {
	return predicate.Notification(sql.FieldEQ(FieldRead, v))
}

// ReadNEQ applies the NEQ predicate on the "read" field.
func ReadNEQ(v bool) predicate.Notification {
	return predicate.Notification(sql.FieldNEQ(FieldRead, v))
}

// HasRecipient applies the HasEdge predicate on the "recipient" edge.
func HasRecipient() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RecipientTable, RecipientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecipientWith applies the HasEdge predicate on the "recipient" edge with a given conditions (other predicates).
func HasRecipientWith(preds ...predicate.User) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := newRecipientStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		p(s.Not())
	})
}