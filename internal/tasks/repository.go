package tasks

import (
	"EM-Api-testTask/internal/models"
	"context"
)

type Repository interface {
	CreateTask(ctx context.Context, dto *models.CreateTaskDto) (error, *models.Task)
	Get(ctx context.Context) (error, *models.User)
	DeleteTask(ctx context.Context, id int) error
}
