package delivery

import (
	"EM-Api-testTask/internal/jobs"
	"github.com/gorilla/mux"
)

func RegisterJobsRoutes(router *mux.Router, useCase jobs.UseCase) {
	h := NewHandler(useCase)
	router.Handle("/jobs/user", h.GetJobs()).Methods("GET")
	router.Handle("/jobs", h.AddJob()).Methods("POST")

}
