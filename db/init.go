package db

import (
	"fmt"
	"database/sql"
)

func InitDatabase() error {

	db, err := Connect_SQL()
	if err != nil {
		return err
	}

	defer db.Close()
	if err := createPostsTable(db); err != nil {
		fmt.Println(err)
	}
	if err := createUsersTable(db); err != nil {
		fmt.Println(err)
	}
	return nil
}

func createUsersTable(db *sql.DB) error {
	query := fmt.Sprint(`CREATE TABLE Users
						(User varchar(26),
						Password varchar(80),
						PRIMARY KEY (User))`)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createPostsTable(db *sql.DB) error {
	query := fmt.Sprint(`CREATE TABLE Posts
						(ID int NOT NULL AUTO_INCREMENT,
						Description varchar(200),
						User varchar(26),
						Date varchar(50),
						PRIMARY KEY (ID))`)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
