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

func (u UsersRepo) GetJob(ctx context.Context, filters models.FiltersDto) (error, *[]models.ShowUserDto) {

	//query := `SELECT id,passportnumber,name,surname,patronymic,address FROM users WHERE (name=$1 or $1='') and (surname=$2 or $2='') and (patronymic=$3 or $3='') and ($4=0) limit $5 offset $6`
	//	query := `SELECT id,passportnumber,name,surname,patronymic,address FROM users WHERE (name=$1 or $1='')`

	query := `SELECT
    j.task_id,
    j.user_id,
    u.name AS user_name,u.surname,u.patronymic,u.address,u.passportnumber,
    t.name as task_name,
    SUM(CASE WHEN j.stopped - j.started >= INTERVAL '0' THEN j.stopped - j.started ELSE INTERVAL '0' END) AS total_work
FROM
    jobs j
        JOIN
    users u ON j.user_id = u.id
        JOIN
    tasks t ON j.task_id = t.id
WHERE
   (u.id=$1 or $1=0)and (u.name=$2 or $2='') and (u.surname=$3 or $3='') and (u.patronymic=$4 or $4='') and ($5=0) 
GROUP BY
    j.task_id, j.user_id, u.name,u.surname,u.patronymic,t.name,u.address,u.passportnumber
ORDER BY
    total_work DESC
limit $6 offset $7;`

	values, err := filtersToValues(filters)

	if err != nil {
		return err, nil
	}

	rows, err := u.db.Query(ctx, query, values...)
	if err != nil {
		glg.Debugf("err from psql GET", err)
		return err, nil
	}
	defer rows.Close()

	usrs := []models.ShowUserDto{}

	for rows.Next() {
		usr := models.ShowUserDto{}
		//err = rows.Scan(&usr.TaskId, &usr.UserId, &usr.Name, &usr.Surname, &usr.Patronymic, &usr.Address, &usr.PassportNumber, &usr.TaskName, &usr.TotalWork.TotalWork)
		if err != nil {
			glg.Debugf("Row scan failed: %v\n", err)
			return err, nil
		}

		usrs = append(usrs, usr)

	}

	return nil, &usrs

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
	return nil, &id
}
