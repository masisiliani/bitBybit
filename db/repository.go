package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"github.com/masisiliani/bitBybit/types"
)


type Repository interface {
	FindPostsByUser(username string) ([]types.Post, error)
	FindPostByID(ID int) (types.Post, error)
	InsertPost(p types.Post) (error)
	DeletePost(ID int) (error)
	UpdatePost(p types.Post) (error)
	InsertUser(u types.User) error
	FindUser(username string) (types.User, error)
	ChangePassword(username, password string) error
}

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

//Find a post by id
func (r *MySQLRepository) FindPostByID(ID int) (types.Post, error) {
	rows, err := r.DB.Query(`SELECT
								ID,
								Description,
								User,
								Date
							FROM Post
							WHERE ID = ` + strconv.Itoa(ID))

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var posts []types.Post
	var post types.Post

	for rows.Next() {

		err = rows.Scan(
			&post.ID,
			&post.Description,
			&post.UserName,
			&post.Date,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			posts = append(posts, post)
		}

	}

	return post, err
}

//FindByUser select all posts by username
func (r *MySQLRepository) FindPostsByUser(username string) ([]types.Post, error) {
	rows, err := r.DB.Query(`SELECT
								ID,
								Description,
								User,
								Date
							FROM Post
							WHERE User = ` + username)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var posts []types.Post
	var post types.Post

	for rows.Next() {

		err = rows.Scan(
			&post.ID,
			&post.Description,
			&post.UserName,
			&post.Date,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			posts = append(posts, post)
		}

	}

	return posts, err
}

//Insert insert a new post
func (r *MySQLRepository) InsertPost(post types.Post) error {
	result, err := r.DB.Exec(`INSERT INTO
							Posts (Description, User, Date)
							VALUES
							(` + post.Description + "," + post.UserName + "," + post.Date + ")")

	if err != nil {
		fmt.Println(err)
		return err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return err
	}

	post.ID = int(postID)

	return nil
}

//Delete a post
func (r *MySQLRepository) DeletePost(ID int) error {
	_, err := r.DB.Exec(`DELETE FROM
							Posts
							VALUES
							WHERE
							ID = ` + strconv.Itoa(ID))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//Update post
func (r *MySQLRepository) UpdatePost(post types.Post) error {
	_, err := r.DB.Exec(`UPDATE
							Posts
							SET
							Description = ` + post.Description +
		` WHERE
							ID = ` + strconv.Itoa(post.ID))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//Find a user by username
func (r *MySQLRepository) FindUser(username string) (types.User, error) {
	rows, err := r.DB.Query(`SELECT
								User,
								Password
							FROM Users
							WHERE User = ` + username)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var user types.User

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
func (r *MySQLRepository) InsertUser(u types.User) error {
	_, err := r.DB.Exec(`INSERT INTO
							Users (User, Password)
							VALUES
							(` + u.UserName + "," + u.Password + ")")

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

