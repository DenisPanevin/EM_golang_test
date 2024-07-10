package delivery

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/tasks"
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"net/http"
)

type Handler struct {
	useCase tasks.UseCase
}

func NewHandler(uc tasks.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(models.CreateTaskDto)
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			glg.Debugf("task decoder", err)
			h.customErr(w, r, http.StatusBadRequest, err)
			return
		}

		/*if err := validators.Validate(*dto); err != nil {
			glg.Debugf("task validator", err)
			h.customErr(w, r, http.StatusUnprocessableEntity, internal.ApiWrongInput)
			return
		}*/

		err, id := h.useCase.Create(r.Context(), dto)
		if err != nil {
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}

		h.customRespond(w, r, http.StatusOK, fmt.Sprintf("Created! task id is %v", *id))

	}

}
func (h *Handler) customErr(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.customRespond(w, r, code, map[string]string{"error:": err.Error()})
}
func (h *Handler) customRespond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
