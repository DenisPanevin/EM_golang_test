package tasks

import (
	"EM-Api-testTask/internal/models"
	"context"
	"net/url"
)

type UseCase interface {
	Create(ctx context.Context, dto *models.CreateTaskDto) (error, *models.Task)
	GetAll(ctx context.Context, v url.Values) (error, *[]models.Task)

	DeleteTask(ctx context.Context, id int) error
}
