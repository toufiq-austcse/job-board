// Code generated by ent, DO NOT EDIT.

package job

import (
	"entgo.io/ent/dialect/sql"
	"github.com/toufiq-austcse/go-api-boilerplate/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldTitle, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldSlug, v))
}

// ApplyTo applies equality check predicate on the "apply_to" field. It's identical to ApplyToEQ.
func ApplyTo(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldApplyTo, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldDescription, v))
}

// CompanyID applies equality check predicate on the "company_id" field. It's identical to CompanyIDEQ.
func CompanyID(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCompanyID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldTitle, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldSlug, v))
}

// ApplyToEQ applies the EQ predicate on the "apply_to" field.
func ApplyToEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldApplyTo, v))
}

// ApplyToNEQ applies the NEQ predicate on the "apply_to" field.
func ApplyToNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldApplyTo, v))
}

// ApplyToIn applies the In predicate on the "apply_to" field.
func ApplyToIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldApplyTo, vs...))
}

// ApplyToNotIn applies the NotIn predicate on the "apply_to" field.
func ApplyToNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldApplyTo, vs...))
}

// ApplyToGT applies the GT predicate on the "apply_to" field.
func ApplyToGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldApplyTo, v))
}

// ApplyToGTE applies the GTE predicate on the "apply_to" field.
func ApplyToGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldApplyTo, v))
}

// ApplyToLT applies the LT predicate on the "apply_to" field.
func ApplyToLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldApplyTo, v))
}

// ApplyToLTE applies the LTE predicate on the "apply_to" field.
func ApplyToLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldApplyTo, v))
}

// ApplyToContains applies the Contains predicate on the "apply_to" field.
func ApplyToContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldApplyTo, v))
}

// ApplyToHasPrefix applies the HasPrefix predicate on the "apply_to" field.
func ApplyToHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldApplyTo, v))
}

// ApplyToHasSuffix applies the HasSuffix predicate on the "apply_to" field.
func ApplyToHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldApplyTo, v))
}

// ApplyToEqualFold applies the EqualFold predicate on the "apply_to" field.
func ApplyToEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldApplyTo, v))
}

// ApplyToContainsFold applies the ContainsFold predicate on the "apply_to" field.
func ApplyToContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldApplyTo, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldDescription, v))
}

// CompanyIDEQ applies the EQ predicate on the "company_id" field.
func CompanyIDEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCompanyID, v))
}

// CompanyIDNEQ applies the NEQ predicate on the "company_id" field.
func CompanyIDNEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCompanyID, v))
}

// CompanyIDIn applies the In predicate on the "company_id" field.
func CompanyIDIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldCompanyID, vs...))
}

// CompanyIDNotIn applies the NotIn predicate on the "company_id" field.
func CompanyIDNotIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldCompanyID, vs...))
}

// CompanyIDGT applies the GT predicate on the "company_id" field.
func CompanyIDGT(v int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldCompanyID, v))
}

// CompanyIDGTE applies the GTE predicate on the "company_id" field.
func CompanyIDGTE(v int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldCompanyID, v))
}

// CompanyIDLT applies the LT predicate on the "company_id" field.
func CompanyIDLT(v int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldCompanyID, v))
}

// CompanyIDLTE applies the LTE predicate on the "company_id" field.
func CompanyIDLTE(v int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldCompanyID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		p(s.Not())
	})
}