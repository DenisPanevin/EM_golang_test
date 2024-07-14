package usecase

import (
	"EM-Api-testTask/internal/jobs"
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/pkg/handler"
	Helpers "EM-Api-testTask/pkg/helpers"
	"context"
	"errors"
	"github.com/kpango/glg"
	"net/url"
	"strconv"
	"time"
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
func (t *TasksUseCase) GetJobs(ctx context.Context, vals url.Values) (error, *[]models.ShowUserDto) {

	filters := models.UserFiltersDto{}

	if vals.Get("id") != "" {
		id, err := strconv.Atoi(vals.Get("id"))
		if err != nil || id <= 0 {
			glg.Debugf("error parsing id in filters %s", err)
			return handler.ApiWrongInput, nil
		} else {
			filters.Id = id
		}
	}

	filters.PassportNumber = vals.Get("passportnumber")

	if filters.Id == 0 && filters.PassportNumber == "" {
		err := errors.New("Id or passport numberr need to be provided")
		return err, nil
	}

	v := Helpers.NewValidator()
	err := v.Validate(filters)
	if err != nil {
		glg.Debugf("error validating filters %s", err)
		return handler.ApiWrongInput, nil
	}

	pageFilters := models.PageFiltersDto{
		Limit: 50,
		Page:  1,
	}
	if vals.Get("limit") != "" {
		limit, e := strconv.Atoi(vals.Get("limit"))
		if e != nil || limit <= 0 {
			glg.Debugf("error parsing page in Pagefilters %s", e)
			return handler.ApiWrongInput, nil
		} else {
			pageFilters.Limit = limit
		}
	}
	if vals.Get("page") != "" {
		page, er := strconv.Atoi(vals.Get("page"))
		if er != nil || page <= 0 {
			glg.Debugf("error parsing page in Pagefilters %s", er)
			return handler.ApiWrongInput, nil
		} else {
			pageFilters.Page = page
		}
	}

	jobInterval := models.JobIntervalDto{}

	if vals.Get("started") != "" {
		started, e := time.Parse("2006.01.02", vals.Get("started"))
		if e != nil {
			glg.Debugf("cant parse started in usecase %s", e)
			return e, nil
		}
		jobInterval.DateStart = started
	}
	ti, _ := time.Parse("2006.01.02", "9999.12.31")
	jobInterval.DateEnd = ti
	if vals.Get("ended") != "" {
		ended, e := time.Parse("2006.01.02", vals.Get("started"))
		if e != nil {
			glg.Debugf("cant parse ended in usecase %s", e)

			return e, nil
		}
		jobInterval.DateEnd = ended
	}

	pageFilters.CalculateOffset()

	err, u := t.repo.GetJob(ctx, filters, jobInterval, pageFilters)
	if err != nil {
		glg.Debugf("cant get from getJob repo in usecase %s", err)
		return err, nil
	}
	return nil, u

}
