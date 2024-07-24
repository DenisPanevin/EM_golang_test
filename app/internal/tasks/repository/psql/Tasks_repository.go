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

func (t TasksRepo) GetAll(ctx context.Context, filters models.Task, pageFilters models.PageFiltersDto) (error, *[]models.Task) {
	query := `select id,name
   FROM
    tasks
    where    
 (id=$1 or $1=0) and (name=$2 or $2='')and(name !='idle')
ORDER BY
    id DESC
    limit $3 offset $4`

	rows, err := t.db.Query(ctx, query, filters.Id, filters.Name, pageFilters.Limit, pageFilters.Offset)
	if err != nil {
		glg.Debugf("err from psql GET all tasks %s", err)
		return err, nil
	}
	defer rows.Close()
	var tasksOutput []models.Task
	for rows.Next() {
		task := models.Task{}

		err = rows.Scan(&task.Id, &task.Name)
		if err != nil {
			glg.Debugf("Row scan failed: %v\n", err)
			return err, nil
		}
		tasksOutput = append(tasksOutput, task)
	}
	if len(tasksOutput) < 1 {
		return pkg.NotFound, nil
	}

	return nil, &tasksOutput
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
