package models

import "time"

// Video represents the tbl_videos table.
type Video struct {
	ID           uint       `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Description  string     `db:"description" json:"description"`
	VideoURL     string     `db:"video_url" json:"video_url"`
	Config       string     `db:"config" json:"config"`
	PathResource string     `db:"path_resource" json:"path_resource"`
	LevelSystem  string     `db:"level_system" json:"level_system"`
	Status       string     `db:"status" json:"status"`
	Note         string     `db:"note" json:"note"`
	Assign       string     `db:"assign" json:"assign"`
	Author       string     `db:"author" json:"author"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
}
