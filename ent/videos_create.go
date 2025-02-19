// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/glu/video-real-time-ranking/ent/comments"
	"github.com/glu/video-real-time-ranking/ent/objects"
	"github.com/glu/video-real-time-ranking/ent/reactions"
	"github.com/glu/video-real-time-ranking/ent/videos"
	"github.com/glu/video-real-time-ranking/ent/viewers"
)

// VideosCreate is the builder for creating a Videos entity.
type VideosCreate struct {
	config
	mutation *VideosMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vc *VideosCreate) SetName(s string) *VideosCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetDescription sets the "description" field.
func (vc *VideosCreate) SetDescription(s string) *VideosCreate {
	vc.mutation.SetDescription(s)
	return vc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (vc *VideosCreate) SetNillableDescription(s *string) *VideosCreate {
	if s != nil {
		vc.SetDescription(*s)
	}
	return vc
}

// SetVideoURL sets the "video_url" field.
func (vc *VideosCreate) SetVideoURL(s string) *VideosCreate {
	vc.mutation.SetVideoURL(s)
	return vc
}

// SetConfig sets the "config" field.
func (vc *VideosCreate) SetConfig(s string) *VideosCreate {
	vc.mutation.SetConfig(s)
	return vc
}

// SetPathResource sets the "path_resource" field.
func (vc *VideosCreate) SetPathResource(s string) *VideosCreate {
	vc.mutation.SetPathResource(s)
	return vc
}

// SetLevelSystem sets the "level_system" field.
func (vc *VideosCreate) SetLevelSystem(s string) *VideosCreate {
	vc.mutation.SetLevelSystem(s)
	return vc
}

// SetStatus sets the "status" field.
func (vc *VideosCreate) SetStatus(s string) *VideosCreate {
	vc.mutation.SetStatus(s)
	return vc
}

// SetNote sets the "note" field.
func (vc *VideosCreate) SetNote(s string) *VideosCreate {
	vc.mutation.SetNote(s)
	return vc
}

// SetAssign sets the "assign" field.
func (vc *VideosCreate) SetAssign(s string) *VideosCreate {
	vc.mutation.SetAssign(s)
	return vc
}

// SetVersion sets the "version" field.
func (vc *VideosCreate) SetVersion(u uint) *VideosCreate {
	vc.mutation.SetVersion(u)
	return vc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (vc *VideosCreate) SetNillableVersion(u *uint) *VideosCreate {
	if u != nil {
		vc.SetVersion(*u)
	}
	return vc
}

// SetAuthor sets the "Author" field.
func (vc *VideosCreate) SetAuthor(s string) *VideosCreate {
	vc.mutation.SetAuthor(s)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VideosCreate) SetCreatedAt(t time.Time) *VideosCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VideosCreate) SetNillableCreatedAt(t *time.Time) *VideosCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VideosCreate) SetUpdatedAt(t time.Time) *VideosCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VideosCreate) SetNillableUpdatedAt(t *time.Time) *VideosCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetDeletedAt sets the "deleted_at" field.
func (vc *VideosCreate) SetDeletedAt(t time.Time) *VideosCreate {
	vc.mutation.SetDeletedAt(t)
	return vc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vc *VideosCreate) SetNillableDeletedAt(t *time.Time) *VideosCreate {
	if t != nil {
		vc.SetDeletedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VideosCreate) SetID(u uint) *VideosCreate {
	vc.mutation.SetID(u)
	return vc
}

// AddTblCommentIDs adds the "tbl_comments" edge to the Comments entity by IDs.
func (vc *VideosCreate) AddTblCommentIDs(ids ...uint) *VideosCreate {
	vc.mutation.AddTblCommentIDs(ids...)
	return vc
}

// AddTblComments adds the "tbl_comments" edges to the Comments entity.
func (vc *VideosCreate) AddTblComments(c ...*Comments) *VideosCreate {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vc.AddTblCommentIDs(ids...)
}

// AddTblReactionIDs adds the "tbl_reactions" edge to the Reactions entity by IDs.
func (vc *VideosCreate) AddTblReactionIDs(ids ...uint) *VideosCreate {
	vc.mutation.AddTblReactionIDs(ids...)
	return vc
}

// AddTblReactions adds the "tbl_reactions" edges to the Reactions entity.
func (vc *VideosCreate) AddTblReactions(r ...*Reactions) *VideosCreate {
	ids := make([]uint, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return vc.AddTblReactionIDs(ids...)
}

// AddTblViewerIDs adds the "tbl_viewers" edge to the Viewers entity by IDs.
func (vc *VideosCreate) AddTblViewerIDs(ids ...uint) *VideosCreate {
	vc.mutation.AddTblViewerIDs(ids...)
	return vc
}

// AddTblViewers adds the "tbl_viewers" edges to the Viewers entity.
func (vc *VideosCreate) AddTblViewers(v ...*Viewers) *VideosCreate {
	ids := make([]uint, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vc.AddTblViewerIDs(ids...)
}

// AddTblObjectIDs adds the "tbl_objects" edge to the Objects entity by IDs.
func (vc *VideosCreate) AddTblObjectIDs(ids ...uint) *VideosCreate {
	vc.mutation.AddTblObjectIDs(ids...)
	return vc
}

// AddTblObjects adds the "tbl_objects" edges to the Objects entity.
func (vc *VideosCreate) AddTblObjects(o ...*Objects) *VideosCreate {
	ids := make([]uint, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return vc.AddTblObjectIDs(ids...)
}

// Mutation returns the VideosMutation object of the builder.
func (vc *VideosCreate) Mutation() *VideosMutation {
	return vc.mutation
}

// Save creates the Videos in the database.
func (vc *VideosCreate) Save(ctx context.Context) (*Videos, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideosCreate) SaveX(ctx context.Context) *Videos {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideosCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideosCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VideosCreate) defaults() {
	if _, ok := vc.mutation.Description(); !ok {
		v := videos.DefaultDescription
		vc.mutation.SetDescription(v)
	}
	if _, ok := vc.mutation.Version(); !ok {
		v := videos.DefaultVersion
		vc.mutation.SetVersion(v)
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := videos.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := videos.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideosCreate) check() error {
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Videos.name"`)}
	}
	if _, ok := vc.mutation.VideoURL(); !ok {
		return &ValidationError{Name: "video_url", err: errors.New(`ent: missing required field "Videos.video_url"`)}
	}
	if _, ok := vc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "Videos.config"`)}
	}
	if _, ok := vc.mutation.PathResource(); !ok {
		return &ValidationError{Name: "path_resource", err: errors.New(`ent: missing required field "Videos.path_resource"`)}
	}
	if _, ok := vc.mutation.LevelSystem(); !ok {
		return &ValidationError{Name: "level_system", err: errors.New(`ent: missing required field "Videos.level_system"`)}
	}
	if _, ok := vc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Videos.status"`)}
	}
	if _, ok := vc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "Videos.note"`)}
	}
	if _, ok := vc.mutation.Assign(); !ok {
		return &ValidationError{Name: "assign", err: errors.New(`ent: missing required field "Videos.assign"`)}
	}
	if _, ok := vc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Videos.version"`)}
	}
	if _, ok := vc.mutation.Author(); !ok {
		return &ValidationError{Name: "Author", err: errors.New(`ent: missing required field "Videos.Author"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Videos.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Videos.updated_at"`)}
	}
	return nil
}

func (vc *VideosCreate) sqlSave(ctx context.Context) (*Videos, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VideosCreate) createSpec() (*Videos, *sqlgraph.CreateSpec) {
	var (
		_node = &Videos{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(videos.Table, sqlgraph.NewFieldSpec(videos.FieldID, field.TypeUint))
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(videos.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.Description(); ok {
		_spec.SetField(videos.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := vc.mutation.VideoURL(); ok {
		_spec.SetField(videos.FieldVideoURL, field.TypeString, value)
		_node.VideoURL = value
	}
	if value, ok := vc.mutation.Config(); ok {
		_spec.SetField(videos.FieldConfig, field.TypeString, value)
		_node.Config = value
	}
	if value, ok := vc.mutation.PathResource(); ok {
		_spec.SetField(videos.FieldPathResource, field.TypeString, value)
		_node.PathResource = value
	}
	if value, ok := vc.mutation.LevelSystem(); ok {
		_spec.SetField(videos.FieldLevelSystem, field.TypeString, value)
		_node.LevelSystem = value
	}
	if value, ok := vc.mutation.Status(); ok {
		_spec.SetField(videos.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := vc.mutation.Note(); ok {
		_spec.SetField(videos.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := vc.mutation.Assign(); ok {
		_spec.SetField(videos.FieldAssign, field.TypeString, value)
		_node.Assign = value
	}
	if value, ok := vc.mutation.Version(); ok {
		_spec.SetField(videos.FieldVersion, field.TypeUint, value)
		_node.Version = value
	}
	if value, ok := vc.mutation.Author(); ok {
		_spec.SetField(videos.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(videos.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(videos.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.DeletedAt(); ok {
		_spec.SetField(videos.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := vc.mutation.TblCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   videos.TblCommentsTable,
			Columns: []string{videos.TblCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comments.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.TblReactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   videos.TblReactionsTable,
			Columns: []string{videos.TblReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reactions.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.TblViewersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   videos.TblViewersTable,
			Columns: []string{videos.TblViewersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(viewers.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.TblObjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   videos.TblObjectsTable,
			Columns: []string{videos.TblObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(objects.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideosCreateBulk is the builder for creating many Videos entities in bulk.
type VideosCreateBulk struct {
	config
	err      error
	builders []*VideosCreate
}

// Save creates the Videos entities in the database.
func (vcb *VideosCreateBulk) Save(ctx context.Context) ([]*Videos, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Videos, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideosMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideosCreateBulk) SaveX(ctx context.Context) []*Videos {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideosCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideosCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
