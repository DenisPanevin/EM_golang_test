package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"EM-Api-testTask/pkg/handler"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"strings"
)

type UserUseCase struct {
	repo users.Repository
}

func NewUserUseCase(ur users.Repository) users.UseCase {
	return &UserUseCase{
		repo: ur,
	}
}
func (uc *UserUseCase) Create(ctx context.Context, dto *models.PassportNumberDto) (error, *int64) {
	SideServerUrl := viper.GetString("app.SideServerUrl")
	splitCredentials := strings.Split(dto.PassportNumber, " ")
	url := fmt.Sprintf("http://%s/info?passportSerie=%s&passportNumber=%s", SideServerUrl, splitCredentials[0], splitCredentials[1])

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glg.Debugf("error forming request to remote %s", err)
		return err, nil
	}

	//req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP POST request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glg.Debugf("cant send to remote server %s", err)
		return err, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		glg.Debugf("cant get from remote server %v", resp.StatusCode)
		return err, nil
	}

	cu := models.CreateUserDto{
		PassportNumber: dto.PassportNumber,
	}
	if err = json.NewDecoder(resp.Body).Decode(&cu); err != nil {
		glg.Debugf("error decoding response from remote %s ", err)
		return err, nil
	}
	//add validation here
	err, id := uc.repo.CreateUser(ctx, cu)
	if err != nil {
		glg.Debugf("error from repo layer in creating user %s ", err)
		return err, nil
	}
	return nil, id
}
func (uc *UserUseCase) GetJob(r *http.Request) (error, *[]models.ShowUserDto) {
	query := r.URL.Query()
	dto := models.FiltersDto{
		Id:         0,
		Name:       "",
		Surname:    "",
		Patronymic: "",
		TaskId:     0,
		Limit:      10,
		Page:       1,
	}

	for key, values := range query {
		println(values[0])
		switch key {
		case "id":
			id, err := strconv.Atoi(values[0])
			if err != nil {
				glg.Debugf("cant Parse int id from url", err)
				return handler.ApiWrongInput, nil
				dto.Page = 1
			}
			if id > 0 {
				dto.Id = id
			} else {
				return handler.ApiWrongInput, nil
			}
		case "name":
			dto.Name = values[0]

		case "surname":
			dto.Surname = values[0]

		case "patronymic":
			dto.Patronymic = values[0]

		case "taskid":

			id, err := strconv.Atoi(values[0])
			if err != nil {
				glg.Debugf("cant Parse int from url", err)
				return handler.ApiWrongInput, nil
				dto.TaskId = 0
			}
			if id > 0 {
				dto.TaskId = id
			} else {
				return handler.ApiWrongInput, nil
			}

		case "limit":

			lim, err := strconv.Atoi(values[0])
			if err != nil {
				glg.Debugf("cant Parse int from url", err)
				return handler.ApiWrongInput, nil
				dto.Limit = 0
			}
			if lim > 0 {
				dto.Limit = lim
			} else {
				return handler.ApiWrongInput, nil
			}
		case "page":

			page, err := strconv.Atoi(values[0])
			if err != nil {
				glg.Debugf("cant Parse int from url", err)
				return handler.ApiWrongInput, nil
				dto.Page = 1
			}
			if page > 0 {
				dto.Page = page
			} else {
				return handler.ApiWrongInput, nil
			}
		}

	}

	err, u := uc.repo.GetJob(r.Context(), dto)
	if err != nil {
		glg.Debugf("usecase error ,%s", err)
		return err, nil
	}

	return nil, u

}

func (uc *UserUseCase) Update(ctx context.Context, user models.UpdateUserDto) (error, *int64) {
	return nil, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, dto models.DeleteUserDto) error {
	return nil
}
