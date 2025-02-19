package models

import "github.com/glu/video-real-time-ranking/ent"

// ObjectsListResponse comments list response with pagination
type ObjectsListResponse struct {
	TotalCount int            `json:"totalCount"`
	TotalPages int            `json:"totalPages"`
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	HasMore    bool           `json:"hasMore"`
	Objects    []*ent.Objects `json:"objects"`
}

type ObjectInteractiveResponse struct {
	Name        string `json:"name"`
	MakerName   string `json:"maker_name"`
	In          int    `json:"in"`
	Out         int    `json:"out"`
	TouchVector string `json:"touch_vector"`
}
