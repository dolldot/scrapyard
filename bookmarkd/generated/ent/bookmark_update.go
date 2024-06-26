// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent/bookmark"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent/predicate"
	"github.com/google/uuid"
)

// BookmarkUpdate is the builder for updating Bookmark entities.
type BookmarkUpdate struct {
	config
	hooks    []Hook
	mutation *BookmarkMutation
}

// Where appends a list predicates to the BookmarkUpdate builder.
func (bu *BookmarkUpdate) Where(ps ...predicate.Bookmark) *BookmarkUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUUID sets the "uuid" field.
func (bu *BookmarkUpdate) SetUUID(u uuid.UUID) *BookmarkUpdate {
	bu.mutation.SetUUID(u)
	return bu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableUUID(u *uuid.UUID) *BookmarkUpdate {
	if u != nil {
		bu.SetUUID(*u)
	}
	return bu
}

// SetName sets the "name" field.
func (bu *BookmarkUpdate) SetName(s string) *BookmarkUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableName(s *string) *BookmarkUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// SetURL sets the "url" field.
func (bu *BookmarkUpdate) SetURL(s string) *BookmarkUpdate {
	bu.mutation.SetURL(s)
	return bu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableURL(s *string) *BookmarkUpdate {
	if s != nil {
		bu.SetURL(*s)
	}
	return bu
}

// SetAccountNumber sets the "account_number" field.
func (bu *BookmarkUpdate) SetAccountNumber(s string) *BookmarkUpdate {
	bu.mutation.SetAccountNumber(s)
	return bu
}

// SetNillableAccountNumber sets the "account_number" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableAccountNumber(s *string) *BookmarkUpdate {
	if s != nil {
		bu.SetAccountNumber(*s)
	}
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BookmarkUpdate) SetCreatedAt(t time.Time) *BookmarkUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableCreatedAt(t *time.Time) *BookmarkUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookmarkUpdate) SetUpdatedAt(t time.Time) *BookmarkUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bu *BookmarkUpdate) SetNillableUpdatedAt(t *time.Time) *BookmarkUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// Mutation returns the BookmarkMutation object of the builder.
func (bu *BookmarkUpdate) Mutation() *BookmarkMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookmarkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookmarkUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookmarkUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookmarkUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookmarkUpdate) check() error {
	if v, ok := bu.mutation.AccountNumber(); ok {
		if err := bookmark.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "Bookmark.account_number": %w`, err)}
		}
	}
	return nil
}

func (bu *BookmarkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(bookmark.Table, bookmark.Columns, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UUID(); ok {
		_spec.SetField(bookmark.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(bookmark.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.URL(); ok {
		_spec.SetField(bookmark.FieldURL, field.TypeString, value)
	}
	if value, ok := bu.mutation.AccountNumber(); ok {
		_spec.SetField(bookmark.FieldAccountNumber, field.TypeString, value)
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(bookmark.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(bookmark.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookmark.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookmarkUpdateOne is the builder for updating a single Bookmark entity.
type BookmarkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookmarkMutation
}

// SetUUID sets the "uuid" field.
func (buo *BookmarkUpdateOne) SetUUID(u uuid.UUID) *BookmarkUpdateOne {
	buo.mutation.SetUUID(u)
	return buo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableUUID(u *uuid.UUID) *BookmarkUpdateOne {
	if u != nil {
		buo.SetUUID(*u)
	}
	return buo
}

// SetName sets the "name" field.
func (buo *BookmarkUpdateOne) SetName(s string) *BookmarkUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableName(s *string) *BookmarkUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// SetURL sets the "url" field.
func (buo *BookmarkUpdateOne) SetURL(s string) *BookmarkUpdateOne {
	buo.mutation.SetURL(s)
	return buo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableURL(s *string) *BookmarkUpdateOne {
	if s != nil {
		buo.SetURL(*s)
	}
	return buo
}

// SetAccountNumber sets the "account_number" field.
func (buo *BookmarkUpdateOne) SetAccountNumber(s string) *BookmarkUpdateOne {
	buo.mutation.SetAccountNumber(s)
	return buo
}

// SetNillableAccountNumber sets the "account_number" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableAccountNumber(s *string) *BookmarkUpdateOne {
	if s != nil {
		buo.SetAccountNumber(*s)
	}
	return buo
}

// SetCreatedAt sets the "created_at" field.
func (buo *BookmarkUpdateOne) SetCreatedAt(t time.Time) *BookmarkUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableCreatedAt(t *time.Time) *BookmarkUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookmarkUpdateOne) SetUpdatedAt(t time.Time) *BookmarkUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buo *BookmarkUpdateOne) SetNillableUpdatedAt(t *time.Time) *BookmarkUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// Mutation returns the BookmarkMutation object of the builder.
func (buo *BookmarkUpdateOne) Mutation() *BookmarkMutation {
	return buo.mutation
}

// Where appends a list predicates to the BookmarkUpdate builder.
func (buo *BookmarkUpdateOne) Where(ps ...predicate.Bookmark) *BookmarkUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookmarkUpdateOne) Select(field string, fields ...string) *BookmarkUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bookmark entity.
func (buo *BookmarkUpdateOne) Save(ctx context.Context) (*Bookmark, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookmarkUpdateOne) SaveX(ctx context.Context) *Bookmark {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookmarkUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookmarkUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookmarkUpdateOne) check() error {
	if v, ok := buo.mutation.AccountNumber(); ok {
		if err := bookmark.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "Bookmark.account_number": %w`, err)}
		}
	}
	return nil
}

func (buo *BookmarkUpdateOne) sqlSave(ctx context.Context) (_node *Bookmark, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(bookmark.Table, bookmark.Columns, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bookmark.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookmark.FieldID)
		for _, f := range fields {
			if !bookmark.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookmark.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UUID(); ok {
		_spec.SetField(bookmark.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(bookmark.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.URL(); ok {
		_spec.SetField(bookmark.FieldURL, field.TypeString, value)
	}
	if value, ok := buo.mutation.AccountNumber(); ok {
		_spec.SetField(bookmark.FieldAccountNumber, field.TypeString, value)
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(bookmark.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(bookmark.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Bookmark{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookmark.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
