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
func (j JobsRepo) GetJob(ctx context.Context, filters models.UserFiltersDto, interval models.JobIntervalDto, pagefilters models.PageFiltersDto) (error, *[]models.ShowUserDto) {

	query := `SELECT
    
    j.user_id ,
    u.name AS user_name,u.surname,u.patronymic,u.address,u.passportnumber,t.id,t.name,
    SUM(CASE WHEN j.stopped - j.started >= INTERVAL '0' THEN j.stopped - j.started ELSE INTERVAL '0' END) AS total_work 
FROM
    jobs j
        JOIN
    users u ON j.user_id = u.id
        JOIN
    tasks t ON j.task_id = t.id
WHERE
    (u.name=$1 or $1='') and (u.surname=$2 or $2='') and (u.patronymic=$3 or $3='')and (u.id=$4 or $4=0)
AND ((j.started >= $5::timestamp AND j.stopped <=$6::timestamp)or task_id=0)
GROUP BY
     j.user_id, u.name,u.surname,u.patronymic,t.name,u.address,u.passportnumber,t.id,t.name
  ORDER BY
    total_work DESC
limit $7 offset $8
 ;
    `

	rows, err := j.db.Query(ctx, query, filters.Name, filters.Surname, filters.Patronymic, filters.Id, interval.DateStart, interval.DateEnd, pagefilters.Limit, pagefilters.Offset)
	if err != nil {
		glg.Debugf("err from psql GET %s", err)
		return err, nil
	}
	defer rows.Close()

	usrs := make(map[int64]models.ShowUserDto)
	for rows.Next() {
		usr := models.ShowUserDto{}
		tsk := models.ShowUserTaskDto{}
		err = rows.Scan(&usr.UserId, &usr.Name, &usr.Surname, &usr.Patronymic, &usr.Address, &usr.PassportNumber, &tsk.TaskId, &tsk.TaskName, &tsk.TotalWork)
		if err != nil {
			glg.Debugf("Row scan failed: %v\n", err)
			return err, nil
		}
		//println(fmt.Sprintf("%v %s %v %s", usr.UserId, usr.Name, tsk.TaskId, tsk.TaskName))
		if existingUser, ok := usrs[usr.UserId]; ok {
			if tsk.TaskId != 0 {
				existingUser.Jobs = append(existingUser.Jobs, tsk)
			}
			usrs[usr.UserId] = existingUser
		} else {
			if tsk.TaskId != 0 {
				usr.Jobs = []models.ShowUserTaskDto{tsk}
			}

			usrs[usr.UserId] = usr
		}
	}

	var usersOutput []models.ShowUserDto
	for _, v := range usrs {
		usersOutput = append(usersOutput, v)
	}

	return nil, &usersOutput

}
