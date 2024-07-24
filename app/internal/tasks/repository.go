package tasks

import (
	"EM-Api-testTask/internal/models"
	"context"
)

type Repository interface {
	CreateTask(ctx context.Context, dto *models.CreateTaskDto) (error, *models.Task)
	GetAll(ctx context.Context, filters models.Task, pageFilters models.PageFiltersDto) (error, *[]models.Task)
	DeleteTask(ctx context.Context, id int) error
}
