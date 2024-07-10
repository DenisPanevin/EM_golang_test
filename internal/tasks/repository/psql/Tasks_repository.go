package repository

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kpango/glg"
	"time"
)

func NewTasksRepo(pool *pgxpool.Pool) tasks.Repository {
	return TasksRepo{
		db: pool,
	}
}

type TasksRepo struct {
	db *pgxpool.Pool
}

func (t TasksRepo) CreateTask(ctx context.Context, dto *models.CreateTaskDto) (error, *int64) {
	query := `INSERT INTO tasks ( name ,status )VALUES ($1, $2)RETURNING id`

	args := []interface{}{dto.Name, dto.Status}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := t.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		glg.Debugf("error add task to psql", err)
		return err, nil
	}
	return nil, &id
}

func (t TasksRepo) Get(ctx context.Context) (error, *models.User) {
	//TODO implement me
	panic("implement me")
}
