package delivery

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"EM-Api-testTask/pkg/handler"
	Helpers "EM-Api-testTask/pkg/helpers"
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"net/http"
)

type Handler struct {
	useCase users.UseCase
}

func NewHandler(uc users.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(models.PassportNumberDto)
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			h.customErr(w, r, http.StatusBadRequest, err)
			return
		}
		v := Helpers.NewValidator()
		if err := v.Validate(dto); err != nil {
			h.customErr(w, r, http.StatusUnprocessableEntity, handler.ApiWrongInput)
			return
		}

		err, id := h.useCase.Create(r.Context(), dto)
		if err != nil {
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}
		if id != nil {
			h.customRespond(w, r, http.StatusOK, fmt.Sprintf("Created! user id is %v", *id))
			return
		}

		h.customRespond(w, r, http.StatusInternalServerError, nil)
		return

	}

}

func (h *Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err, u := h.useCase.GetJob(r)

		if err != nil {
			h.customErr(w, r, 500, err)
		}

		h.customRespond(w, r, http.StatusOK, u)

	}

}
func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		/*err, u := h.useCase.Update(r)

		if err != nil {
			h.customErr(w, r, 500, err)
		}

		h.customRespond(w, r, http.StatusOK, u)
		*/
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

// misc
func (h *Handler) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Check")

	}

}
