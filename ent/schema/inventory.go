package schema

import "entgo.io/ent"

// Inventory holds the schema definition for the Inventory entity.
type Inventory struct {
	ent.Schema
}

// Fields of the Inventory.
func (Inventory) Fields() []ent.Field {
	return nil
}

// Edges of the Inventory.
func (Inventory) Edges() []ent.Edge {
	return nil
}
