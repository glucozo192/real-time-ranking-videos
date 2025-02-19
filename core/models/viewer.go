package models

import "time"

// Viewer represents the tbl_viewers table.
type Viewer struct {
	ID        uint       `db:"id" json:"id"`
	VideoID   uint       `db:"video_id" json:"video_id"`
	Number    uint       `db:"number" json:"number"`
	TimePoint uint       `db:"time_point" json:"time_point"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
