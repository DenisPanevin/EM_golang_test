package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	Apierrors "EM-Api-testTask/pkg"
	Helpers "EM-Api-testTask/pkg/helpers"
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
func (uc *UserUseCase) Create(ctx context.Context, dto *models.PassportNumberDto) (error, *models.User) {
	SideServerUrl := viper.GetString("app.SideServerUrl")
	splitCredentials := strings.Split(dto.PassportNumber, " ")
	url := fmt.Sprintf("http://%s/info?passportSerie=%s&passportNumber=%s", SideServerUrl, splitCredentials[0], splitCredentials[1])

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glg.Debugf("error forming request to remote %s", err)
		return err, nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glg.Debugf("cant send to remote server %s", err)
		return err, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		glg.Debugf("cant get from remote server %v", resp.StatusCode)
		err = Apierrors.NotFound
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

func (uc *UserUseCase) Update(ctx context.Context, dto models.UpdateUserDto) (error, *models.User) {
	vals := url.Values{}
	vals.Set("id", strconv.Itoa(dto.Id))
	err, usr := uc.GetAll(ctx, vals)
	if err != nil {
		glg.Debugf("Error in update from get All usecase %s", err)
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
	e, id := uc.repo.UpdateUser(ctx, dto)
	return e, id

}
func (uc *UserUseCase) GetAll(ctx context.Context, vals url.Values) (error, *[]models.ShowUserDto) {
	filters := models.UserFiltersDto{}

	if vals.Get("id") != "" {
		id, err := strconv.Atoi(vals.Get("id"))
		if err != nil || id <= 0 {
			glg.Debugf("error parsing id in filters %s", err)
			return Apierrors.ApiWrongInput, nil
		} else {
			filters.Id = id
		}
	}

	filters.Name = vals.Get("name")
	filters.Surname = vals.Get("surname")
	filters.Patronymic = vals.Get("patronymic")
	filters.PassportNumber = vals.Get("passportnumber")
	filters.Address = vals.Get("address")
	println(vals.Get("passportnumber"))
	v := Helpers.NewValidator()
	err := v.Validate(filters)
	if err != nil {
		glg.Debugf("error validating filters %s", err)
		return Apierrors.ApiWrongInput, nil
	}

	pageFilters := models.PageFiltersDto{
		Limit: 50,
		Page:  1,
	}
	if vals.Get("limit") != "" {
		limit, e := strconv.Atoi(vals.Get("limit"))
		if e != nil || limit <= 0 {
			glg.Debugf("error parsing page in Pagefilters %s", e)
			return Apierrors.ApiWrongInput, nil
		} else {
			pageFilters.Limit = limit
		}
	}
	if vals.Get("page") != "" {
		page, er := strconv.Atoi(vals.Get("page"))
		if er != nil || page <= 0 {
			glg.Debugf("error parsing page in Pagefilters %s", er)
			return Apierrors.ApiWrongInput, nil
		} else {
			pageFilters.Page = page
		}
	}

	pageFilters.CalculateOffset()

	err, u := uc.repo.GetAll(ctx, filters, pageFilters)
	if err != nil {
		glg.Debugf("cant get all from users repo in usecase %s", err)
		return err, nil
	}
	return nil, u
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int) error {
	err := uc.repo.DeleteUser(ctx, id)
	return err
}
