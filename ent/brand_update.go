// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/leeliwei930/go_commerce/ent/brand"
	"github.com/leeliwei930/go_commerce/ent/predicate"
	"github.com/leeliwei930/go_commerce/ent/product"
)

// BrandUpdate is the builder for updating Brand entities.
type BrandUpdate struct {
	config
	hooks    []Hook
	mutation *BrandMutation
}

// Where appends a list predicates to the BrandUpdate builder.
func (bu *BrandUpdate) Where(ps ...predicate.Brand) *BrandUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetName sets the "name" field.
func (bu *BrandUpdate) SetName(s string) *BrandUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetDescription sets the "description" field.
func (bu *BrandUpdate) SetDescription(s string) *BrandUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetIsActive sets the "is_active" field.
func (bu *BrandUpdate) SetIsActive(b bool) *BrandUpdate {
	bu.mutation.SetIsActive(b)
	return bu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (bu *BrandUpdate) SetNillableIsActive(b *bool) *BrandUpdate {
	if b != nil {
		bu.SetIsActive(*b)
	}
	return bu
}

// SetPublishedAt sets the "published_at" field.
func (bu *BrandUpdate) SetPublishedAt(t time.Time) *BrandUpdate {
	bu.mutation.SetPublishedAt(t)
	return bu
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (bu *BrandUpdate) SetNillablePublishedAt(t *time.Time) *BrandUpdate {
	if t != nil {
		bu.SetPublishedAt(*t)
	}
	return bu
}

// ClearPublishedAt clears the value of the "published_at" field.
func (bu *BrandUpdate) ClearPublishedAt() *BrandUpdate {
	bu.mutation.ClearPublishedAt()
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BrandUpdate) SetCreatedAt(t time.Time) *BrandUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BrandUpdate) SetNillableCreatedAt(t *time.Time) *BrandUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BrandUpdate) SetUpdatedAt(t time.Time) *BrandUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bu *BrandUpdate) SetNillableUpdatedAt(t *time.Time) *BrandUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (bu *BrandUpdate) AddProductIDs(ids ...uuid.UUID) *BrandUpdate {
	bu.mutation.AddProductIDs(ids...)
	return bu
}

// AddProducts adds the "products" edges to the Product entity.
func (bu *BrandUpdate) AddProducts(p ...*Product) *BrandUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bu *BrandUpdate) Mutation() *BrandMutation {
	return bu.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (bu *BrandUpdate) ClearProducts() *BrandUpdate {
	bu.mutation.ClearProducts()
	return bu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (bu *BrandUpdate) RemoveProductIDs(ids ...uuid.UUID) *BrandUpdate {
	bu.mutation.RemoveProductIDs(ids...)
	return bu
}

// RemoveProducts removes "products" edges to Product entities.
func (bu *BrandUpdate) RemoveProducts(p ...*Product) *BrandUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BrandUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BrandUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BrandUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BrandUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BrandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   brand.Table,
			Columns: brand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: brand.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldDescription,
		})
	}
	if value, ok := bu.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: brand.FieldIsActive,
		})
	}
	if value, ok := bu.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldPublishedAt,
		})
	}
	if bu.mutation.PublishedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: brand.FieldPublishedAt,
		})
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldCreatedAt,
		})
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldUpdatedAt,
		})
	}
	if bu.mutation.ProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !bu.mutation.ProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProductsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BrandUpdateOne is the builder for updating a single Brand entity.
type BrandUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BrandMutation
}

// SetName sets the "name" field.
func (buo *BrandUpdateOne) SetName(s string) *BrandUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetDescription sets the "description" field.
func (buo *BrandUpdateOne) SetDescription(s string) *BrandUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetIsActive sets the "is_active" field.
func (buo *BrandUpdateOne) SetIsActive(b bool) *BrandUpdateOne {
	buo.mutation.SetIsActive(b)
	return buo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (buo *BrandUpdateOne) SetNillableIsActive(b *bool) *BrandUpdateOne {
	if b != nil {
		buo.SetIsActive(*b)
	}
	return buo
}

// SetPublishedAt sets the "published_at" field.
func (buo *BrandUpdateOne) SetPublishedAt(t time.Time) *BrandUpdateOne {
	buo.mutation.SetPublishedAt(t)
	return buo
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (buo *BrandUpdateOne) SetNillablePublishedAt(t *time.Time) *BrandUpdateOne {
	if t != nil {
		buo.SetPublishedAt(*t)
	}
	return buo
}

// ClearPublishedAt clears the value of the "published_at" field.
func (buo *BrandUpdateOne) ClearPublishedAt() *BrandUpdateOne {
	buo.mutation.ClearPublishedAt()
	return buo
}

// SetCreatedAt sets the "created_at" field.
func (buo *BrandUpdateOne) SetCreatedAt(t time.Time) *BrandUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BrandUpdateOne) SetNillableCreatedAt(t *time.Time) *BrandUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BrandUpdateOne) SetUpdatedAt(t time.Time) *BrandUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buo *BrandUpdateOne) SetNillableUpdatedAt(t *time.Time) *BrandUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (buo *BrandUpdateOne) AddProductIDs(ids ...uuid.UUID) *BrandUpdateOne {
	buo.mutation.AddProductIDs(ids...)
	return buo
}

// AddProducts adds the "products" edges to the Product entity.
func (buo *BrandUpdateOne) AddProducts(p ...*Product) *BrandUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (buo *BrandUpdateOne) Mutation() *BrandMutation {
	return buo.mutation
}

// ClearProducts clears all "products" edges to the Product entity.
func (buo *BrandUpdateOne) ClearProducts() *BrandUpdateOne {
	buo.mutation.ClearProducts()
	return buo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (buo *BrandUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *BrandUpdateOne {
	buo.mutation.RemoveProductIDs(ids...)
	return buo
}

// RemoveProducts removes "products" edges to Product entities.
func (buo *BrandUpdateOne) RemoveProducts(p ...*Product) *BrandUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemoveProductIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BrandUpdateOne) Select(field string, fields ...string) *BrandUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Brand entity.
func (buo *BrandUpdateOne) Save(ctx context.Context) (*Brand, error) {
	var (
		err  error
		node *Brand
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BrandUpdateOne) SaveX(ctx context.Context) *Brand {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BrandUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BrandUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BrandUpdateOne) sqlSave(ctx context.Context) (_node *Brand, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   brand.Table,
			Columns: brand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: brand.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Brand.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, brand.FieldID)
		for _, f := range fields {
			if !brand.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != brand.FieldID {
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
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldName,
		})
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: brand.FieldDescription,
		})
	}
	if value, ok := buo.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: brand.FieldIsActive,
		})
	}
	if value, ok := buo.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldPublishedAt,
		})
	}
	if buo.mutation.PublishedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: brand.FieldPublishedAt,
		})
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldCreatedAt,
		})
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: brand.FieldUpdatedAt,
		})
	}
	if buo.mutation.ProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !buo.mutation.ProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProductsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Brand{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
