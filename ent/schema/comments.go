package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Comments holds the schema definition for the Comments entity.
type Comments struct {
	ent.Schema
}

// Annotations of the Comments.
func (Comments) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_comments"},
	}
}

// Fields of the Comments.
func (Comments) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Positive().StructTag(`json:"id" bson:"_id,omitempty"`),
		field.Uint("video_id").
			Positive().
			StructTag(`json:"video_id" bson:"video_id"`), // JSON and BSON struct tags for the 'video_id' field.
		field.Text("description").
			Default("").
			Optional().
			StructTag(`json:"description" bson:"description"`), // JSON and BSON struct tags for the 'description' field.
		field.Text("comment").
			Default("").
			NotEmpty().
			StructTag(`json:"comment" bson:"comment"`), // JSON and BSON struct tags for the 'comment' field.
		field.String("user_name").
			MaxLen(100).
			NotEmpty().
			StructTag(`json:"user_name" bson:"user_name"`), // JSON and BSON struct tags for the 'user_name' field.
		field.String("avatar").
			Default("").
			NotEmpty().
			StructTag(`json:"avatar" bson:"avatar"`), // JSON and BSON struct tags for the 'avatar' field.
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

// Edges of the Comments.
func (Comments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tbl_videos", Videos.Type).
			Ref("tbl_comments").
			Field("video_id").
			Unique().
			Required(),
	}
}
