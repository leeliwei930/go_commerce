package schema

import "entgo.io/ent"

// PaymentChannel holds the schema definition for the PaymentChannel entity.
type PaymentChannel struct {
	ent.Schema
}

// Fields of the PaymentChannel.
func (PaymentChannel) Fields() []ent.Field {
	return nil
}

// Edges of the PaymentChannel.
func (PaymentChannel) Edges() []ent.Edge {
	return nil
}
