package schema

import "entgo.io/ent"

// PaymentMethod holds the schema definition for the PaymentMethod entity.
type PaymentMethod struct {
	ent.Schema
}

// Fields of the PaymentMethod.
func (PaymentMethod) Fields() []ent.Field {
	return nil
}

// Edges of the PaymentMethod.
func (PaymentMethod) Edges() []ent.Edge {
	return nil
}
