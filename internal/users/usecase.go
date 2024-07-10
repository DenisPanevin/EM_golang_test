package users

import (
	"EM-Api-testTask/internal/models"
	"context"
	"net/http"
	"os/user"
)

type UseCase interface {
	Create(ctx context.Context, dto models.CreateUserDto) (error, *int64)
	Get(r *http.Request) (error, *models.User)

	Edit(ctx context.Context, user models.CreateUserDto) (error, *user.User)

	DeleteUser(ctx context.Context, dto models.DeleteUserDto) error
}
