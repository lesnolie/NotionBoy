// Code generated by ent, DO NOT EDIT.

package conversation

import (
	"notionboy/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldUpdatedAt, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldDeleted, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldUUID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldUserID, v))
}

// Instruction applies equality check predicate on the "instruction" field. It's identical to InstructionEQ.
func Instruction(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldInstruction, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldTitle, v))
}

// TokenUsage applies equality check predicate on the "token_usage" field. It's identical to TokenUsageEQ.
func TokenUsage(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldTokenUsage, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldDeleted, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldUUID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldUserID, v))
}

// InstructionEQ applies the EQ predicate on the "instruction" field.
func InstructionEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldInstruction, v))
}

// InstructionNEQ applies the NEQ predicate on the "instruction" field.
func InstructionNEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldInstruction, v))
}

// InstructionIn applies the In predicate on the "instruction" field.
func InstructionIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldInstruction, vs...))
}

// InstructionNotIn applies the NotIn predicate on the "instruction" field.
func InstructionNotIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldInstruction, vs...))
}

// InstructionGT applies the GT predicate on the "instruction" field.
func InstructionGT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldInstruction, v))
}

// InstructionGTE applies the GTE predicate on the "instruction" field.
func InstructionGTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldInstruction, v))
}

// InstructionLT applies the LT predicate on the "instruction" field.
func InstructionLT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldInstruction, v))
}

// InstructionLTE applies the LTE predicate on the "instruction" field.
func InstructionLTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldInstruction, v))
}

// InstructionContains applies the Contains predicate on the "instruction" field.
func InstructionContains(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContains(FieldInstruction, v))
}

// InstructionHasPrefix applies the HasPrefix predicate on the "instruction" field.
func InstructionHasPrefix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasPrefix(FieldInstruction, v))
}

// InstructionHasSuffix applies the HasSuffix predicate on the "instruction" field.
func InstructionHasSuffix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasSuffix(FieldInstruction, v))
}

// InstructionIsNil applies the IsNil predicate on the "instruction" field.
func InstructionIsNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldIsNull(FieldInstruction))
}

// InstructionNotNil applies the NotNil predicate on the "instruction" field.
func InstructionNotNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldNotNull(FieldInstruction))
}

// InstructionEqualFold applies the EqualFold predicate on the "instruction" field.
func InstructionEqualFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEqualFold(FieldInstruction, v))
}

// InstructionContainsFold applies the ContainsFold predicate on the "instruction" field.
func InstructionContainsFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContainsFold(FieldInstruction, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Conversation {
	return predicate.Conversation(sql.FieldContainsFold(FieldTitle, v))
}

// TokenUsageEQ applies the EQ predicate on the "token_usage" field.
func TokenUsageEQ(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldEQ(FieldTokenUsage, v))
}

// TokenUsageNEQ applies the NEQ predicate on the "token_usage" field.
func TokenUsageNEQ(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldNEQ(FieldTokenUsage, v))
}

// TokenUsageIn applies the In predicate on the "token_usage" field.
func TokenUsageIn(vs ...int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldIn(FieldTokenUsage, vs...))
}

// TokenUsageNotIn applies the NotIn predicate on the "token_usage" field.
func TokenUsageNotIn(vs ...int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldNotIn(FieldTokenUsage, vs...))
}

// TokenUsageGT applies the GT predicate on the "token_usage" field.
func TokenUsageGT(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldGT(FieldTokenUsage, v))
}

// TokenUsageGTE applies the GTE predicate on the "token_usage" field.
func TokenUsageGTE(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldGTE(FieldTokenUsage, v))
}

// TokenUsageLT applies the LT predicate on the "token_usage" field.
func TokenUsageLT(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldLT(FieldTokenUsage, v))
}

// TokenUsageLTE applies the LTE predicate on the "token_usage" field.
func TokenUsageLTE(v int64) predicate.Conversation {
	return predicate.Conversation(sql.FieldLTE(FieldTokenUsage, v))
}

// TokenUsageIsNil applies the IsNil predicate on the "token_usage" field.
func TokenUsageIsNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldIsNull(FieldTokenUsage))
}

// TokenUsageNotNil applies the NotNil predicate on the "token_usage" field.
func TokenUsageNotNil() predicate.Conversation {
	return predicate.Conversation(sql.FieldNotNull(FieldTokenUsage))
}

// HasConversationMessages applies the HasEdge predicate on the "conversation_messages" edge.
func HasConversationMessages() predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConversationMessagesTable, ConversationMessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConversationMessagesWith applies the HasEdge predicate on the "conversation_messages" edge with a given conditions (other predicates).
func HasConversationMessagesWith(preds ...predicate.ConversationMessage) predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConversationMessagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConversationMessagesTable, ConversationMessagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Conversation) predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Conversation) predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Conversation) predicate.Conversation {
	return predicate.Conversation(func(s *sql.Selector) {
		p(s.Not())
	})
}
