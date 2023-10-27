// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type SchemaMigration struct {
	ent.Schema
}

func (SchemaMigration) Fields() []ent.Field {
	return []ent.Field{field.String("id").StorageKey("version")}
}
func (SchemaMigration) Edges() []ent.Edge {
	return nil
}
func (SchemaMigration) Annotations() []schema.Annotation {
	return nil
}
