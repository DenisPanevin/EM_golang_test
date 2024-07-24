package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	Apierrors "EM-Api-testTask/pkg"
	Helpers "EM-Api-testTask/pkg/helpers"
	"context"
	"github.com/kpango/glg"
	"net/url"
	"strconv"
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

func (t TasksUseCase) GetAll(ctx context.Context, vals url.Values) (error, *[]models.Task) {
	filters := models.Task{}

	if vals.Get("id") != "" {
		id, err := strconv.Atoi(vals.Get("id"))
		if err != nil || id <= 0 {
			glg.Debugf("error parsing id in task filters %s", err)
			return Apierrors.ApiWrongInput, nil
		} else {
			filters.Id = int64(id)
		}
	}

	filters.Name = vals.Get("name")

	v := Helpers.NewValidator()
	err := v.Validate(filters)
	if err != nil {
		glg.Debugf("error validating filters %s", err)
		return Apierrors.ApiWrongInput, nil
	}

	pageFilters := models.PageFiltersDto{
		Limit: 50,
		Page:  1,
	}
	if vals.Get("limit") != "" {
		limit, e := strconv.Atoi(vals.Get("limit"))
		if e != nil || limit <= 0 {
			glg.Debugf("error parsing page in Pagefilters %s", e)
			return Apierrors.ApiWrongInput, nil
		} else {
			pageFilters.Limit = limit
		}
	}
	if vals.Get("page") != "" {
		page, er := strconv.Atoi(vals.Get("page"))
		if er != nil || page <= 0 {
			glg.Debugf("error parsing page in Pagefilters %s", er)
			return Apierrors.ApiWrongInput, nil
		} else {
			pageFilters.Page = page
		}
	}

	pageFilters.CalculateOffset()

	err, u := t.repo.GetAll(ctx, filters, pageFilters)
	if err != nil {
		glg.Debugf("cant get all from task repo in usecase %s", err)
		return err, nil
	}
	return nil, u

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
