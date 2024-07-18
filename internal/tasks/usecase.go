package tasks

import (
	"EM-Api-testTask/internal/models"
	"context"
	"net/http"
)

type UseCase interface {
	Create(ctx context.Context, dto *models.CreateTaskDto) (error, *models.Task)
	Get(r *http.Request) (error, *models.User)

	DeleteTask(ctx context.Context, id int) error
}
