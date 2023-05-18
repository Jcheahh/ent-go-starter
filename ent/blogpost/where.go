// Code generated by ent, DO NOT EDIT.

package blogpost

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldContent, v))
}

// DateCreated applies equality check predicate on the "dateCreated" field. It's identical to DateCreatedEQ.
func DateCreated(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldDateCreated, v))
}

// DateUpdated applies equality check predicate on the "dateUpdated" field. It's identical to DateUpdatedEQ.
func DateUpdated(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldDateUpdated, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContainsFold(FieldContent, v))
}

// DateCreatedEQ applies the EQ predicate on the "dateCreated" field.
func DateCreatedEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "dateCreated" field.
func DateCreatedNEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "dateCreated" field.
func DateCreatedIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "dateCreated" field.
func DateCreatedNotIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "dateCreated" field.
func DateCreatedGT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "dateCreated" field.
func DateCreatedGTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "dateCreated" field.
func DateCreatedLT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "dateCreated" field.
func DateCreatedLTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLTE(FieldDateCreated, v))
}

// DateCreatedContains applies the Contains predicate on the "dateCreated" field.
func DateCreatedContains(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContains(FieldDateCreated, v))
}

// DateCreatedHasPrefix applies the HasPrefix predicate on the "dateCreated" field.
func DateCreatedHasPrefix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasPrefix(FieldDateCreated, v))
}

// DateCreatedHasSuffix applies the HasSuffix predicate on the "dateCreated" field.
func DateCreatedHasSuffix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasSuffix(FieldDateCreated, v))
}

// DateCreatedEqualFold applies the EqualFold predicate on the "dateCreated" field.
func DateCreatedEqualFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEqualFold(FieldDateCreated, v))
}

// DateCreatedContainsFold applies the ContainsFold predicate on the "dateCreated" field.
func DateCreatedContainsFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContainsFold(FieldDateCreated, v))
}

// DateUpdatedEQ applies the EQ predicate on the "dateUpdated" field.
func DateUpdatedEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEQ(FieldDateUpdated, v))
}

// DateUpdatedNEQ applies the NEQ predicate on the "dateUpdated" field.
func DateUpdatedNEQ(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNEQ(FieldDateUpdated, v))
}

// DateUpdatedIn applies the In predicate on the "dateUpdated" field.
func DateUpdatedIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldIn(FieldDateUpdated, vs...))
}

// DateUpdatedNotIn applies the NotIn predicate on the "dateUpdated" field.
func DateUpdatedNotIn(vs ...string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldNotIn(FieldDateUpdated, vs...))
}

// DateUpdatedGT applies the GT predicate on the "dateUpdated" field.
func DateUpdatedGT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGT(FieldDateUpdated, v))
}

// DateUpdatedGTE applies the GTE predicate on the "dateUpdated" field.
func DateUpdatedGTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldGTE(FieldDateUpdated, v))
}

// DateUpdatedLT applies the LT predicate on the "dateUpdated" field.
func DateUpdatedLT(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLT(FieldDateUpdated, v))
}

// DateUpdatedLTE applies the LTE predicate on the "dateUpdated" field.
func DateUpdatedLTE(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldLTE(FieldDateUpdated, v))
}

// DateUpdatedContains applies the Contains predicate on the "dateUpdated" field.
func DateUpdatedContains(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContains(FieldDateUpdated, v))
}

// DateUpdatedHasPrefix applies the HasPrefix predicate on the "dateUpdated" field.
func DateUpdatedHasPrefix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasPrefix(FieldDateUpdated, v))
}

// DateUpdatedHasSuffix applies the HasSuffix predicate on the "dateUpdated" field.
func DateUpdatedHasSuffix(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldHasSuffix(FieldDateUpdated, v))
}

// DateUpdatedEqualFold applies the EqualFold predicate on the "dateUpdated" field.
func DateUpdatedEqualFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldEqualFold(FieldDateUpdated, v))
}

// DateUpdatedContainsFold applies the ContainsFold predicate on the "dateUpdated" field.
func DateUpdatedContainsFold(v string) predicate.BlogPost {
	return predicate.BlogPost(sql.FieldContainsFold(FieldDateUpdated, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.BlogPost {
	return predicate.BlogPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.UserSeller) predicate.BlogPost {
	return predicate.BlogPost(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BlogPost) predicate.BlogPost {
	return predicate.BlogPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BlogPost) predicate.BlogPost {
	return predicate.BlogPost(func(s *sql.Selector) {
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
func Not(p predicate.BlogPost) predicate.BlogPost {
	return predicate.BlogPost(func(s *sql.Selector) {
		p(s.Not())
	})
}
