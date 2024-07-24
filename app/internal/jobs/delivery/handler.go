package delivery

import (
	"EM-Api-testTask/internal/jobs"
	"EM-Api-testTask/internal/models"
	Apierrors "EM-Api-testTask/pkg"
	Helpers "EM-Api-testTask/pkg/helpers"
	helpers "EM-Api-testTask/pkg/http"
	"encoding/json"
	"github.com/kpango/glg"
	"net/http"
)

type Handler struct {
	useCase jobs.UseCase
}

func NewHandler(uc jobs.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

// @Summary      Assign user to a task...
// @Description  Assign user to a task by user id and task id first post starts the task and second stops it
// @Accept       json
// @Param request body models.AddJobDto	true	"{"userid":1,"taskid":1}"
//
//	@Success 200   "status OK"
//	@Failure 400	{object}	helpers.ErrorDto
//	@Failure 422	{object}	helpers.ErrorDto
//	@Failure 500	{object}	helpers.ErrorDto
//
// @Router       /jobs [post]
func (h *Handler) AddJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(models.AddJobDto)
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			glg.Debugf("job decoder %s", err)
			helpers.SendError(w, r, http.StatusBadRequest, err)
			return
		}
		v := Helpers.NewValidator()
		if err := v.Validate(dto); err != nil {

			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)
			return
		}

		err, _ := h.useCase.AddJob(r.Context(), dto)
		if err != nil {
			glg.Debugf("job usecase %s", err)
			helpers.SendError(w, r, http.StatusInternalServerError, err)
			return
		}

		helpers.SendData(w, r, http.StatusOK, nil)

	}

}

// @Summary      Get list of jobs for user by user id
// @Description  return list of jobs for user by user id
// @Produce      json
// @Param id query int false "User ID"
// @Param page query int false "page"
// @Param limit query int false "limit"
//
//	@Success 200  {object} models.ShowUserDto{TotalWorkTime=nil} "returning Users objects"
//	@Failure 500	{object}	helpers.ErrorDto
//
// @Router       /jobs/user [get]
func (h *Handler) GetJobs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err, u := h.useCase.GetJobs(r.Context(), r.URL.Query())

		if err != nil {
			helpers.SendError(w, r, 500, err)
			return
		}

		helpers.SendData(w, r, http.StatusOK, u)

	}

}
