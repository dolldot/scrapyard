// Code generated by ent, DO NOT EDIT.

package bookmark

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldUUID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldURL, v))
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldAccountNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldUpdatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldUUID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldContainsFold(FieldName, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldContainsFold(FieldURL, v))
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldAccountNumber, v))
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldAccountNumber, v))
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldAccountNumber, vs...))
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldAccountNumber, vs...))
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldAccountNumber, v))
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldAccountNumber, v))
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldAccountNumber, v))
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldAccountNumber, v))
}

// AccountNumberContains applies the Contains predicate on the "account_number" field.
func AccountNumberContains(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldContains(FieldAccountNumber, v))
}

// AccountNumberHasPrefix applies the HasPrefix predicate on the "account_number" field.
func AccountNumberHasPrefix(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldHasPrefix(FieldAccountNumber, v))
}

// AccountNumberHasSuffix applies the HasSuffix predicate on the "account_number" field.
func AccountNumberHasSuffix(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldHasSuffix(FieldAccountNumber, v))
}

// AccountNumberEqualFold applies the EqualFold predicate on the "account_number" field.
func AccountNumberEqualFold(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEqualFold(FieldAccountNumber, v))
}

// AccountNumberContainsFold applies the ContainsFold predicate on the "account_number" field.
func AccountNumberContainsFold(v string) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldContainsFold(FieldAccountNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Bookmark {
	return predicate.Bookmark(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bookmark) predicate.Bookmark {
	return predicate.Bookmark(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bookmark) predicate.Bookmark {
	return predicate.Bookmark(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bookmark) predicate.Bookmark {
	return predicate.Bookmark(sql.NotPredicates(p))
}
