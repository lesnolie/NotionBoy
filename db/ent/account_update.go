// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notionboy/db/ent/account"
	"notionboy/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetDeleted sets the "deleted" field.
func (au *AccountUpdate) SetDeleted(b bool) *AccountUpdate {
	au.mutation.SetDeleted(b)
	return au
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeleted(b *bool) *AccountUpdate {
	if b != nil {
		au.SetDeleted(*b)
	}
	return au
}

// SetUUID sets the "uuid" field.
func (au *AccountUpdate) SetUUID(u uuid.UUID) *AccountUpdate {
	au.mutation.SetUUID(u)
	return au
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUUID(u *uuid.UUID) *AccountUpdate {
	if u != nil {
		au.SetUUID(*u)
	}
	return au
}

// ClearUUID clears the value of the "uuid" field.
func (au *AccountUpdate) ClearUUID() *AccountUpdate {
	au.mutation.ClearUUID()
	return au
}

// SetUserID sets the "user_id" field.
func (au *AccountUpdate) SetUserID(s string) *AccountUpdate {
	au.mutation.SetUserID(s)
	return au
}

// SetUserType sets the "user_type" field.
func (au *AccountUpdate) SetUserType(at account.UserType) *AccountUpdate {
	au.mutation.SetUserType(at)
	return au
}

// SetNillableUserType sets the "user_type" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUserType(at *account.UserType) *AccountUpdate {
	if at != nil {
		au.SetUserType(*at)
	}
	return au
}

// ClearUserType clears the value of the "user_type" field.
func (au *AccountUpdate) ClearUserType() *AccountUpdate {
	au.mutation.ClearUserType()
	return au
}

// SetDatabaseID sets the "database_id" field.
func (au *AccountUpdate) SetDatabaseID(s string) *AccountUpdate {
	au.mutation.SetDatabaseID(s)
	return au
}

// SetNillableDatabaseID sets the "database_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDatabaseID(s *string) *AccountUpdate {
	if s != nil {
		au.SetDatabaseID(*s)
	}
	return au
}

// ClearDatabaseID clears the value of the "database_id" field.
func (au *AccountUpdate) ClearDatabaseID() *AccountUpdate {
	au.mutation.ClearDatabaseID()
	return au
}

// SetAccessToken sets the "access_token" field.
func (au *AccountUpdate) SetAccessToken(s string) *AccountUpdate {
	au.mutation.SetAccessToken(s)
	return au
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAccessToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetAccessToken(*s)
	}
	return au
}

// ClearAccessToken clears the value of the "access_token" field.
func (au *AccountUpdate) ClearAccessToken() *AccountUpdate {
	au.mutation.ClearAccessToken()
	return au
}

// SetNotionUserID sets the "notion_user_id" field.
func (au *AccountUpdate) SetNotionUserID(s string) *AccountUpdate {
	au.mutation.SetNotionUserID(s)
	return au
}

// SetNillableNotionUserID sets the "notion_user_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableNotionUserID(s *string) *AccountUpdate {
	if s != nil {
		au.SetNotionUserID(*s)
	}
	return au
}

// ClearNotionUserID clears the value of the "notion_user_id" field.
func (au *AccountUpdate) ClearNotionUserID() *AccountUpdate {
	au.mutation.ClearNotionUserID()
	return au
}

// SetNotionUserEmail sets the "notion_user_email" field.
func (au *AccountUpdate) SetNotionUserEmail(s string) *AccountUpdate {
	au.mutation.SetNotionUserEmail(s)
	return au
}

// SetNillableNotionUserEmail sets the "notion_user_email" field if the given value is not nil.
func (au *AccountUpdate) SetNillableNotionUserEmail(s *string) *AccountUpdate {
	if s != nil {
		au.SetNotionUserEmail(*s)
	}
	return au
}

// ClearNotionUserEmail clears the value of the "notion_user_email" field.
func (au *AccountUpdate) ClearNotionUserEmail() *AccountUpdate {
	au.mutation.ClearNotionUserEmail()
	return au
}

// SetIsLatestSchema sets the "is_latest_schema" field.
func (au *AccountUpdate) SetIsLatestSchema(b bool) *AccountUpdate {
	au.mutation.SetIsLatestSchema(b)
	return au
}

// SetNillableIsLatestSchema sets the "is_latest_schema" field if the given value is not nil.
func (au *AccountUpdate) SetNillableIsLatestSchema(b *bool) *AccountUpdate {
	if b != nil {
		au.SetIsLatestSchema(*b)
	}
	return au
}

// SetIsOpenaiAPIUser sets the "is_openai_api_user" field.
func (au *AccountUpdate) SetIsOpenaiAPIUser(b bool) *AccountUpdate {
	au.mutation.SetIsOpenaiAPIUser(b)
	return au
}

// SetNillableIsOpenaiAPIUser sets the "is_openai_api_user" field if the given value is not nil.
func (au *AccountUpdate) SetNillableIsOpenaiAPIUser(b *bool) *AccountUpdate {
	if b != nil {
		au.SetIsOpenaiAPIUser(*b)
	}
	return au
}

// SetOpenaiAPIKey sets the "openai_api_key" field.
func (au *AccountUpdate) SetOpenaiAPIKey(s string) *AccountUpdate {
	au.mutation.SetOpenaiAPIKey(s)
	return au
}

// SetNillableOpenaiAPIKey sets the "openai_api_key" field if the given value is not nil.
func (au *AccountUpdate) SetNillableOpenaiAPIKey(s *string) *AccountUpdate {
	if s != nil {
		au.SetOpenaiAPIKey(*s)
	}
	return au
}

// ClearOpenaiAPIKey clears the value of the "openai_api_key" field.
func (au *AccountUpdate) ClearOpenaiAPIKey() *AccountUpdate {
	au.mutation.ClearOpenaiAPIKey()
	return au
}

// SetAPIKey sets the "api_key" field.
func (au *AccountUpdate) SetAPIKey(u uuid.UUID) *AccountUpdate {
	au.mutation.SetAPIKey(u)
	return au
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAPIKey(u *uuid.UUID) *AccountUpdate {
	if u != nil {
		au.SetAPIKey(*u)
	}
	return au
}

// ClearAPIKey clears the value of the "api_key" field.
func (au *AccountUpdate) ClearAPIKey() *AccountUpdate {
	au.mutation.ClearAPIKey()
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.UserType(); ok {
		if err := account.UserTypeValidator(v); err != nil {
			return &ValidationError{Name: "user_type", err: fmt.Errorf(`ent: validator failed for field "Account.user_type": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Deleted(); ok {
		_spec.SetField(account.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := au.mutation.UUID(); ok {
		_spec.SetField(account.FieldUUID, field.TypeUUID, value)
	}
	if au.mutation.UUIDCleared() {
		_spec.ClearField(account.FieldUUID, field.TypeUUID)
	}
	if value, ok := au.mutation.UserID(); ok {
		_spec.SetField(account.FieldUserID, field.TypeString, value)
	}
	if value, ok := au.mutation.UserType(); ok {
		_spec.SetField(account.FieldUserType, field.TypeEnum, value)
	}
	if au.mutation.UserTypeCleared() {
		_spec.ClearField(account.FieldUserType, field.TypeEnum)
	}
	if value, ok := au.mutation.DatabaseID(); ok {
		_spec.SetField(account.FieldDatabaseID, field.TypeString, value)
	}
	if au.mutation.DatabaseIDCleared() {
		_spec.ClearField(account.FieldDatabaseID, field.TypeString)
	}
	if value, ok := au.mutation.AccessToken(); ok {
		_spec.SetField(account.FieldAccessToken, field.TypeString, value)
	}
	if au.mutation.AccessTokenCleared() {
		_spec.ClearField(account.FieldAccessToken, field.TypeString)
	}
	if value, ok := au.mutation.NotionUserID(); ok {
		_spec.SetField(account.FieldNotionUserID, field.TypeString, value)
	}
	if au.mutation.NotionUserIDCleared() {
		_spec.ClearField(account.FieldNotionUserID, field.TypeString)
	}
	if value, ok := au.mutation.NotionUserEmail(); ok {
		_spec.SetField(account.FieldNotionUserEmail, field.TypeString, value)
	}
	if au.mutation.NotionUserEmailCleared() {
		_spec.ClearField(account.FieldNotionUserEmail, field.TypeString)
	}
	if value, ok := au.mutation.IsLatestSchema(); ok {
		_spec.SetField(account.FieldIsLatestSchema, field.TypeBool, value)
	}
	if value, ok := au.mutation.IsOpenaiAPIUser(); ok {
		_spec.SetField(account.FieldIsOpenaiAPIUser, field.TypeBool, value)
	}
	if value, ok := au.mutation.OpenaiAPIKey(); ok {
		_spec.SetField(account.FieldOpenaiAPIKey, field.TypeString, value)
	}
	if au.mutation.OpenaiAPIKeyCleared() {
		_spec.ClearField(account.FieldOpenaiAPIKey, field.TypeString)
	}
	if value, ok := au.mutation.APIKey(); ok {
		_spec.SetField(account.FieldAPIKey, field.TypeUUID, value)
	}
	if au.mutation.APIKeyCleared() {
		_spec.ClearField(account.FieldAPIKey, field.TypeUUID)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetDeleted sets the "deleted" field.
func (auo *AccountUpdateOne) SetDeleted(b bool) *AccountUpdateOne {
	auo.mutation.SetDeleted(b)
	return auo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeleted(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetDeleted(*b)
	}
	return auo
}

// SetUUID sets the "uuid" field.
func (auo *AccountUpdateOne) SetUUID(u uuid.UUID) *AccountUpdateOne {
	auo.mutation.SetUUID(u)
	return auo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUUID(u *uuid.UUID) *AccountUpdateOne {
	if u != nil {
		auo.SetUUID(*u)
	}
	return auo
}

// ClearUUID clears the value of the "uuid" field.
func (auo *AccountUpdateOne) ClearUUID() *AccountUpdateOne {
	auo.mutation.ClearUUID()
	return auo
}

// SetUserID sets the "user_id" field.
func (auo *AccountUpdateOne) SetUserID(s string) *AccountUpdateOne {
	auo.mutation.SetUserID(s)
	return auo
}

// SetUserType sets the "user_type" field.
func (auo *AccountUpdateOne) SetUserType(at account.UserType) *AccountUpdateOne {
	auo.mutation.SetUserType(at)
	return auo
}

// SetNillableUserType sets the "user_type" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUserType(at *account.UserType) *AccountUpdateOne {
	if at != nil {
		auo.SetUserType(*at)
	}
	return auo
}

// ClearUserType clears the value of the "user_type" field.
func (auo *AccountUpdateOne) ClearUserType() *AccountUpdateOne {
	auo.mutation.ClearUserType()
	return auo
}

// SetDatabaseID sets the "database_id" field.
func (auo *AccountUpdateOne) SetDatabaseID(s string) *AccountUpdateOne {
	auo.mutation.SetDatabaseID(s)
	return auo
}

// SetNillableDatabaseID sets the "database_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDatabaseID(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetDatabaseID(*s)
	}
	return auo
}

// ClearDatabaseID clears the value of the "database_id" field.
func (auo *AccountUpdateOne) ClearDatabaseID() *AccountUpdateOne {
	auo.mutation.ClearDatabaseID()
	return auo
}

// SetAccessToken sets the "access_token" field.
func (auo *AccountUpdateOne) SetAccessToken(s string) *AccountUpdateOne {
	auo.mutation.SetAccessToken(s)
	return auo
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAccessToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetAccessToken(*s)
	}
	return auo
}

// ClearAccessToken clears the value of the "access_token" field.
func (auo *AccountUpdateOne) ClearAccessToken() *AccountUpdateOne {
	auo.mutation.ClearAccessToken()
	return auo
}

// SetNotionUserID sets the "notion_user_id" field.
func (auo *AccountUpdateOne) SetNotionUserID(s string) *AccountUpdateOne {
	auo.mutation.SetNotionUserID(s)
	return auo
}

// SetNillableNotionUserID sets the "notion_user_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableNotionUserID(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetNotionUserID(*s)
	}
	return auo
}

// ClearNotionUserID clears the value of the "notion_user_id" field.
func (auo *AccountUpdateOne) ClearNotionUserID() *AccountUpdateOne {
	auo.mutation.ClearNotionUserID()
	return auo
}

// SetNotionUserEmail sets the "notion_user_email" field.
func (auo *AccountUpdateOne) SetNotionUserEmail(s string) *AccountUpdateOne {
	auo.mutation.SetNotionUserEmail(s)
	return auo
}

// SetNillableNotionUserEmail sets the "notion_user_email" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableNotionUserEmail(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetNotionUserEmail(*s)
	}
	return auo
}

// ClearNotionUserEmail clears the value of the "notion_user_email" field.
func (auo *AccountUpdateOne) ClearNotionUserEmail() *AccountUpdateOne {
	auo.mutation.ClearNotionUserEmail()
	return auo
}

// SetIsLatestSchema sets the "is_latest_schema" field.
func (auo *AccountUpdateOne) SetIsLatestSchema(b bool) *AccountUpdateOne {
	auo.mutation.SetIsLatestSchema(b)
	return auo
}

// SetNillableIsLatestSchema sets the "is_latest_schema" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableIsLatestSchema(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetIsLatestSchema(*b)
	}
	return auo
}

// SetIsOpenaiAPIUser sets the "is_openai_api_user" field.
func (auo *AccountUpdateOne) SetIsOpenaiAPIUser(b bool) *AccountUpdateOne {
	auo.mutation.SetIsOpenaiAPIUser(b)
	return auo
}

// SetNillableIsOpenaiAPIUser sets the "is_openai_api_user" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableIsOpenaiAPIUser(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetIsOpenaiAPIUser(*b)
	}
	return auo
}

// SetOpenaiAPIKey sets the "openai_api_key" field.
func (auo *AccountUpdateOne) SetOpenaiAPIKey(s string) *AccountUpdateOne {
	auo.mutation.SetOpenaiAPIKey(s)
	return auo
}

// SetNillableOpenaiAPIKey sets the "openai_api_key" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableOpenaiAPIKey(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetOpenaiAPIKey(*s)
	}
	return auo
}

// ClearOpenaiAPIKey clears the value of the "openai_api_key" field.
func (auo *AccountUpdateOne) ClearOpenaiAPIKey() *AccountUpdateOne {
	auo.mutation.ClearOpenaiAPIKey()
	return auo
}

// SetAPIKey sets the "api_key" field.
func (auo *AccountUpdateOne) SetAPIKey(u uuid.UUID) *AccountUpdateOne {
	auo.mutation.SetAPIKey(u)
	return auo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAPIKey(u *uuid.UUID) *AccountUpdateOne {
	if u != nil {
		auo.SetAPIKey(*u)
	}
	return auo
}

// ClearAPIKey clears the value of the "api_key" field.
func (auo *AccountUpdateOne) ClearAPIKey() *AccountUpdateOne {
	auo.mutation.ClearAPIKey()
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.UserType(); ok {
		if err := account.UserTypeValidator(v); err != nil {
			return &ValidationError{Name: "user_type", err: fmt.Errorf(`ent: validator failed for field "Account.user_type": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Deleted(); ok {
		_spec.SetField(account.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := auo.mutation.UUID(); ok {
		_spec.SetField(account.FieldUUID, field.TypeUUID, value)
	}
	if auo.mutation.UUIDCleared() {
		_spec.ClearField(account.FieldUUID, field.TypeUUID)
	}
	if value, ok := auo.mutation.UserID(); ok {
		_spec.SetField(account.FieldUserID, field.TypeString, value)
	}
	if value, ok := auo.mutation.UserType(); ok {
		_spec.SetField(account.FieldUserType, field.TypeEnum, value)
	}
	if auo.mutation.UserTypeCleared() {
		_spec.ClearField(account.FieldUserType, field.TypeEnum)
	}
	if value, ok := auo.mutation.DatabaseID(); ok {
		_spec.SetField(account.FieldDatabaseID, field.TypeString, value)
	}
	if auo.mutation.DatabaseIDCleared() {
		_spec.ClearField(account.FieldDatabaseID, field.TypeString)
	}
	if value, ok := auo.mutation.AccessToken(); ok {
		_spec.SetField(account.FieldAccessToken, field.TypeString, value)
	}
	if auo.mutation.AccessTokenCleared() {
		_spec.ClearField(account.FieldAccessToken, field.TypeString)
	}
	if value, ok := auo.mutation.NotionUserID(); ok {
		_spec.SetField(account.FieldNotionUserID, field.TypeString, value)
	}
	if auo.mutation.NotionUserIDCleared() {
		_spec.ClearField(account.FieldNotionUserID, field.TypeString)
	}
	if value, ok := auo.mutation.NotionUserEmail(); ok {
		_spec.SetField(account.FieldNotionUserEmail, field.TypeString, value)
	}
	if auo.mutation.NotionUserEmailCleared() {
		_spec.ClearField(account.FieldNotionUserEmail, field.TypeString)
	}
	if value, ok := auo.mutation.IsLatestSchema(); ok {
		_spec.SetField(account.FieldIsLatestSchema, field.TypeBool, value)
	}
	if value, ok := auo.mutation.IsOpenaiAPIUser(); ok {
		_spec.SetField(account.FieldIsOpenaiAPIUser, field.TypeBool, value)
	}
	if value, ok := auo.mutation.OpenaiAPIKey(); ok {
		_spec.SetField(account.FieldOpenaiAPIKey, field.TypeString, value)
	}
	if auo.mutation.OpenaiAPIKeyCleared() {
		_spec.ClearField(account.FieldOpenaiAPIKey, field.TypeString)
	}
	if value, ok := auo.mutation.APIKey(); ok {
		_spec.SetField(account.FieldAPIKey, field.TypeUUID, value)
	}
	if auo.mutation.APIKeyCleared() {
		_spec.ClearField(account.FieldAPIKey, field.TypeUUID)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
