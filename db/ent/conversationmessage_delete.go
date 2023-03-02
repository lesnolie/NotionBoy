// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"notionboy/db/ent/conversationmessage"
	"notionboy/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConversationMessageDelete is the builder for deleting a ConversationMessage entity.
type ConversationMessageDelete struct {
	config
	hooks    []Hook
	mutation *ConversationMessageMutation
}

// Where appends a list predicates to the ConversationMessageDelete builder.
func (cmd *ConversationMessageDelete) Where(ps ...predicate.ConversationMessage) *ConversationMessageDelete {
	cmd.mutation.Where(ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *ConversationMessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cmd.hooks) == 0 {
		affected, err = cmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConversationMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmd.mutation = mutation
			affected, err = cmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmd.hooks) - 1; i >= 0; i-- {
			if cmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *ConversationMessageDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *ConversationMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: conversationmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversationmessage.FieldID,
			},
		},
	}
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ConversationMessageDeleteOne is the builder for deleting a single ConversationMessage entity.
type ConversationMessageDeleteOne struct {
	cmd *ConversationMessageDelete
}

// Exec executes the deletion query.
func (cmdo *ConversationMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{conversationmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *ConversationMessageDeleteOne) ExecX(ctx context.Context) {
	cmdo.cmd.ExecX(ctx)
}
