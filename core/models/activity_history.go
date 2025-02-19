package models

import "time"

// ActivityHistory represents the tbl_activity_histories table.
type ActivityHistory struct {
	ID          uint       `db:"id" json:"id"`
	VideoID     uint       `db:"video_id" json:"video_id"`
	Note        string     `db:"note" json:"note"`
	LevelSystem string     `db:"level_system" json:"level_system"`
	Actions     string     `db:"actions" json:"actions"`
	User        string     `db:"user" json:"user"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
