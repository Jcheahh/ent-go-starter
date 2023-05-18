// Code generated by ent, DO NOT EDIT.

package userseller

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldLTE(FieldID, id))
}

// BrandName applies equality check predicate on the "brandName" field. It's identical to BrandNameEQ.
func BrandName(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldEQ(FieldBrandName, v))
}

// BrandNameEQ applies the EQ predicate on the "brandName" field.
func BrandNameEQ(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldEQ(FieldBrandName, v))
}

// BrandNameNEQ applies the NEQ predicate on the "brandName" field.
func BrandNameNEQ(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldNEQ(FieldBrandName, v))
}

// BrandNameIn applies the In predicate on the "brandName" field.
func BrandNameIn(vs ...string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldIn(FieldBrandName, vs...))
}

// BrandNameNotIn applies the NotIn predicate on the "brandName" field.
func BrandNameNotIn(vs ...string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldNotIn(FieldBrandName, vs...))
}

// BrandNameGT applies the GT predicate on the "brandName" field.
func BrandNameGT(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldGT(FieldBrandName, v))
}

// BrandNameGTE applies the GTE predicate on the "brandName" field.
func BrandNameGTE(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldGTE(FieldBrandName, v))
}

// BrandNameLT applies the LT predicate on the "brandName" field.
func BrandNameLT(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldLT(FieldBrandName, v))
}

// BrandNameLTE applies the LTE predicate on the "brandName" field.
func BrandNameLTE(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldLTE(FieldBrandName, v))
}

// BrandNameContains applies the Contains predicate on the "brandName" field.
func BrandNameContains(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldContains(FieldBrandName, v))
}

// BrandNameHasPrefix applies the HasPrefix predicate on the "brandName" field.
func BrandNameHasPrefix(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldHasPrefix(FieldBrandName, v))
}

// BrandNameHasSuffix applies the HasSuffix predicate on the "brandName" field.
func BrandNameHasSuffix(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldHasSuffix(FieldBrandName, v))
}

// BrandNameEqualFold applies the EqualFold predicate on the "brandName" field.
func BrandNameEqualFold(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldEqualFold(FieldBrandName, v))
}

// BrandNameContainsFold applies the ContainsFold predicate on the "brandName" field.
func BrandNameContainsFold(v string) predicate.UserSeller {
	return predicate.UserSeller(sql.FieldContainsFold(FieldBrandName, v))
}

// HasUserProfile applies the HasEdge predicate on the "userProfile" edge.
func HasUserProfile() predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserProfileTable, UserProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserProfileWith applies the HasEdge predicate on the "userProfile" edge with a given conditions (other predicates).
func HasUserProfileWith(preds ...predicate.User) predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
		step := newUserProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShops applies the HasEdge predicate on the "shops" edge.
func HasShops() predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShopsTable, ShopsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShopsWith applies the HasEdge predicate on the "shops" edge with a given conditions (other predicates).
func HasShopsWith(preds ...predicate.Shop) predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
		step := newShopsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserSeller) predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserSeller) predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
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
func Not(p predicate.UserSeller) predicate.UserSeller {
	return predicate.UserSeller(func(s *sql.Selector) {
		p(s.Not())
	})
}
