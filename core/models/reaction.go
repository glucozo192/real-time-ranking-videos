package models

import "time"

// Reaction represents the tbl_reactions table.
type Reaction struct {
	ID          uint       `db:"id" json:"id"`
	VideoID     uint       `db:"video_id" json:"video_id"`
	Description string     `db:"description" json:"description"`
	Name        string     `db:"name" json:"name"`
	Number      uint       `db:"number" json:"number"`
	TimePoint   uint       `db:"time_point" json:"time_point"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
