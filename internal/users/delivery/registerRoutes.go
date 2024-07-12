package delivery

import (
	"EM-Api-testTask/internal/users"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserRoutes(router *mux.Router, useCase users.UseCase) {
	h := NewHandler(useCase)
	router.Handle("/check", h.CheckHealth()).Methods("GET")
	router.Handle("/users", h.CreateUser()).Methods("POST")
	router.Handle("/users", h.UpdateUser()).Methods(http.MethodPatch)

	router.Handle("/users", h.Get()).Methods("GET")

}
