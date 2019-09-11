package main

import (
	db "github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/api/router"
	"github.com/masisiliani/bitBybit/pkg/user"
	"github.com/masisiliani/bitBybit/pkg/post"
	"os"
	"net/http"
	"fmt"
)



func main(){
	if err := db.InitDatabase(); err != nil{
		fmt.Println(err)
	}

	database, err := db.Connect_SQL()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	r := router.Router{
		UserController: user.UserController{
			Repository: &db.MySQLRepository{
				DB: database,
			},
		},
		PostController: post.PostController{
			Repository: &db.MySQLRepository{
				DB: database,
			},
		},
	}
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Working!")
	})
	http.HandleFunc("/login", r.Login)
	http.HandleFunc("/register", r.Register)
	http.HandleFunc("/changePassword", r.ChangePassword)
	http.HandleFunc("/post/new", r.NewPost)
	http.HandleFunc("/post", r.GetPost)
	http.HandleFunc("/post/all", r.GetPosts)
	http.HandleFunc("/post/delete", r.DeletePost)

	fmt.Println("Starting server...")
	if err := http.ListenAndServe("0.0.0.0:3001", nil); err != nil {
		fmt.Println(err)
        os.Exit(1)
    }
}