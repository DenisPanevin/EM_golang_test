package repository

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	"EM-Api-testTask/pkg"
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

func (t TasksRepo) CreateTask(ctx context.Context, dto *models.CreateTaskDto) (error, *models.Task) {
	query := `INSERT INTO tasks ( name  )VALUES ($1)RETURNING id,name`

	args := []interface{}{dto.Name}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var task models.Task
	err := t.db.QueryRow(ctx, query, args...).Scan(&task.Id, &task.Name)
	if err != nil {
		glg.Debugf("error add task to psql %s", err)
		return err, nil
	}
	return nil, &task
}

func (t TasksRepo) Get(ctx context.Context) (error, *models.User) {
	//TODO implement me
	panic("implement me")
}
func (t TasksRepo) DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM tasks
WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tag, err := t.db.Exec(ctx, query, id)
	if err != nil {
		glg.Debugf("error delete to psql %s", err)
		return err
	}
	if tag.RowsAffected() != 1 {
		println(tag.RowsAffected())
		return pkg.NotFound
	}

	return nil

}
