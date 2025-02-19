package models

import "github.com/glu/video-real-time-ranking/ent"

// ViewerListResponse comments list response with pagination
type ViewerListResponse struct {
	TotalCount int            `json:"totalCount"`
	TotalPages int            `json:"totalPages"`
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	HasMore    bool           `json:"hasMore"`
	Viewers    []*ent.Viewers `json:"viewers"`
}
