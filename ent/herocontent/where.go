// Code generated by ent, DO NOT EDIT.

package herocontent

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldLTE(FieldID, id))
}

// PrimaryMessage applies equality check predicate on the "primaryMessage" field. It's identical to PrimaryMessageEQ.
func PrimaryMessage(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEQ(FieldPrimaryMessage, v))
}

// SecondaryMessage applies equality check predicate on the "secondaryMessage" field. It's identical to SecondaryMessageEQ.
func SecondaryMessage(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEQ(FieldSecondaryMessage, v))
}

// PrimaryMessageEQ applies the EQ predicate on the "primaryMessage" field.
func PrimaryMessageEQ(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEQ(FieldPrimaryMessage, v))
}

// PrimaryMessageNEQ applies the NEQ predicate on the "primaryMessage" field.
func PrimaryMessageNEQ(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldNEQ(FieldPrimaryMessage, v))
}

// PrimaryMessageIn applies the In predicate on the "primaryMessage" field.
func PrimaryMessageIn(vs ...string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldIn(FieldPrimaryMessage, vs...))
}

// PrimaryMessageNotIn applies the NotIn predicate on the "primaryMessage" field.
func PrimaryMessageNotIn(vs ...string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldNotIn(FieldPrimaryMessage, vs...))
}

// PrimaryMessageGT applies the GT predicate on the "primaryMessage" field.
func PrimaryMessageGT(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldGT(FieldPrimaryMessage, v))
}

// PrimaryMessageGTE applies the GTE predicate on the "primaryMessage" field.
func PrimaryMessageGTE(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldGTE(FieldPrimaryMessage, v))
}

// PrimaryMessageLT applies the LT predicate on the "primaryMessage" field.
func PrimaryMessageLT(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldLT(FieldPrimaryMessage, v))
}

// PrimaryMessageLTE applies the LTE predicate on the "primaryMessage" field.
func PrimaryMessageLTE(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldLTE(FieldPrimaryMessage, v))
}

// PrimaryMessageContains applies the Contains predicate on the "primaryMessage" field.
func PrimaryMessageContains(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldContains(FieldPrimaryMessage, v))
}

// PrimaryMessageHasPrefix applies the HasPrefix predicate on the "primaryMessage" field.
func PrimaryMessageHasPrefix(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldHasPrefix(FieldPrimaryMessage, v))
}

// PrimaryMessageHasSuffix applies the HasSuffix predicate on the "primaryMessage" field.
func PrimaryMessageHasSuffix(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldHasSuffix(FieldPrimaryMessage, v))
}

// PrimaryMessageEqualFold applies the EqualFold predicate on the "primaryMessage" field.
func PrimaryMessageEqualFold(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEqualFold(FieldPrimaryMessage, v))
}

// PrimaryMessageContainsFold applies the ContainsFold predicate on the "primaryMessage" field.
func PrimaryMessageContainsFold(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldContainsFold(FieldPrimaryMessage, v))
}

// SecondaryMessageEQ applies the EQ predicate on the "secondaryMessage" field.
func SecondaryMessageEQ(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEQ(FieldSecondaryMessage, v))
}

// SecondaryMessageNEQ applies the NEQ predicate on the "secondaryMessage" field.
func SecondaryMessageNEQ(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldNEQ(FieldSecondaryMessage, v))
}

// SecondaryMessageIn applies the In predicate on the "secondaryMessage" field.
func SecondaryMessageIn(vs ...string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldIn(FieldSecondaryMessage, vs...))
}

// SecondaryMessageNotIn applies the NotIn predicate on the "secondaryMessage" field.
func SecondaryMessageNotIn(vs ...string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldNotIn(FieldSecondaryMessage, vs...))
}

// SecondaryMessageGT applies the GT predicate on the "secondaryMessage" field.
func SecondaryMessageGT(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldGT(FieldSecondaryMessage, v))
}

// SecondaryMessageGTE applies the GTE predicate on the "secondaryMessage" field.
func SecondaryMessageGTE(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldGTE(FieldSecondaryMessage, v))
}

// SecondaryMessageLT applies the LT predicate on the "secondaryMessage" field.
func SecondaryMessageLT(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldLT(FieldSecondaryMessage, v))
}

// SecondaryMessageLTE applies the LTE predicate on the "secondaryMessage" field.
func SecondaryMessageLTE(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldLTE(FieldSecondaryMessage, v))
}

// SecondaryMessageContains applies the Contains predicate on the "secondaryMessage" field.
func SecondaryMessageContains(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldContains(FieldSecondaryMessage, v))
}

// SecondaryMessageHasPrefix applies the HasPrefix predicate on the "secondaryMessage" field.
func SecondaryMessageHasPrefix(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldHasPrefix(FieldSecondaryMessage, v))
}

// SecondaryMessageHasSuffix applies the HasSuffix predicate on the "secondaryMessage" field.
func SecondaryMessageHasSuffix(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldHasSuffix(FieldSecondaryMessage, v))
}

// SecondaryMessageEqualFold applies the EqualFold predicate on the "secondaryMessage" field.
func SecondaryMessageEqualFold(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldEqualFold(FieldSecondaryMessage, v))
}

// SecondaryMessageContainsFold applies the ContainsFold predicate on the "secondaryMessage" field.
func SecondaryMessageContainsFold(v string) predicate.HeroContent {
	return predicate.HeroContent(sql.FieldContainsFold(FieldSecondaryMessage, v))
}

// HasImage applies the HasEdge predicate on the "image" edge.
func HasImage() predicate.HeroContent {
	return predicate.HeroContent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImageTable, ImageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImageWith applies the HasEdge predicate on the "image" edge with a given conditions (other predicates).
func HasImageWith(preds ...predicate.Image) predicate.HeroContent {
	return predicate.HeroContent(func(s *sql.Selector) {
		step := newImageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HeroContent) predicate.HeroContent {
	return predicate.HeroContent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HeroContent) predicate.HeroContent {
	return predicate.HeroContent(func(s *sql.Selector) {
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
func Not(p predicate.HeroContent) predicate.HeroContent {
	return predicate.HeroContent(func(s *sql.Selector) {
		p(s.Not())
	})
}