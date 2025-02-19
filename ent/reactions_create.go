// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/glu/video-real-time-ranking/ent/reactions"
	"github.com/glu/video-real-time-ranking/ent/videos"
)

// ReactionsCreate is the builder for creating a Reactions entity.
type ReactionsCreate struct {
	config
	mutation *ReactionsMutation
	hooks    []Hook
}

// SetVideoID sets the "video_id" field.
func (rc *ReactionsCreate) SetVideoID(u uint) *ReactionsCreate {
	rc.mutation.SetVideoID(u)
	return rc
}

// SetDescription sets the "description" field.
func (rc *ReactionsCreate) SetDescription(s string) *ReactionsCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *ReactionsCreate) SetNillableDescription(s *string) *ReactionsCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *ReactionsCreate) SetName(s string) *ReactionsCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetNumber sets the "number" field.
func (rc *ReactionsCreate) SetNumber(i int) *ReactionsCreate {
	rc.mutation.SetNumber(i)
	return rc
}

// SetTimePoint sets the "time_point" field.
func (rc *ReactionsCreate) SetTimePoint(f float64) *ReactionsCreate {
	rc.mutation.SetTimePoint(f)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReactionsCreate) SetCreatedAt(t time.Time) *ReactionsCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReactionsCreate) SetNillableCreatedAt(t *time.Time) *ReactionsCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ReactionsCreate) SetUpdatedAt(t time.Time) *ReactionsCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ReactionsCreate) SetNillableUpdatedAt(t *time.Time) *ReactionsCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *ReactionsCreate) SetDeletedAt(t time.Time) *ReactionsCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *ReactionsCreate) SetNillableDeletedAt(t *time.Time) *ReactionsCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReactionsCreate) SetID(u uint) *ReactionsCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetTblVideosID sets the "tbl_videos" edge to the Videos entity by ID.
func (rc *ReactionsCreate) SetTblVideosID(id uint) *ReactionsCreate {
	rc.mutation.SetTblVideosID(id)
	return rc
}

// SetTblVideos sets the "tbl_videos" edge to the Videos entity.
func (rc *ReactionsCreate) SetTblVideos(v *Videos) *ReactionsCreate {
	return rc.SetTblVideosID(v.ID)
}

// Mutation returns the ReactionsMutation object of the builder.
func (rc *ReactionsCreate) Mutation() *ReactionsMutation {
	return rc.mutation
}

// Save creates the Reactions in the database.
func (rc *ReactionsCreate) Save(ctx context.Context) (*Reactions, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReactionsCreate) SaveX(ctx context.Context) *Reactions {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReactionsCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReactionsCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReactionsCreate) defaults() {
	if _, ok := rc.mutation.Description(); !ok {
		v := reactions.DefaultDescription
		rc.mutation.SetDescription(v)
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := reactions.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := reactions.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReactionsCreate) check() error {
	if _, ok := rc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`ent: missing required field "Reactions.video_id"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Reactions.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := reactions.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Reactions.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Reactions.number"`)}
	}
	if v, ok := rc.mutation.Number(); ok {
		if err := reactions.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Reactions.number": %w`, err)}
		}
	}
	if _, ok := rc.mutation.TimePoint(); !ok {
		return &ValidationError{Name: "time_point", err: errors.New(`ent: missing required field "Reactions.time_point"`)}
	}
	if v, ok := rc.mutation.TimePoint(); ok {
		if err := reactions.TimePointValidator(v); err != nil {
			return &ValidationError{Name: "time_point", err: fmt.Errorf(`ent: validator failed for field "Reactions.time_point": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Reactions.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Reactions.updated_at"`)}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := reactions.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Reactions.id": %w`, err)}
		}
	}
	if _, ok := rc.mutation.TblVideosID(); !ok {
		return &ValidationError{Name: "tbl_videos", err: errors.New(`ent: missing required edge "Reactions.tbl_videos"`)}
	}
	return nil
}

func (rc *ReactionsCreate) sqlSave(ctx context.Context) (*Reactions, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReactionsCreate) createSpec() (*Reactions, *sqlgraph.CreateSpec) {
	var (
		_node = &Reactions{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(reactions.Table, sqlgraph.NewFieldSpec(reactions.FieldID, field.TypeUint))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(reactions.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(reactions.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Number(); ok {
		_spec.SetField(reactions.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := rc.mutation.TimePoint(); ok {
		_spec.SetField(reactions.FieldTimePoint, field.TypeFloat64, value)
		_node.TimePoint = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(reactions.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(reactions.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(reactions.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := rc.mutation.TblVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reactions.TblVideosTable,
			Columns: []string{reactions.TblVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(videos.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VideoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReactionsCreateBulk is the builder for creating many Reactions entities in bulk.
type ReactionsCreateBulk struct {
	config
	err      error
	builders []*ReactionsCreate
}

// Save creates the Reactions entities in the database.
func (rcb *ReactionsCreateBulk) Save(ctx context.Context) ([]*Reactions, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Reactions, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReactionsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReactionsCreateBulk) SaveX(ctx context.Context) []*Reactions {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReactionsCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReactionsCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
