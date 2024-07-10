package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	"context"
	"net/http"
	"os/user"
)

type TasksUseCase struct {
	repo tasks.Repository
}

func (t TasksUseCase) Create(ctx context.Context, dto *models.CreateTaskDto) (error, *int64) {
	err, id := t.repo.CreateTask(ctx, dto)
	if err != nil {
		return err, nil
	}
	return nil, id
}

func (t TasksUseCase) Get(r *http.Request) (error, *models.User) {
	//TODO implement me
	panic("implement me")
}

func (t TasksUseCase) Edit(ctx context.Context) (error, *user.User) {
	//TODO implement me
	panic("implement me")
}

func (t TasksUseCase) DeleteTask(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewTasksUseCase(ur tasks.Repository) tasks.UseCase {
	return &TasksUseCase{
		repo: ur,
	}
}
