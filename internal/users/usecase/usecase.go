package usecase

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"context"
	"github.com/kpango/glg"
	"net/http"
	"os/user"
	"strconv"
)

type UserUseCase struct {
	repo users.Repository
}

func NewUserUseCase(ur users.Repository) users.UseCase {
	return &UserUseCase{
		repo: ur,
	}
}
func (uc *UserUseCase) Create(ctx context.Context, dto models.CreateUserDto) (error, *int64) {
	//check dto for empty settings
	err, id := uc.repo.CreateUser(ctx, dto)
	if err != nil {
		return err, nil
	}
	return nil, id
}
func (uc *UserUseCase) Get(r *http.Request) (error, *models.User) {
	query := r.URL.Query()
	dto := models.FiltersDto{
		Name:       "",
		Surname:    "",
		Patronymic: "",
		TaskId:     0,
	}

	for key, values := range query {
		println(key)
		switch key {
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
			}
			if id > 0 {
				dto.TaskId = id
			}
		}

	}
	////////////////////
	err, u := uc.repo.Get(r.Context(), dto)
	if err != nil {
		return err, nil
	}
	return nil, u

}

func (uc *UserUseCase) Edit(ctx context.Context, user models.CreateUserDto) (error, *user.User) {
	return nil, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, dto models.DeleteUserDto) error {
	return nil
}
