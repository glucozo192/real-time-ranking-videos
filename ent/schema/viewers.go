package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Viewers holds the schema definition for the Viewers entity.
type Viewers struct {
	ent.Schema
}

// Annotations of the Viewers.
func (Viewers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_viewers"},
	}
}

// Fields of the Viewers.
func (Viewers) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Positive().StructTag(`json:"id" bson:"_id,omitempty"`),
		field.Uint("video_id").
			Positive().
			StructTag(`json:"video_id" bson:"video_id"`), // JSON and BSON struct tags for the 'video_id' field.
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

// Edges of the Viewers.
func (Viewers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tbl_videos", Videos.Type).
			Ref("tbl_viewers").
			Field("video_id").
			Unique().
			Required(),
	}
}
