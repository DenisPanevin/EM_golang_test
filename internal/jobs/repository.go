package jobs

import (
	"EM-Api-testTask/internal/models"
	"context"
)

type Repository interface {
	AddJob(ctx context.Context, dto *models.AddJobDto) (error, *int64)
	StopJob(ctx context.Context, id *int64) error
	Get(ctx context.Context) (error, *models.Job)
	CheckExist(ctx context.Context, dto *models.AddJobDto) (*models.Job, error)
}
