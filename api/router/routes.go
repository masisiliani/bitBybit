package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/pkg/post"
	"github.com/masisiliani/bitBybit/pkg/user"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var controllers Controller

func init() {
	if err := db.InitDatabase(); err != nil {
		fmt.Println(err)
	}

	database, err := db.Connect_SQL()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	controllers = Controller{
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

}

var routes = Routes{
	Route{
		"Healthcheck",
		"GET",
		"/healthcheck",
		Index,
	},
	Route{
		"Login",
		"POST",
		"/login",
		controllers.Login,
	},
	Route{
		"Create an user",
		"POST",
		"/user",
		controllers.Register,
	},
	Route{
		"Change user password",
		"PUT",
		"/user",
		controllers.ChangePassword,
	},
	Route{
		"Insert new post",
		"POST",
		"/post",
		controllers.NewPost,
	},
	Route{
		"Delete post",
		"DELETE",
		"/post",
		controllers.DeletePost,
	},
	Route{
		"Get all posts",
		"GET",
		"/post",
		controllers.GetPosts,
	},
	Route{
		"Get all posts",
		"GET",
		"/post/{postid}",
		controllers.GetPostByID,
	},
}
