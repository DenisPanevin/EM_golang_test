package usecase

import (
	"EM-Api-testTask/internal/jobs"
	"EM-Api-testTask/internal/models"
	"context"
	"github.com/kpango/glg"
)

type TasksUseCase struct {
	repo jobs.Repository
}

func NewJobsUseCase(ur jobs.Repository) jobs.UseCase {
	return &TasksUseCase{
		repo: ur,
	}
}

func (t TasksUseCase) AddJob(ctx context.Context, dto *models.AddJobDto) (error, *int64) {

	job, err := t.repo.CheckExist(ctx, dto)
	if err != nil {
		switch err.Error() {
		case "no rows in result set":
			e, id := t.repo.AddJob(ctx, dto)
			if e != nil {
				return e, nil
			}
			return nil, id
		default:
			glg.Debugf("edit job query error %s", err)
			return err, nil
		}
	}
	if job.Stoped.Day() == 1 && job.Stoped.Year() == 0001 && job.Stoped.Month() == 1 {
		e := t.repo.StopJob(ctx, &job.Id)
		if e != nil {
			return e, nil
		}
		return nil, &job.Id
	}

	e, id := t.repo.AddJob(ctx, dto)
	if e != nil {
		return e, nil
	}
	return nil, id

}

func (t TasksUseCase) DeleteJob(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
