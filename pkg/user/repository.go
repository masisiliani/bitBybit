package user

import (
	"database/sql"
	"fmt"
)

//MySQLRepository MySql repo
type MySQLRepository struct {
	DB *sql.DB
}

//NewMySQLRepository create new repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		DB: db,
	}
}

//Find a user by username
func (r *MySQLRepository) Find(username string) (User, error) {
	rows, err := r.DB.Query(`SELECT
								User,
								Password
							FROM Users
							WHERE User = ` + username)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var user User

	for rows.Next() {

		err = rows.Scan(
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			fmt.Println(err)
		}

	}
	return user, err
}

//Insert a new User on Database
func (r *MySQLRepository) Insert(username, password string) error {
	_, err := r.DB.Exec(`INSERT INTO
							Users (User, Password)
							VALUES
							(` + username + "," + password + ")")

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//ChangePassword udpdate the user's password
func (r *MySQLRepository) ChangePassword(username, password string) error {
	_, err := r.DB.Exec(`UPDATE
							Users
							SET
							Password = ` + password +
		` WHERE
							User = ` + username)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
