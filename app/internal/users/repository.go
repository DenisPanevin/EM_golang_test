package users

import (
	"EM-Api-testTask/internal/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, dto models.CreateUserDto) (error, *models.User)
	UpdateUser(ctx context.Context, filters models.UpdateUserDto) (error, *models.User)
	DeleteUser(ctx context.Context, id int) error
	GetAll(ctx context.Context, filters models.UserFiltersDto, pagefilters models.PageFiltersDto) (error, *[]models.ShowUserDto)
	/*GetByEmail(ctx context.Context, email string) (error, *models.User)
	GetList(ctx context.Context, page int, size int) (error, []models.User)
	Update(ctx context.Context, user CreateUserDto) (error, *user.User)
	AddSubscription(ctx context.Context, subscribeeId int) (error, int)
	GetAllSubscriptions(ctx context.Context, subscriberId int) (error, []models.User)
	DeleteUser(ctx context.Context, dto DeleteUserDto) error*/
}
