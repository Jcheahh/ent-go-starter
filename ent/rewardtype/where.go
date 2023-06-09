// Code generated by ent, DO NOT EDIT.

package rewardtype

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RewardType {
	return predicate.RewardType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RewardType {
	return predicate.RewardType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RewardType {
	return predicate.RewardType(sql.FieldLTE(FieldID, id))
}

// Val applies equality check predicate on the "val" field. It's identical to ValEQ.
func Val(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldEQ(FieldVal, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.RewardType {
	return predicate.RewardType(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.RewardType {
	return predicate.RewardType(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.RewardType {
	return predicate.RewardType(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.RewardType {
	return predicate.RewardType(sql.FieldNotIn(FieldType, vs...))
}

// ValEQ applies the EQ predicate on the "val" field.
func ValEQ(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldEQ(FieldVal, v))
}

// ValNEQ applies the NEQ predicate on the "val" field.
func ValNEQ(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldNEQ(FieldVal, v))
}

// ValIn applies the In predicate on the "val" field.
func ValIn(vs ...int) predicate.RewardType {
	return predicate.RewardType(sql.FieldIn(FieldVal, vs...))
}

// ValNotIn applies the NotIn predicate on the "val" field.
func ValNotIn(vs ...int) predicate.RewardType {
	return predicate.RewardType(sql.FieldNotIn(FieldVal, vs...))
}

// ValGT applies the GT predicate on the "val" field.
func ValGT(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldGT(FieldVal, v))
}

// ValGTE applies the GTE predicate on the "val" field.
func ValGTE(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldGTE(FieldVal, v))
}

// ValLT applies the LT predicate on the "val" field.
func ValLT(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldLT(FieldVal, v))
}

// ValLTE applies the LTE predicate on the "val" field.
func ValLTE(v int) predicate.RewardType {
	return predicate.RewardType(sql.FieldLTE(FieldVal, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RewardType) predicate.RewardType {
	return predicate.RewardType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RewardType) predicate.RewardType {
	return predicate.RewardType(func(s *sql.Selector) {
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
func Not(p predicate.RewardType) predicate.RewardType {
	return predicate.RewardType(func(s *sql.Selector) {
		p(s.Not())
	})
}
