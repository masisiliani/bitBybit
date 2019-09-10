package database

import (
	"fmt"
)

func InitDatabase() error {

	err := Connect_SQL()
	if err != nil {
		return err
	}

	defer DB.Close()
	if err := createPostsTable(); err != nil {
		return err
	}
	if err := createUsersTable(); err != nil {
		return err
	}
	return nil
}

func createUsersTable() error {
	query := fmt.Sprint(`CREATE TABLE Users
						(Username varchar(26),
						Password varchar(80),
						PRIMARY KEY (Username))`)
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createPostsTable() error {
	query := fmt.Sprint(`CREATE TABLE Posts
						(ID int NOT NULL AUTO_INCREMENT,
						Description varchar(200),
						Username varchar(26),
						Date varchar(20),
						PRIMARY KEY (ID))`)
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
