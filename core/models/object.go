package models

import "time"

// Object represents the tbl_objects table.
type Object struct {
	ID          uint       `db:"id" json:"id"`
	VideoID     uint       `db:"video_id" json:"video_id"`
	Description string     `db:"description" json:"description"`
	CoordinateX uint       `db:"coordinate_x" json:"coordinate_x"`
	CoordinateY uint       `db:"coordinate_y" json:"coordinate_y"`
	Length      uint       `db:"length" json:"length"`
	Width       uint       `db:"width" json:"width"`
	Order       uint       `db:"order" json:"order"`
	TimePoint   uint       `db:"time_point" json:"time_point"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
