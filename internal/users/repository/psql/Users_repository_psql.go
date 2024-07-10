package repository

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"context"
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

func (u UsersRepo) Get(ctx context.Context, filters models.FiltersDto) (error, *models.User) {

	query := `SELECT id,passportnumber,name,surname,patronymic,address FROM users WHERE (name=$1 or $1='') and (surname=$2 or $2='') and (patronymic=$3 or $3='') and ($4=0)`
	//query := `SELECT id,passportnumber,name,surname,patronymic,address FROM users WHERE (name=$1 or $1='')`

	values := filtersToValues(filters)

	rows, err := u.db.Query(ctx, query, values...)
	if err != nil {
		glg.Debugf("err from psql GET", err)
		return err, nil
	}
	defer rows.Close()

	usrs := make([]models.User, 0)

	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.Id, &u.PassportNumber, &u.Name, &u.Surname, &u.Patronymic, &u.Adress)
		if err != nil {
			glg.Debugf("Row scan failed: %v\n", err)
		}
		usrs = append(usrs, u)

	}

	for _, v := range usrs {
		println(v.Name)
	}

	return nil, nil

}

func (u UsersRepo) CreateUser(ctx context.Context, dto models.CreateUserDto) (error, *int64) {
	query := `INSERT INTO users ( passportnumber ,name ,surname ,patronymic, address)VALUES ($1, $2,$3,$4,$5)RETURNING id`

	args := []interface{}{dto.PassportNumber, dto.Name, dto.Surname, dto.Patronymic, dto.Address}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := u.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		glg.Debugf("error add to psql", err)
		return err, nil
	}
	return nil, &id
}
