package models

import (
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/utils"
)

type ActivityHistory struct {
	ID          uint       `bson:"_id,omitempty" json:"id"`
	VideoID     uint       `bson:"video_id" json:"video_id"`
	Note        string     `bson:"note,omitempty" json:"note"`
	LevelSystem string     `bson:"level_system" json:"level_system"`
	Actions     string     `bson:"actions" json:"actions"`
	User        string     `bson:"user" json:"user"`
	CreatedAt   time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`
}

// ActivityHistoryList list response with pagination
type ActivityHistoryList struct {
	TotalCount        int64              `json:"totalCount" bson:"totalCount"`
	TotalPages        int64              `json:"totalPages" bson:"totalPages"`
	Page              int64              `json:"page" bson:"page"`
	Size              int64              `json:"size" bson:"size"`
	HasMore           bool               `json:"hasMore" bson:"hasMore"`
	ActivityHistories []*ActivityHistory `json:"ActivityHistories" bson:"ActivityHistories"`
}

func NewActivityHistoryListWithPagination(activityHistories []*ActivityHistory, count int64, pagination *utils.Pagination) *ActivityHistoryList {
	return &ActivityHistoryList{
		TotalCount:        count,
		TotalPages:        int64(pagination.GetTotalPages(int(count))),
		Page:              int64(pagination.GetPage()),
		Size:              int64(pagination.GetSize()),
		HasMore:           pagination.GetHasMore(int(count)),
		ActivityHistories: activityHistories,
	}
}
