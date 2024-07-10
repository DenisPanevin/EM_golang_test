package repository

import (
	"EM-Api-testTask/internal/jobs"
	"EM-Api-testTask/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kpango/glg"
	"time"
)

func NewJobsRepo(pool *pgxpool.Pool) jobs.Repository {
	return JobsRepo{
		db: pool,
	}
}

type JobsRepo struct {
	db *pgxpool.Pool
}

func (j JobsRepo) AddJob(ctx context.Context, dto *models.AddJobDto) (error, *int64) {
	query := `INSERT INTO jobs ( user_id ,task_id,started,stopped)VALUES ($1, $2,$3,$4)RETURNING id`

	args := []interface{}{dto.UserId, dto.TaskId, time.Now(), time.Time{}}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64

	err := j.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		glg.Debugf("error creating job %s", err)
		return err, &id
	}
	return nil, &id

}

func (j JobsRepo) CheckExist(ctx context.Context, dto *models.AddJobDto) (*models.Job, error) {

	query := `SELECT id, user_id, task_id, started, stopped  FROM jobs WHERE user_id=$1 and task_id=$2 ORDER BY started DESC limit 1`
	args := []interface{}{dto.UserId, dto.TaskId}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	job := models.Job{}
	err := j.db.QueryRow(ctx, query, args...).Scan(&job.Id, &job.UserId, &job.TaskId, &job.Started, &job.Stoped)
	if err != nil {
		glg.Debugf("error in checkExist repo %s", err)
		return nil, err
	}

	return &job, nil

}
func (j JobsRepo) StopJob(ctx context.Context, id *int64) error {
	query := `
        UPDATE jobs
        SET stopped=$1
        WHERE id=$2
    `

	args := []interface{}{time.Now(), *id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	j.db.QueryRow(ctx, query, args...)
	return nil

}

func (j JobsRepo) Get(ctx context.Context) (error, *models.Job) {
	//TODO implement me
	panic("implement me")
}
