// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/leeliwei930/go_commerce/ent/brand"
	"github.com/leeliwei930/go_commerce/ent/product"
)

// BrandCreate is the builder for creating a Brand entity.
type BrandCreate struct {
	config
	mutation *BrandMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BrandCreate) SetName(s string) *BrandCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BrandCreate) SetDescription(s string) *BrandCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetIsActive sets the "is_active" field.
func (bc *BrandCreate) SetIsActive(b bool) *BrandCreate {
	bc.mutation.SetIsActive(b)
	return bc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (bc *BrandCreate) SetNillableIsActive(b *bool) *BrandCreate {
	if b != nil {
		bc.SetIsActive(*b)
	}
	return bc
}

// SetPublishedAt sets the "published_at" field.
func (bc *BrandCreate) SetPublishedAt(t time.Time) *BrandCreate {
	bc.mutation.SetPublishedAt(t)
	return bc
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (bc *BrandCreate) SetNillablePublishedAt(t *time.Time) *BrandCreate {
	if t != nil {
		bc.SetPublishedAt(*t)
	}
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BrandCreate) SetCreatedAt(t time.Time) *BrandCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BrandCreate) SetNillableCreatedAt(t *time.Time) *BrandCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BrandCreate) SetUpdatedAt(t time.Time) *BrandCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BrandCreate) SetNillableUpdatedAt(t *time.Time) *BrandCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BrandCreate) SetID(u uuid.UUID) *BrandCreate {
	bc.mutation.SetID(u)
	return bc
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (bc *BrandCreate) AddProductIDs(ids ...uuid.UUID) *BrandCreate {
	bc.mutation.AddProductIDs(ids...)
	return bc
}

// AddProducts adds the "products" edges to the Product entity.
func (bc *BrandCreate) AddProducts(p ...*Product) *BrandCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bc *BrandCreate) Mutation() *BrandMutation {
	return bc.mutation
}

// Save creates the Brand in the database.
func (bc *BrandCreate) Save(ctx context.Context) (*Brand, error) {
	var (
		err  error
		node *Brand
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BrandCreate) SaveX(ctx context.Context) *Brand {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BrandCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BrandCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BrandCreate) defaults() {
	if _, ok := bc.mutation.IsActive(); !ok {
		v := brand.DefaultIsActive
		bc.mutation.SetIsActive(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := brand.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := brand.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := brand.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BrandCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := bc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := bc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "is_active"`)}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (bc *BrandCreate) sqlSave(ctx context.Context) (*Brand, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (bc *BrandCreate) createSpec() (*Brand, *sqlgraph.CreateSpec) {
	var (
		_node = &Brand{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: brand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: brand.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := bc.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: brand.FieldIsActive,
		})
		_node.IsActive = value
	}
	if value, ok := bc.mutation.PublishedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldPublishedAt,
		})
		_node.PublishedAt = &value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductsTable,
			Columns: []string{brand.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BrandCreateBulk is the builder for creating many Brand entities in bulk.
type BrandCreateBulk struct {
	config
	builders []*BrandCreate
}

// Save creates the Brand entities in the database.
func (bcb *BrandCreateBulk) Save(ctx context.Context) ([]*Brand, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Brand, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BrandMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BrandCreateBulk) SaveX(ctx context.Context) []*Brand {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BrandCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BrandCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}