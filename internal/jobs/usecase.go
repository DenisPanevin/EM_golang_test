package jobs

import (
	"EM-Api-testTask/internal/models"
	"context"
)

type UseCase interface {
	AddJob(ctx context.Context, dto *models.AddJobDto) (error, *int64)
	DeleteJob(ctx context.Context, id int64) error
}
