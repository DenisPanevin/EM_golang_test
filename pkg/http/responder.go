package helpers

import (
	"encoding/json"
	"github.com/kpango/glg"
	"net/http"
)

type ErrorDto struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
type DataDto struct {
	Data interface{} `json:"data,omitempty"`
}

func SendError(w http.ResponseWriter, r *http.Request, code int, err error) {
	var e string
	if err != nil {
		e = err.Error()
	}
	SendJSON(w, r, code, ErrorDto{Code: code, Message: e})

}

func SendData(w http.ResponseWriter, r *http.Request, code int, data interface{}) {

	SendJSON(w, r, code, DataDto{Data: data})
}

func SendJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			glg.Debugf("cant encode json", err)
		}
	}
}
