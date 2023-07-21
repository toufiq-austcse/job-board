// Code generated by ent, DO NOT EDIT.

package taxonomy

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldUpdatedAt, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldParentID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldTitle, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldSlug, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldUpdatedAt, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldParentID, v))
}

// ParentIDContains applies the Contains predicate on the "parent_id" field.
func ParentIDContains(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContains(FieldParentID, v))
}

// ParentIDHasPrefix applies the HasPrefix predicate on the "parent_id" field.
func ParentIDHasPrefix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasPrefix(FieldParentID, v))
}

// ParentIDHasSuffix applies the HasSuffix predicate on the "parent_id" field.
func ParentIDHasSuffix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasSuffix(FieldParentID, v))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotNull(FieldParentID))
}

// ParentIDEqualFold applies the EqualFold predicate on the "parent_id" field.
func ParentIDEqualFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEqualFold(FieldParentID, v))
}

// ParentIDContainsFold applies the ContainsFold predicate on the "parent_id" field.
func ParentIDContainsFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContainsFold(FieldParentID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContainsFold(FieldTitle, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugIsNil applies the IsNil predicate on the "slug" field.
func SlugIsNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIsNull(FieldSlug))
}

// SlugNotNil applies the NotNil predicate on the "slug" field.
func SlugNotNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotNull(FieldSlug))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContainsFold(FieldSlug, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Taxonomy {
	return predicate.Taxonomy(sql.FieldContainsFold(FieldType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Taxonomy) predicate.Taxonomy {
	return predicate.Taxonomy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Taxonomy) predicate.Taxonomy {
	return predicate.Taxonomy(func(s *sql.Selector) {
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
func Not(p predicate.Taxonomy) predicate.Taxonomy {
	return predicate.Taxonomy(func(s *sql.Selector) {
		p(s.Not())
	})
}
