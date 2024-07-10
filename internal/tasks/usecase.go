package tasks

import (
	"EM-Api-testTask/internal/models"
	"context"
	"net/http"
	"os/user"
)

type UseCase interface {
	Create(ctx context.Context, dto *models.CreateTaskDto) (error, *int64)
	Get(r *http.Request) (error, *models.User)

	Edit(ctx context.Context) (error, *user.User)

	DeleteTask(ctx context.Context) error
}
