package post

import (
	"database/sql"
	"fmt"
	"strconv"
)

<<<<<<< HEAD
type Repository interface {
	FindPosts(username string) ([]Post, error)
	FindPost(ID int) (Post, error)
	InsertPost(username string, description string) (error)
	DeletePost(ID int) (error)
	UpdatePost(ID int, description string) (error)
}

//SQLServerRepository mongodb repo
type SQLServerRepository struct {
=======
//MySQLRepository MySql repo
type MySQLRepository struct {
>>>>>>> c0907c2572c30583feae235947dc883146b53c7f
	DB *sql.DB
}

//NewMySQLRepository create new repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		DB: db,
	}
}

//Find a post by id
func (r *MySQLRepository) Find(ID int) ([]Post, error) {
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

	var posts []Post
	var post Post

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

//FindByUser select all posts by username
func (r *MySQLRepository) FindByUser(username string) ([]Post, error) {
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

	var posts []Post
	var post Post

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
func (r *MySQLRepository) Insert(post *Post) error {
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
func (r *MySQLRepository) Delete(ID int) error {
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
func (r *MySQLRepository) Update(post *Post) error {
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
