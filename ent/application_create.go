// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/ent/group"
)

// ApplicationCreate is the builder for creating a Application entity.
type ApplicationCreate struct {
	config
	mutation *ApplicationMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ac *ApplicationCreate) SetName(s string) *ApplicationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetStatus sets the status field.
func (ac *ApplicationCreate) SetStatus(s string) *ApplicationCreate {
	ac.mutation.SetStatus(s)
	return ac
}

// SetCreatedAt sets the created_at field.
func (ac *ApplicationCreate) SetCreatedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreatedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the updated_at field.
func (ac *ApplicationCreate) SetUpdatedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableUpdatedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetScenario sets the scenario field.
func (ac *ApplicationCreate) SetScenario(s string) *ApplicationCreate {
	ac.mutation.SetScenario(s)
	return ac
}

// SetTags sets the tags field.
func (ac *ApplicationCreate) SetTags(s string) *ApplicationCreate {
	ac.mutation.SetTags(s)
	return ac
}

// SetNillableTags sets the tags field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableTags(s *string) *ApplicationCreate {
	if s != nil {
		ac.SetTags(*s)
	}
	return ac
}

// AddGroupIDs adds the groups edge to Group by ids.
func (ac *ApplicationCreate) AddGroupIDs(ids ...int) *ApplicationCreate {
	ac.mutation.AddGroupIDs(ids...)
	return ac
}

// AddGroups adds the groups edges to Group.
func (ac *ApplicationCreate) AddGroups(g ...*Group) *ApplicationCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ac.AddGroupIDs(ids...)
}

// Mutation returns the ApplicationMutation object of the builder.
func (ac *ApplicationCreate) Mutation() *ApplicationMutation {
	return ac.mutation
}

// Save creates the Application in the database.
func (ac *ApplicationCreate) Save(ctx context.Context) (*Application, error) {
	if err := ac.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Application
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApplicationCreate) SaveX(ctx context.Context) *Application {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *ApplicationCreate) preSave() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := application.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := application.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Scenario(); !ok {
		return &ValidationError{Name: "scenario", err: errors.New("ent: missing required field \"scenario\"")}
	}
	if _, ok := ac.mutation.Tags(); !ok {
		v := application.DefaultTags
		ac.mutation.SetTags(v)
	}
	return nil
}

func (ac *ApplicationCreate) sqlSave(ctx context.Context) (*Application, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *ApplicationCreate) createSpec() (*Application, *sqlgraph.CreateSpec) {
	var (
		a     = &Application{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: application.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: application.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldName,
		})
		a.Name = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldStatus,
		})
		a.Status = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldCreatedAt,
		})
		a.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldUpdatedAt,
		})
		a.UpdatedAt = value
	}
	if value, ok := ac.mutation.Scenario(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldScenario,
		})
		a.Scenario = value
	}
	if value, ok := ac.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldTags,
		})
		a.Tags = value
	}
	if nodes := ac.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.GroupsTable,
			Columns: []string{application.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}

// ApplicationCreateBulk is the builder for creating a bulk of Application entities.
type ApplicationCreateBulk struct {
	config
	builders []*ApplicationCreate
}

// Save creates the Application entities in the database.
func (acb *ApplicationCreateBulk) Save(ctx context.Context) ([]*Application, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Application, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ApplicationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (acb *ApplicationCreateBulk) SaveX(ctx context.Context) []*Application {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
