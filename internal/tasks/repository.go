package tasks

import (
	"EM-Api-testTask/internal/models"
	"context"
)

type Repository interface {
	CreateTask(ctx context.Context, dto *models.CreateTaskDto) (error, *int64)
	Get(ctx context.Context) (error, *models.User)
}
