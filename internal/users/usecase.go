package users

import (
	"EM-Api-testTask/internal/models"
	"context"
	"net/url"
)

type UseCase interface {
	Create(ctx context.Context, dto *models.PassportNumberDto) (error, *int64)
	GetAll(ctx context.Context, vals url.Values) (error, *[]models.ShowUserDto)
	Update(ctx context.Context, user models.UpdateUserDto) (error, *int64)

	DeleteUser(ctx context.Context, id int) error
}
