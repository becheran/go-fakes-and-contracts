// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/ingredient"
	"github.com/quii/go-fakes-and-contracts/adapters/driven/persistence/sqlite/ent/pantry"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Pantry is the model entity for the Pantry schema.
type Pantry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PantryQuery when eager-loading is set.
	Edges        PantryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PantryEdges holds the relations/edges for other nodes in the graph.
type PantryEdges struct {
	// Ingredient holds the value of the ingredient edge.
	Ingredient *Ingredient `json:"ingredient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// IngredientOrErr returns the Ingredient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PantryEdges) IngredientOrErr() (*Ingredient, error) {
	if e.loadedTypes[0] {
		if e.Ingredient == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ingredient.Label}
		}
		return e.Ingredient, nil
	}
	return nil, &NotLoadedError{edge: "ingredient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pantry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pantry.FieldID, pantry.FieldQuantity:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pantry fields.
func (pa *Pantry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pantry.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case pantry.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				pa.Quantity = int(value.Int64)
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pantry.
// This includes values selected through modifiers, order, etc.
func (pa *Pantry) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryIngredient queries the "ingredient" edge of the Pantry entity.
func (pa *Pantry) QueryIngredient() *IngredientQuery {
	return NewPantryClient(pa.config).QueryIngredient(pa)
}

// Update returns a builder for updating this Pantry.
// Note that you need to call Pantry.Unwrap() before calling this method if this Pantry
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Pantry) Update() *PantryUpdateOne {
	return NewPantryClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Pantry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Pantry) Unwrap() *Pantry {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pantry is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Pantry) String() string {
	var builder strings.Builder
	builder.WriteString("Pantry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", pa.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// Pantries is a parsable slice of Pantry.
type Pantries []*Pantry
