package repository

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kpango/glg"
	"time"
)

func NewUsersRepo(pool *pgxpool.Pool) users.Repository {
	return UsersRepo{
		db: pool,
	}
}

type UsersRepo struct {
	db *pgxpool.Pool
}

func (u UsersRepo) CreateUser(ctx context.Context, dto models.CreateUserDto) (error, *int64) {
	query := `INSERT INTO users ( passportnumber ,name ,surname ,patronymic, address)VALUES ($1, $2,$3,$4,$5)RETURNING id`

	args := []interface{}{dto.PassportNumber, dto.Name, dto.Surname, dto.Patronymic, dto.Address}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := u.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		glg.Debugf("error add to psql %s", err)
		return err, nil
	}
	query = `INSERT INTO jobs ( user_id ,task_id,started,stopped ) VALUES ($1,$2,$3,$4)`
	endTime, _ := time.Parse("2006.01.02", "9999.12.31")
	u.db.QueryRow(ctx, query, id, 0, time.Time{}, endTime)
	return nil, &id
}
func (u UsersRepo) UpdateUser(ctx context.Context, dto models.UpdateUserDto) (error, *int64) {
	query := `UPDATE users
SET name = $1, surname = $2, patronymic = $3, passportnumber = $4, address = $5
WHERE id = $6
RETURNING id`

	args := []interface{}{dto.Name, dto.Surname, dto.Patronymic, dto.PassportNumber, dto.Address, dto.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := u.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		glg.Debugf("error update to psql %s", err)
		return err, nil
	}

	return nil, &id

}
func (u UsersRepo) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users
WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tag, err := u.db.Exec(ctx, query, id)
	if err != nil {
		glg.Debugf("error delete to psql %s", err)
		return err
	}
	if tag.RowsAffected() != 1 {
		println(tag.RowsAffected())
		err = errors.New("no such user already")
		return err
	}

	return nil

}
