// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/predicate"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/recipeingredient"
)

// RecipeIngredientDelete is the builder for deleting a RecipeIngredient entity.
type RecipeIngredientDelete struct {
	config
	hooks    []Hook
	mutation *RecipeIngredientMutation
}

// Where appends a list predicates to the RecipeIngredientDelete builder.
func (rid *RecipeIngredientDelete) Where(ps ...predicate.RecipeIngredient) *RecipeIngredientDelete {
	rid.mutation.Where(ps...)
	return rid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rid *RecipeIngredientDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rid.sqlExec, rid.mutation, rid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rid *RecipeIngredientDelete) ExecX(ctx context.Context) int {
	n, err := rid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rid *RecipeIngredientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(recipeingredient.Table, sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt))
	if ps := rid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rid.mutation.done = true
	return affected, err
}

// RecipeIngredientDeleteOne is the builder for deleting a single RecipeIngredient entity.
type RecipeIngredientDeleteOne struct {
	rid *RecipeIngredientDelete
}

// Where appends a list predicates to the RecipeIngredientDelete builder.
func (rido *RecipeIngredientDeleteOne) Where(ps ...predicate.RecipeIngredient) *RecipeIngredientDeleteOne {
	rido.rid.mutation.Where(ps...)
	return rido
}

// Exec executes the deletion query.
func (rido *RecipeIngredientDeleteOne) Exec(ctx context.Context) error {
	n, err := rido.rid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{recipeingredient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rido *RecipeIngredientDeleteOne) ExecX(ctx context.Context) {
	if err := rido.Exec(ctx); err != nil {
		panic(err)
	}
}
