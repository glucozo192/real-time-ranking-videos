package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/utils"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
)

type IActivityHistoryRepositoryReader interface {
	CreateActivityHistory(ctx context.Context, activityHistory *models.ActivityHistory) (*models.ActivityHistory, error)
	UpdateActivityHistory(ctx context.Context, activityHistory *models.ActivityHistory) (*models.ActivityHistory, error)
	DeleteActivityHistory(ctx context.Context, id uint) error

	GetActivityHistoryById(ctx context.Context, id uint) (*models.ActivityHistory, error)
	Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.ActivityHistoryList, error)
}
