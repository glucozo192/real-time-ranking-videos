package models

import "time"

// Comment represents the tbl_comments table.
type Comment struct {
	ID          uint       `db:"id" json:"id"`
	VideoID     uint       `db:"video_id" json:"video_id"`
	Description string     `db:"description" json:"description"`
	Comment     string     `db:"comment" json:"comment"`
	UserName    string     `db:"user_name" json:"user_name"`
	Avatar      string     `db:"avatar" json:"avatar"`
	TimePoint   uint       `db:"time_point" json:"time_point"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
