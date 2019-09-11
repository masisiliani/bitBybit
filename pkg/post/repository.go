package post

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Repository interface {
	FindPosts(username string) ([]Post, error)
	FindPost(ID int) (Post, error)
	InsertPost(username string, description string) (error)
	DeletePost(ID int) (error)
	UpdatePost(ID int, description string) (error)
}

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

func (r *SQLServerRepository) FindPosts(username string) ([]Post, error) {
	rows, err := r.DB.Query(`SELECT
							ID, 
							Description, 
							User
						FROM Post 
						WHERE User = ` + username)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var posts []Post
	var post Post

	for rows.Next() {

		err = rows.Scan(
			&post.ID,
			&post.Description,
			&post.User,
		)

		if err != nil {
			fmt.Println(err)
		}else{
			posts = append(posts, post)
		}

	}

	return posts, err
}


func (r *SQLServerRepository) FindPost(ID int) (Post, error) {
	rows, err := r.DB.Query(`SELECT
							ID, 
							Description, 
							User
						FROM Post 
						WHERE ID = ` + strconv.Itoa(int(ID)))

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var post Post

	for rows.Next() {

		err = rows.Scan(
			&post.ID,
			&post.Description,
			&post.User,
		)

		if err != nil {
			fmt.Println(err)
		}

	}

	return post, err
}


func (r *SQLServerRepository) InsertPost(username string, description string) (error) {
	rows, err := r.DB.Query(`INSERT INTO
							Posts (Description, User)
							VALUES
							(` + description + "," + username + ")")

	if err != nil {
		fmt.Println(err)
		return err
	}
	rows.Close()
	return  nil
}


func (r *SQLServerRepository) DeletePost(ID int) (error) {
	rows, err := r.DB.Query(`DELETE FROM
							Posts
							VALUES
							WHERE
							ID = ` + strconv.Itoa(ID))

	if err != nil {
		fmt.Println(err)
		return err
	}
	rows.Close()
	return  nil
}

func (r *SQLServerRepository) UpdatePost(ID int, description string) (error) {
	rows, err := r.DB.Query(`UPDATE
							Posts
							SET
							Description = ` + description + 
							` WHERE
							ID = ` + strconv.Itoa(ID))

	if err != nil {
		fmt.Println(err)
		return err
	}
	rows.Close()
	return  nil
}