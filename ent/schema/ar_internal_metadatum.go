// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type ArInternalMetadatum struct {
	ent.Schema
}

func (ArInternalMetadatum) Fields() []ent.Field {
	return []ent.Field{field.String("id").StorageKey("key"), field.String("value").Optional(), field.Time("created_at"), field.Time("updated_at")}
}
func (ArInternalMetadatum) Edges() []ent.Edge {
	return nil
}
func (ArInternalMetadatum) Annotations() []schema.Annotation {
	return nil
}
