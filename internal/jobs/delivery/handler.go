package delivery

import (
	"EM-Api-testTask/internal/jobs"
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/pkg/handler"
	Helpers "EM-Api-testTask/pkg/helpers"
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

func (h *Handler) AddJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(models.AddJobDto)
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			glg.Debugf("job decoder %s", err)
			h.customErr(w, r, http.StatusBadRequest, err)
			return
		}
		v := Helpers.NewValidator()
		if err := v.Validate(dto); err != nil {

			h.customErr(w, r, http.StatusUnprocessableEntity, handler.ApiWrongInput)
			return
		}

		err, _ := h.useCase.AddJob(r.Context(), dto)
		if err != nil {
			glg.Debugf("job usecase %s", err)
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}

		//		h.customRespond(w, r, http.StatusOK, fmt.Sprintf("Added! Job id is %v", *id))

	}

}
func (h *Handler) GetJobs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err, u := h.useCase.GetJobs(r.Context(), r.URL.Query())

		if err != nil {
			h.customErr(w, r, 500, err)
			return
		}

		h.customRespond(w, r, http.StatusOK, u)

	}

}

func (h *Handler) customErr(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.customRespond(w, r, code, map[string]string{"error:": err.Error()})
}
func (h *Handler) customRespond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			glg.Debugf("", err)
		}
	}
}
