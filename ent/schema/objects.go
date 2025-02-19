package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Objects holds the schema definition for the Objects entity.
type Objects struct {
	ent.Schema
}

// Annotations of the Objects.
func (Objects) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_objects"},
	}
}

// Fields of the Objects.
func (Objects) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Positive().StructTag(`json:"id" bson:"_id,omitempty"`),
		field.Uint("video_id").
			StructTag(`json:"video_id" bson:"video_id"`), // JSON and BSON struct tags for the 'video_id' field.
		field.Text("description").
			Default("").
			Optional().
			StructTag(`json:"description" bson:"description"`), // JSON and BSON struct tags for the 'description' field.
		field.Int("coordinate_x").
			Positive().
			StructTag(`json:"coordinate_x" bson:"coordinate_x"`), // JSON and BSON struct tags for the 'coordinate_x' field.
		field.Int("coordinate_y").
			Positive().
			StructTag(`json:"coordinate_y" bson:"coordinate_y"`), // JSON and BSON struct tags for the 'coordinate_y' field.
		field.Int("length").
			Positive().
			StructTag(`json:"length" bson:"length"`), // JSON and BSON struct tags for the 'length' field.
		field.Int("width").
			Positive().
			StructTag(`json:"width" bson:"width"`), // JSON and BSON struct tags for the 'width' field.
		field.Int("order").
			Positive().
			StructTag(`json:"order" bson:"order"`), // JSON and BSON struct tags for the 'order' field.
		field.Float("time_start").
			StructTag(`json:"time_start" bson:"time_start"`), // JSON and BSON struct tags for the 'time_start' field.
		field.Float("time_end").
			StructTag(`json:"time_end" bson:"time_end"`), // JSON and BSON struct tags for the 'time_end' field.
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
		field.Text("touch_vector").
			Default("").
			StructTag(`json:"touch_vector" bson:"touch_vector"`), // JSON and BSON struct tags for the 'touch_vector' field.
		field.Text("marker_name").
			Default("").
			StructTag(`json:"marker_name" bson:"marker_name"`), // JSON and BSON struct tags for the 'marker_name' field.
		field.Float("time_point").
			Default(0).
			StructTag(`json:"time_point" bson:"time_point"`), // JSON and BSON struct tags for the 'time_point' field.
	}
}

// Edges of the Objects.
func (Objects) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tbl_videos", Videos.Type).
			Ref("tbl_objects").
			Field("video_id").
			Unique().
			Required(),
	}
}
