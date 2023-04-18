// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"notionboy/db/ent/prompt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Prompt is the model entity for the Prompt schema.
type Prompt struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// UUID
	UUID uuid.UUID `json:"uuid,omitempty"`
	// user id
	UserID uuid.UUID `json:"user_id,omitempty"`
	// role name
	Act string `json:"act,omitempty"`
	// prompt text
	Prompt string `json:"prompt,omitempty"`
	// is user custom prompt
	IsCustom bool `json:"is_custom,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prompt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case prompt.FieldDeleted, prompt.FieldIsCustom:
			values[i] = new(sql.NullBool)
		case prompt.FieldID:
			values[i] = new(sql.NullInt64)
		case prompt.FieldAct, prompt.FieldPrompt:
			values[i] = new(sql.NullString)
		case prompt.FieldCreatedAt, prompt.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case prompt.FieldUUID, prompt.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Prompt", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prompt fields.
func (pr *Prompt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prompt.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case prompt.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case prompt.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case prompt.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				pr.Deleted = value.Bool
			}
		case prompt.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				pr.UUID = *value
			}
		case prompt.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				pr.UserID = *value
			}
		case prompt.FieldAct:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field act", values[i])
			} else if value.Valid {
				pr.Act = value.String
			}
		case prompt.FieldPrompt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prompt", values[i])
			} else if value.Valid {
				pr.Prompt = value.String
			}
		case prompt.FieldIsCustom:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_custom", values[i])
			} else if value.Valid {
				pr.IsCustom = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Prompt.
// Note that you need to call Prompt.Unwrap() before calling this method if this Prompt
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prompt) Update() *PromptUpdateOne {
	return NewPromptClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Prompt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Prompt) Unwrap() *Prompt {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prompt is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prompt) String() string {
	var builder strings.Builder
	builder.WriteString("Prompt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", pr.Deleted))
	builder.WriteString(", ")
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", pr.UUID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.UserID))
	builder.WriteString(", ")
	builder.WriteString("act=")
	builder.WriteString(pr.Act)
	builder.WriteString(", ")
	builder.WriteString("prompt=")
	builder.WriteString(pr.Prompt)
	builder.WriteString(", ")
	builder.WriteString("is_custom=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsCustom))
	builder.WriteByte(')')
	return builder.String()
}

// Prompts is a parsable slice of Prompt.
type Prompts []*Prompt
