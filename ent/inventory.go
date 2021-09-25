// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/go_commerce/ent/inventory"
)

// Inventory is the model entity for the Inventory schema.
type Inventory struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Inventory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case inventory.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Inventory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Inventory fields.
func (i *Inventory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case inventory.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Inventory.
// Note that you need to call Inventory.Unwrap() before calling this method if this Inventory
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Inventory) Update() *InventoryUpdateOne {
	return (&InventoryClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Inventory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Inventory) Unwrap() *Inventory {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Inventory is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Inventory) String() string {
	var builder strings.Builder
	builder.WriteString("Inventory(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Inventories is a parsable slice of Inventory.
type Inventories []*Inventory

func (i Inventories) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}