package delivery

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
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
	useCase users.UseCase
}

func NewHandler(uc users.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

// @Summary      Creates new user
// @Description  Creates new user by passport number
// @Accept       json
// @Produce      json
// @Param request body models.PassportNumberDto	true	"Passport number in xxxx xxxxxx format"
//
//	@Success 201  {object} models.User "returning User object"
//	@Failure 400	{object}	helpers.ErrorDto
//	@Failure 422	{object}	helpers.ErrorDto
//	@Failure 500	{object}	helpers.ErrorDto
//
// @Router       /users [post]
func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := new(models.PassportNumberDto)
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			helpers.SendError(w, r, http.StatusBadRequest, err)
			return
		}
		v := Helpers.NewValidator()
		if err := v.Validate(dto); err != nil {
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)
			return
		}

		err, usr := h.useCase.Create(r.Context(), dto)

		if err != nil {

			helpers.SendError(w, r, http.StatusInternalServerError, err)
			return
		}

		helpers.SendData(w, r, http.StatusCreated, usr)
		return

	}

}

// @Summary      Get list of users
// @Description  return list of users
// @Produce      json
// @Param id query int false "User ID"
// @Param name query string false "User name"
// @Param surname query int false "User surname"
// @Param patronymic query string false "User patronymic"
// @Param address query string false "User address"
// @Param passportnumber query string false "User passpor tnumber"
// @Param page query int false "page"
// @Param limit query int false "limit"
//
//	@Success 200  {object} []models.ShowUserDto{Jobs=nil} "returning Users objects"
//	@Failure 500	{object}	helpers.ErrorDto
//
// @Router       /users [get]
func (h *Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err, u := h.useCase.GetAll(r.Context(), r.URL.Query())

		if err != nil {
			helpers.SendError(w, r, 500, err)
			return
		}

		helpers.SendData(w, r, 200, u)

	}

}

// @Summary      Updates a user
// @Description  Updates a user by id
// @Accept       json
// @Produce      json
// @Param id path int true "User ID"
// @Param request body models.UpdateUserDto	true	"fields need to be updated"
//
//	@Success 200  {object} models.User "returning User object"
//	@Failure 500	{object}	helpers.ErrorDto
//
// @Router       /users/{id} [patch]
func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			glg.Logf("cant parse id from updat user handler %s", err)
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)
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
			helpers.SendError(w, r, http.StatusBadRequest, Apierrors.ApiWrongInput)
			return
		}
		v := Helpers.NewValidator()
		if err = v.Validate(dto); err != nil {
			glg.Debugf("cant validate json in update handler %s", err)
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)
			return
		}

		if dto.Name == "" && dto.Surname == "" && dto.Patronymic == "" && dto.PassportNumber == "" && dto.Address == "" {
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)
			return
		}
		dto.Id = id

		err, usr := h.useCase.Update(r.Context(), dto)
		if err != nil {
			glg.Debugf("error in update usecase returning %s", err)
			helpers.SendError(w, r, http.StatusInternalServerError, err)
			return
		}
		helpers.SendData(w, r, http.StatusOK, *usr)
	}

}

// @Summary      Delete a user
// @Description  Deletes a user by id
// @Param id path int true "User ID"
//
//	@Success 200  {string }	string	 "{Deleted id: id}"
//	@Failure 500	{object}	helpers.ErrorDto
//	@Failure 422	{object}	helpers.ErrorDto
//
// @Router       /users/{id} [delete]
func (h *Handler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			glg.Logf("cant parse id from updat user handler %s", err)
			helpers.SendError(w, r, http.StatusUnprocessableEntity, Apierrors.ApiWrongInput)

			return
		}
		err = h.useCase.DeleteUser(r.Context(), id)
		if err != nil {
			glg.Debugf("error in delete usecase returning %s", err)
			helpers.SendError(w, r, http.StatusInternalServerError, err)
			return
		}
		helpers.SendData(w, r, http.StatusOK, fmt.Sprintf(fmt.Sprintf(`{Deleted id: %v}`, id)))
	}

}
func (h *Handler) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		helpers.SendData(w, r, 200, fmt.Sprint("all good"))

	}

}
