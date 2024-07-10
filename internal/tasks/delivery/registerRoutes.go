package delivery

import (
	"EM-Api-testTask/internal/tasks"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(router *mux.Router, useCase tasks.UseCase) {
	h := NewHandler(useCase)
	router.Handle("/tasks", h.CreateTask()).Methods("POST")
	//	router.Handle("/users", h.CreateUser()).Methods("POST")

	//router.Handle("/users", h.Get()).Methods("GET")

}
