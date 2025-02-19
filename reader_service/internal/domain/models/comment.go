package models

import (
	"github.com/glu/video-real-time-ranking/ent"
)

// CommentsListResponse comments list response with pagination
type CommentsListResponse struct {
	TotalCount int             `json:"totalCount"`
	TotalPages int             `json:"totalPages"`
	Page       int             `json:"page"`
	Size       int             `json:"size"`
	HasMore    bool            `json:"hasMore"`
	Comments   []*ent.Comments `json:"comments"`
}
