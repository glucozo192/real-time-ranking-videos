package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Reactions holds the schema definition for the Reactions entity.
type Reactions struct {
	ent.Schema
}

// Annotations of the Reactions.
func (Reactions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_reactions"},
	}
}

// Fields of the Reactions.
func (Reactions) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Positive().StructTag(`json:"id" bson:"_id,omitempty"`),
		field.Uint("video_id").
			StructTag(`json:"video_id" bson:"video_id"`), // JSON and BSON struct tags for the 'video_id' field.
		field.Text("description").
			Default("").
			Optional().
			StructTag(`json:"description" bson:"description"`), // JSON and BSON struct tags for the 'description' field.
		field.String("name").
			MaxLen(50).
			NotEmpty().
			StructTag(`json:"name" bson:"name"`), // JSON and BSON struct tags for the 'name' field.
		field.Int("number").
			Positive().
			StructTag(`json:"number" bson:"number"`), // JSON and BSON struct tags for the 'number' field.
		field.Float("time_point").
			Positive().
			StructTag(`json:"time_point" bson:"time_point"`), // JSON and BSON struct tags for the 'time_point' field.
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			StructTag(`json:"created_at" bson:"created_at"`), // JSON and BSON struct tags for the 'created_at' field.
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			StructTag(`json:"updated_at" bson:"updated_at"`), // JSON and BSON struct tags for the 'updated_at' field.
		field.Time("deleted_at").
			Optional().
			Nillable().
			StructTag(`json:"deleted_at" bson:"deleted_at"`), // JSON and BSON struct tags for the 'deleted_at' field.
	}
}

// Edges of the Reactions.
func (Reactions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tbl_videos", Videos.Type).
			Ref("tbl_reactions").
			Field("video_id").
			Unique().
			Required(),
	}
}
