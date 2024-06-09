package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("url"),
		field.String("account_number").
			MinLen(12).
			MaxLen(12).
			Default(""),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return nil
}
