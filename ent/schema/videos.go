package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Videos holds the schema definition for the Videos entity.
type Videos struct {
	ent.Schema
}

// Annotations of the User.
func (Videos) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_videos"},
	}
}

// Fields of the Videos.
func (Videos) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			StructTag(`json:"id" bson:"_id,omitempty"`), // JSON and BSON struct tags for the 'id' field.
		field.String("name").
			StructTag(`json:"name" bson:"name"`), // JSON and BSON struct tags for the 'name' field.
		field.Text("description").
			Default("").
			Optional().
			StructTag(`json:"description" bson:"description"`), // JSON and BSON struct tags for the 'description' field.
		field.String("video_url").
			StructTag(`json:"video_url" bson:"video_url"`), // JSON and BSON struct tags for the 'video_url' field.
		field.Text("config").
			StructTag(`json:"config" bson:"config"`), // JSON and BSON struct tags for the 'config' field.
		field.Text("path_resource").
			StructTag(`json:"path_resource" bson:"path_resource"`), // JSON and BSON struct tags for the 'path_resource' field.
		field.String("level_system").
			StructTag(`json:"level_system" bson:"level_system"`), // JSON and BSON struct tags for the 'level_system' field.
		field.String("status").
			StructTag(`json:"status" bson:"status"`), // JSON and BSON struct tags for the 'status' field.
		field.Text("note").
			StructTag(`json:"note" bson:"note"`), // JSON and BSON struct tags for the 'note' field.
		field.String("assign").
			StructTag(`json:"assign" bson:"assign"`), // JSON and BSON struct tags for the 'assign' field.
		field.Uint("version").
			Default(1).
			StructTag(`json:"version" bson:"version,omitempty"`), // JSON and BSON struct tags for the 'version' field.
		field.String("Author").
			StructTag(`json:"author" bson:"author"`), // JSON and BSON struct tags for the 'Author' field.
		field.Time("created_at").
			Default(time.Now).
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

// Edges of the Videos.
func (Videos) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tbl_comments", Comments.Type),
		edge.To("tbl_reactions", Reactions.Type),
		edge.To("tbl_viewers", Viewers.Type),
		edge.To("tbl_objects", Objects.Type),
	}
}
