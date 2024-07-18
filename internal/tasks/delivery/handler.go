package delivery

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	Apierrors "EM-Api-testTask/pkg"
	Helpers "EM-Api-testTask/pkg/helpers"
	helpers "EM-Api-testTask/pkg/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kpango/glg"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase tasks.UseCase
}

func NewHandler(uc tasks.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

// @Summary      Creates new task
// @Description  Creates new task
// @Accept       json
// @Produce      json
// @Param request body models.CreateTaskDto	true	"new task data"
//
//	@Success 201  {object} models.Task "returning task object"
//	@Failure 400	{object}	helpers.ErrorDto
//	@Failure 422	{object}	helpers.ErrorDto
//	@Failure 500	{object}	helpers.ErrorDto
//
// @Router       /tasks [post]
func (h *Handler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(models.CreateTaskDto)
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			glg.Debugf("task decoder %s", err)
			helpers.SendError(w, r, http.StatusBadRequest, err)
			return
		}
		v := Helpers.NewValidator()
		if err := v.Validate(dto); err != nil {
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)
			return
		}
		err, task := h.useCase.Create(r.Context(), dto)
		if err != nil {
			helpers.SendError(w, r, http.StatusInternalServerError, err)
			return
		}

		helpers.SendData(w, r, http.StatusOK, task)

	}

}

// @Summary      Delete task
// @Description  Deletes task by id
// @Param id path int true "Task ID"
//
//	@Success 200  {string }	string	 "{Deleted id: id}"
//	@Failure 500	{object}	helpers.ErrorDto
//	@Failure 422	{object}	helpers.ErrorDto
//
// @Router       /tasks/{id} [delete]
func (h *Handler) DeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			glg.Logf("cant parse id from delete user handler %s", err)
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)

			return
		}
		err = h.useCase.DeleteTask(r.Context(), id)
		if err != nil {
			glg.Debugf("error in delete usecase returning %s", err)
			helpers.SendError(w, r, http.StatusInternalServerError, err)
			return
		}
		helpers.SendData(w, r, http.StatusOK, fmt.Sprintf(fmt.Sprintf(`{Deleted id: %v}`, id)))
	}

}
