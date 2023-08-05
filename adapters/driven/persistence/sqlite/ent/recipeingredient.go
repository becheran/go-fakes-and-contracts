// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/recipe"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/recipeingredient"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RecipeIngredient is the model entity for the RecipeIngredient schema.
type RecipeIngredient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecipeIngredientQuery when eager-loading is set.
	Edges                   RecipeIngredientEdges `json:"edges"`
	recipe_recipeingredient *int
	selectValues            sql.SelectValues
}

// RecipeIngredientEdges holds the relations/edges for other nodes in the graph.
type RecipeIngredientEdges struct {
	// Recipe holds the value of the recipe edge.
	Recipe *Recipe `json:"recipe,omitempty"`
	// Ingredient holds the value of the ingredient edge.
	Ingredient []*Ingredient `json:"ingredient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RecipeOrErr returns the Recipe value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecipeIngredientEdges) RecipeOrErr() (*Recipe, error) {
	if e.loadedTypes[0] {
		if e.Recipe == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: recipe.Label}
		}
		return e.Recipe, nil
	}
	return nil, &NotLoadedError{edge: "recipe"}
}

// IngredientOrErr returns the Ingredient value or an error if the edge
// was not loaded in eager-loading.
func (e RecipeIngredientEdges) IngredientOrErr() ([]*Ingredient, error) {
	if e.loadedTypes[1] {
		return e.Ingredient, nil
	}
	return nil, &NotLoadedError{edge: "ingredient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RecipeIngredient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case recipeingredient.FieldID, recipeingredient.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case recipeingredient.ForeignKeys[0]: // recipe_recipeingredient
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RecipeIngredient fields.
func (ri *RecipeIngredient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recipeingredient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ri.ID = int(value.Int64)
		case recipeingredient.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				ri.Quantity = int(value.Int64)
			}
		case recipeingredient.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field recipe_recipeingredient", value)
			} else if value.Valid {
				ri.recipe_recipeingredient = new(int)
				*ri.recipe_recipeingredient = int(value.Int64)
			}
		default:
			ri.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RecipeIngredient.
// This includes values selected through modifiers, order, etc.
func (ri *RecipeIngredient) Value(name string) (ent.Value, error) {
	return ri.selectValues.Get(name)
}

// QueryRecipe queries the "recipe" edge of the RecipeIngredient entity.
func (ri *RecipeIngredient) QueryRecipe() *RecipeQuery {
	return NewRecipeIngredientClient(ri.config).QueryRecipe(ri)
}

// QueryIngredient queries the "ingredient" edge of the RecipeIngredient entity.
func (ri *RecipeIngredient) QueryIngredient() *IngredientQuery {
	return NewRecipeIngredientClient(ri.config).QueryIngredient(ri)
}

// Update returns a builder for updating this RecipeIngredient.
// Note that you need to call RecipeIngredient.Unwrap() before calling this method if this RecipeIngredient
// was returned from a transaction, and the transaction was committed or rolled back.
func (ri *RecipeIngredient) Update() *RecipeIngredientUpdateOne {
	return NewRecipeIngredientClient(ri.config).UpdateOne(ri)
}

// Unwrap unwraps the RecipeIngredient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ri *RecipeIngredient) Unwrap() *RecipeIngredient {
	_tx, ok := ri.config.driver.(*txDriver)
	if !ok {
		panic("ent: RecipeIngredient is not a transactional entity")
	}
	ri.config.driver = _tx.drv
	return ri
}

// String implements the fmt.Stringer.
func (ri *RecipeIngredient) String() string {
	var builder strings.Builder
	builder.WriteString("RecipeIngredient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ri.ID))
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", ri.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// RecipeIngredients is a parsable slice of RecipeIngredient.
type RecipeIngredients []*RecipeIngredient
