package delivery

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"EM-Api-testTask/pkg/handler"
	Helpers "EM-Api-testTask/pkg/helpers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kpango/glg"
	"golang.org/x/net/webdav"
	"net/http"
	"strconv"
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

		err, u := h.useCase.GetAll(r.Context(), r.URL.Query())

		if err != nil {
			h.customErr(w, r, 500, err)
			return
		}

		h.customRespond(w, r, http.StatusOK, u)

	}

}
func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			glg.Logf("cant parse id from updat user handler %s", err)
			h.customErr(w, r, http.StatusUnprocessableEntity, handler.ApiWrongInput)
			return
		}
		dto := models.UpdateUserDto{
			Id:             0,
			Name:           "",
			Surname:        "",
			Patronymic:     "",
			PassportNumber: "",
			Address:        "",
		}
		if err = json.NewDecoder(r.Body).Decode(&dto); err != nil {
			glg.Debugf("cant decode json in update handler %s", err)
			h.customErr(w, r, http.StatusBadRequest, handler.ApiWrongInput)
			return
		}
		v := Helpers.NewValidator()
		if err = v.Validate(dto); err != nil {
			glg.Debugf("cant validate json in update handler %s", err)
			h.customErr(w, r, http.StatusUnprocessableEntity, handler.ApiWrongInput)
			return
		}

		if dto.Name == "" && dto.Surname == "" && dto.Patronymic == "" && dto.PassportNumber == "" && dto.Address == "" {
			h.customErr(w, r, webdav.StatusUnprocessableEntity, handler.ApiWrongInput)
			return
		}
		dto.Id = id

		err, uid := h.useCase.Update(r.Context(), dto)
		if err != nil {
			glg.Debugf("error in update usecase returning %s", err)
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}
		h.customRespond(w, r, http.StatusOK, fmt.Sprintf("Updated, user id is %v", *uid))
	}

}
func (h *Handler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			glg.Logf("cant parse id from updat user handler %s", err)
			h.customErr(w, r, http.StatusUnprocessableEntity, handler.ApiWrongInput)
			return
		}
		err = h.useCase.DeleteUser(r.Context(), id)
		if err != nil {
			glg.Debugf("error in delete usecase returning %s", err)
			h.customErr(w, r, http.StatusInternalServerError, err)
			return
		}
		h.customRespond(w, r, http.StatusOK, fmt.Sprintf("Deleted user with id %v", id))
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
