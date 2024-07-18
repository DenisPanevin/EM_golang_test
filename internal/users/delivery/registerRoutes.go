package delivery

import (
	"EM-Api-testTask/internal/users"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserRoutes(router *mux.Router, useCase users.UseCase) {
	h := NewHandler(useCase)

	router.Handle("/users", h.CreateUser()).Methods(http.MethodPost)
	router.Handle("/users/{id}", h.UpdateUser()).Methods(http.MethodPatch)
	router.Handle("/users/{id}", h.DeleteUser()).Methods(http.MethodDelete)

	router.Handle("/users", h.Get()).Methods(http.MethodGet)
	router.Handle("/check", h.CheckHealth()).Methods(http.MethodGet)

}
