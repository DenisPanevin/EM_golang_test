package delivery

import (
	"EM-Api-testTask/internal/tasks"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(router *mux.Router, useCase tasks.UseCase) {
	h := NewHandler(useCase)
	router.Handle("/tasks", h.CreateTask()).Methods("POST")
	router.Handle("/tasks/{id}", h.DeleteTask()).Methods("DELETE")
	router.Handle("/tasks", h.GetAll()).Methods("GET")

}
