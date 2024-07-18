package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	"context"
	"net/http"
)

type TasksUseCase struct {
	repo tasks.Repository
}

func (t TasksUseCase) Create(ctx context.Context, dto *models.CreateTaskDto) (error, *models.Task) {
	err, task := t.repo.CreateTask(ctx, dto)
	if err != nil {
		return err, nil
	}
	return nil, task
}

func (t TasksUseCase) Get(r *http.Request) (error, *models.User) {
	//TODO implement me
	panic("implement me")
}

func (t *TasksUseCase) DeleteTask(ctx context.Context, id int) error {
	err := t.repo.DeleteTask(ctx, id)
	return err
}

func NewTasksUseCase(ur tasks.Repository) tasks.UseCase {
	return &TasksUseCase{
		repo: ur,
	}
}
