package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
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

func (uc *UserUseCase) Update(ctx context.Context, dto models.UpdateUserDto) (error, *int64) {
	vals := url.Values{}
	vals.Set("id", strconv.Itoa(dto.Id))
	err, usr := uc.GetAll(ctx, vals)
	if err != nil {
		glg.Debugf("Error in update from get job usecase %s", err)
		return err, nil
	}
	if dto.Name == "" {
		dto.Name = (*usr)[0].Name
	}
	if dto.Surname == "" {
		dto.Surname = (*usr)[0].Surname
	}
	if dto.Patronymic == "" {
		dto.Patronymic = (*usr)[0].Patronymic
	}
	if dto.PassportNumber == "" {
		dto.PassportNumber = (*usr)[0].PassportNumber
	}
	if dto.Address == "" {
		dto.Address = (*usr)[0].Address
	}
	err, id := uc.repo.UpdateUser(ctx, dto)
	return nil, id
}
func (uc *UserUseCase) GetAll(ctx context.Context, vals url.Values) (error, *[]models.ShowUserDto) {
	return nil, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int) error {
	err := uc.repo.DeleteUser(ctx, id)
	return err
}
