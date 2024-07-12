package users

import (
	"EM-Api-testTask/internal/models"
	"context"
	"net/http"
)

type UseCase interface {
	Create(ctx context.Context, dto *models.PassportNumberDto) (error, *int64)
	Get(r *http.Request) (error, *[]models.ShowUserDto)

	Update(ctx context.Context, user models.UpdateUserDto) (error, *int64)

	DeleteUser(ctx context.Context, dto models.DeleteUserDto) error
}
