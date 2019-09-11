package post

import (
	db "github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/types"

	"time"
	"fmt"
)

type PostController struct {
	Repository db.Repository
}

func (pc *PostController) GetPostByID(id int) (types.Post, error){
	post, err := pc.Repository.FindPostByID(id)
	return post, err
}

func  (pc *PostController)  GetPosts(username string) ([]types.Post, error){
	posts, err := pc.Repository.FindPostsByUser(username)
	return posts, err
}

func  (pc *PostController) InsertPost(p types.Post) (error){
	date := time.Now().Format("2006-01-02 15:04:05")
	p.Date = date
	fmt.Println(p.Date)
	return pc.Repository.InsertPost(p)
}

func  (pc *PostController) DeletePost(id int) error{
	return pc.Repository.DeletePost(id)
}

func  (pc *PostController) UpdatePost(p types.Post) error{
	return pc.Repository.UpdatePost(p)
}
