package repository

import (
	"EM-Api-testTask/internal/models"
	"EM-Api-testTask/internal/users"
	"EM-Api-testTask/pkg"
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

func (u UsersRepo) CreateUser(ctx context.Context, dto models.CreateUserDto) (error, *models.User) {
	query := `INSERT INTO users ( 
                   passportnumber ,name ,surname ,patronymic, address
                   )VALUES ($1, $2,$3,$4,$5)RETURNING id,passportnumber,name,surname,patronymic,address`

	args := []interface{}{dto.PassportNumber, dto.Name, dto.Surname, dto.Patronymic, dto.Address}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var usr models.User
	err := u.db.QueryRow(ctx, query, args...).Scan(&usr.Id, &usr.PassportNumber, &usr.Name, &usr.Surname, &usr.Patronymic, &usr.Address)
	if err != nil {
		glg.Debugf("error add to psql %s", err)
		return err, nil
	}
	query = `INSERT INTO jobs ( user_id ,task_id,started,stopped ) VALUES ($1,$2,$3,$4)`
	endTime, _ := time.Parse("2006.01.02", "9999.12.31")
	u.db.QueryRow(ctx, query, usr.Id, 0, time.Time{}, endTime)
	return nil, &usr
}
func (u UsersRepo) UpdateUser(ctx context.Context, dto models.UpdateUserDto) (error, *models.User) {
	query := `UPDATE users
SET name = $1, surname = $2, patronymic = $3, passportnumber = $4, address = $5
WHERE id = $6
RETURNING id,passportnumber,name,surname,patronymic,address`

	args := []interface{}{dto.Name, dto.Surname, dto.Patronymic, dto.PassportNumber, dto.Address, dto.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var usr models.User
	err := u.db.QueryRow(ctx, query, args...).Scan(&usr.Id, &usr.PassportNumber, &usr.Name, &usr.Surname, &usr.Patronymic, &usr.Address)
	if err != nil {
		glg.Debugf("error update to psql %s", err)
		return err, nil
	}

	return nil, &usr

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
		return pkg.NotFound
	}

	return nil

}
func (u UsersRepo) GetAll(ctx context.Context, filters models.UserFiltersDto, pagefilters models.PageFiltersDto) (error, *[]models.ShowUserDto) {
	query := `select j.user_id ,
    u.name AS user_name,u.surname,u.patronymic,u.address,u.passportnumber,
    SUM(CASE WHEN j.task_id!=0  THEN j.stopped - j.started ELSE INTERVAL '0' END) AS total_work
FROM
    jobs j
        JOIN
    users u ON j.user_id = u.id
    where    
 (u.name=$1 or $1='') and (u.surname=$2 or $2='') and (u.patronymic=$3 or $3='')and (u.passportnumber=$4 or $4='')and (u.id=$5 or $5=0)
GROUP BY
    j.user_id, u.name,u.surname,u.patronymic,u.address,u.passportnumber
ORDER BY
    total_work DESC
    limit $6 offset $7`

	rows, err := u.db.Query(ctx, query, filters.Name, filters.Surname, filters.Patronymic, filters.PassportNumber, filters.Id, pagefilters.Limit, pagefilters.Offset)
	if err != nil {
		glg.Debugf("err from psql GET %s", err)
		return err, nil
	}
	defer rows.Close()
	var usersOutput []models.ShowUserDto
	for rows.Next() {
		usr := models.ShowUserDto{}

		err = rows.Scan(&usr.UserId, &usr.Name, &usr.Surname, &usr.Patronymic, &usr.Address, &usr.PassportNumber, &usr.TotalWorkTime)
		if err != nil {
			glg.Debugf("Row scan failed: %v\n", err)
			return err, nil
		}
		usersOutput = append(usersOutput, usr)
	}
	if len(usersOutput) < 1 {
		return pkg.NotFound, nil
	}

	return nil, &usersOutput
}
