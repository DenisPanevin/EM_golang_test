package delivery

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"bytes"
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

		/*if err := validators.Validate(*dto); err != nil {
			h.customErr(w, r, http.StatusUnprocessableEntity, internal.ApiWrongInput)
			return
		}*/

		err, cu := SendCredentials("http://localhost:9090/", *dto)
		if err != nil {
			glg.Debugf("send credentials", err)
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}

		/*if err = validators.Validate(*cu); err != nil {
			glg.Debugf("validation error", err)
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}*/

		err, id := h.useCase.Create(r.Context(), *cu)
		if err != nil {
			h.customRespond(w, r, http.StatusInternalServerError, nil)
			return
		}

		h.customRespond(w, r, http.StatusOK, fmt.Sprintf("Created! user id is %v", *id))

	}

}

func (h *Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		h.useCase.Get(r)

	}

}

func SendCredentials(url string, credentials models.PassportNumberDto) (error, *models.CreateUserDto) {
	//url := "http://localhost:9090/"

	jsonData, err := json.Marshal(credentials)
	if err != nil {
		return err, nil
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err, nil
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP POST request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err, nil
	}

	cu := models.CreateUserDto{
		PassportNumber: credentials.PassportNumber,
	}
	if err := json.NewDecoder(resp.Body).Decode(&cu); err != nil {
		return err, nil
	}

	return nil, &cu

}
func (h *Handler) customErr(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.customRespond(w, r, code, map[string]string{"error:": err.Error()})
}
func (h *Handler) customRespond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// misc
func (h *Handler) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Check")

	}

}
