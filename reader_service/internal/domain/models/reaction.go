package models

import "github.com/glu/video-real-time-ranking/ent"

// ReactionsListResponse comments list response with pagination
type ReactionsListResponse struct {
	TotalCount int              `json:"totalCount"`
	TotalPages int              `json:"totalPages"`
	Page       int              `json:"page"`
	Size       int              `json:"size"`
	HasMore    bool             `json:"hasMore"`
	Reactions  []*ent.Reactions `json:"reactions"`
}
