package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/leeliwei930/go_commerce/validators"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.Text("description"),
		field.Bool("is_active").Default(false),
		field.String("published_at").
			SchemaType(map[string]string{
				dialect.MySQL: "timestamp",
			}).
			Validate(validators.IsValidDate()).
			Nillable().
			Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("brand", Brand.Type).Ref("products").Unique(),
	}
}
