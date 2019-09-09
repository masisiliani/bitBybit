package user

import (
	"database/sql"
	"fmt"
	"strconv"
)

//SQLServerRepository mongodb repo
type SQLServerRepository struct {
	DB *sql.DB
}

//NewSqlServerRepository create new repository
func NewSqlServerRepository(db *sql.DB) *SQLServerRepository {
	return &SQLServerRepository{
		DB: db,
	}
}

func (r *SQLServerRepository) Find(ID int32) (User, error) {
	rows, err := r.DB.Query(`SELECT 
							ID, 
							Login, 
							Password
						FROM Users 
						WHERE ID = ` + strconv.Itoa(int(ID)))

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var user User

	for rows.Next() {

		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.Password,
		)

		if err != nil {
			fmt.Println(err)
		}

	}

	return user, err

}
