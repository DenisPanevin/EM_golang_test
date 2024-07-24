package dummyServer

import (
	"encoding/json"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"github.com/kpango/glg"
	"net/http"
	"time"
)

func StartDummy() {
	mux := mux2.NewRouter()

	// Register handler functions
	mux.HandleFunc("/info", greetHandler).Methods(http.MethodGet)

	httpServer := &http.Server{
		Addr:           "localhost:9090",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		glg.Fatal("Failed to listen and serve: %+v", err)
	}

}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Retrieve specific query parameters by their keys
	passportSerie := query.Get("passportSerie")
	passportNumber := query.Get("passportNumber")

	glg.Debugf("got from main app %s %s", passportSerie, passportNumber)

	d := map[string]respondUser{
		"1234 000001": {
			Surname:    "Иванов",
			Name:       "Иван",
			Patronymic: "Иванович",
			Address:    "г. Москва, ул. Ленина, д. 5, кв. 1",
		},
		"1234 000002": {
			Surname:    "Петров",
			Name:       "Петр",
			Patronymic: "Петрович",
			Address:    "г. Краснодар, ул. Сталина, д. 6, кв. 2",
		},
		"1234 000003": {
			Surname:    "Денисов",
			Name:       "Денис",
			Patronymic: "Денисович",
			Address:    "г. Екатеринбург, ул. Берии, д. 7, кв. 3",
		},
		"1234 000004": {
			Surname:    "Семенов",
			Name:       "Семен",
			Patronymic: "Семенович",
			Address:    "г. Тула, ул. Троцкого, д. 8, кв. 4",
		},
		"1234 000005": {
			Surname:    "Александров",
			Name:       "Александр",
			Patronymic: "Александрович",
			Address:    "г. Новосибирск, ул. Бухарина, д. 9, кв. 5",
		},
		"1234 100001": {
			Surname:    "Иванов",
			Name:       "Иван",
			Patronymic: "Александрович",
			Address:    "г. Новосибирск, ул. Бухарина, д. 9, кв. 5",
		},
	}

	if _, ok := d[fmt.Sprintf("%s %s", passportSerie, passportNumber)]; ok {
		glg.Debugf("returningOk")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(d[fmt.Sprintf("%s %s", passportSerie, passportNumber)])
		return
	}
	glg.Debugf("returning 500")
	w.WriteHeader(http.StatusBadRequest)
	return

	//w.WriteHeader(http.StatusInternalServerError)

}

type respondUser struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}
