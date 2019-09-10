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
							User, 
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
			&user.Login,
			&user.Password,
		)

		if err != nil {
			fmt.Println(err)
		}

	}
	return user, err
}

func (r *SQLServerRepository) InsertUser(username, password string) (error) {
	rows, err := r.DB.Query(`INSERT INTO
							Users (User, Password)
							VALUES
							(` + username + "," + password + ")")

	if err != nil {
		fmt.Println(err)
		return err
	}

	rows.Close()
	return  nil
}


func (r *SQLServerRepository) ChangePassword(username, password string) (error) {
	rows, err := r.DB.Query(`UPDATE
							Users
							SET
							Password = ` + password + 
							` WHERE
							User = ` + username)

	if err != nil {
		fmt.Println(err)
		return err
	}
	rows.Close()
	return  nil
}