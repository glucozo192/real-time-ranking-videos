package models

import (
	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/ent"
)

// HTTP
type SearchVideoRequest struct {
	ID          *int              `json:"id"`
	Name        *string           `json:"name"`
	Author      *string           `json:"author"`
	Assign      *string           `json:"assign"`
	LevelSystem *string           `json:"level_system"`
	Pagination  *utils.Pagination `json:"pagination"`
}

// VideosListResponse videos list response with pagination
type VideosListResponse struct {
	TotalCount int           `json:"totalCount"`
	TotalPages int           `json:"totalPages"`
	Page       int           `json:"page"`
	Size       int           `json:"size"`
	HasMore    bool          `json:"hasMore"`
	Videos     []*ent.Videos `json:"videos"`
}
