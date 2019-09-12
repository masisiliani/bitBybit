package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masisiliani/bitBybit/pkg/post"
	"github.com/masisiliani/bitBybit/pkg/user"
	"github.com/masisiliani/bitBybit/types"

	"errors"
	"fmt"
	"strconv"
)

type Controller struct {
	UserController user.UserController
	PostController post.PostController
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Working!!!")
}

func (r *Controller) Login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	var parameters types.User
	err := decoder.Decode(&parameters)
	fmt.Println(parameters)
	if err != nil {
		newError := errors.New("invalid format")
		r.writeError(w, newError)
	}

	err = r.UserController.Login(parameters)
	if err != nil {
		r.writeError(w, err)
	}

	w.Header().Set("Session", newSession(parameters.UserName))
	w.WriteHeader(http.StatusOK)
}

func (r *Controller) Register(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	var parameters types.User
	err := decoder.Decode(&parameters)
	if err != nil {
		newError := errors.New("invalid format")
		r.writeError(w, newError)
	}

	err = r.UserController.Insert(parameters)
	if err != nil {
		r.writeError(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (r *Controller) ChangePassword(w http.ResponseWriter, req *http.Request) {
	session := req.Header.Get("Session")
	u, err := decodeSession(session)

	if err != nil {
		r.writeError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&u)

	if err != nil {
		r.writeError(w, err)
	}

	err = r.UserController.ChangePassword(u)
	if err != nil {
		r.writeError(w, err)
	}

	w.WriteHeader(http.StatusOK)
}

func (r *Controller) NewPost(w http.ResponseWriter, req *http.Request) {
	session := req.Header.Get("Session")
	u, err := decodeSession(session)
	fmt.Println("username: ", u.UserName)
	if err != nil {
		r.writeError(w, err)
	}

	p := types.Post{}

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&p)

	if err != nil {
		r.writeError(w, err)
	}
	p.UserName = u.UserName
	fmt.Println(p)
	fmt.Println(p.UserName)
	err = r.PostController.InsertPost(p)
	if err != nil {
		r.writeError(w, err)
	}
	w.WriteHeader(http.StatusOK)
}

func (r *Controller) GetPosts(w http.ResponseWriter, req *http.Request) {
	session := req.Header.Get("Session")
	u, err := decodeSession(session)

	if err != nil {
		r.writeError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")

	posts, err := r.PostController.GetPosts(u.UserName)
	if err != nil {
		r.writeError(w, err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (r *Controller) GetPostByID(w http.ResponseWriter, req *http.Request) {

	fmt.Println("GetPost")

	parameters := mux.Vars(req)["postid"]

	postID, err := strconv.Atoi(parameters)

	if err != nil {
		r.writeError(w, err)
		return
	}

	fmt.Println(postID)

	session := req.Header.Get("Session")
	_, err = decodeSession(session)
	if err != nil {
		r.writeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	posts, err := r.PostController.GetPostByID(postID)
	if err != nil {
		r.writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (r *Controller) DeletePost(w http.ResponseWriter, req *http.Request) {
	session := req.Header.Get("Session")
	_, err := decodeSession(session)

	if err != nil {
		r.writeError(w, err)
	}

	idStr := req.URL.Query()["id"]
	if len(idStr) == 0 {
		newError := errors.New("invalid parameters")
		r.writeError(w, newError)
	}
	id, _ := strconv.Atoi(idStr[0])

	w.Header().Set("Content-Type", "application/json")

	err = r.PostController.DeletePost(id)
	if err != nil {
		r.writeError(w, err)
	}
	w.WriteHeader(http.StatusOK)
}

func (r *Controller) writeError(w http.ResponseWriter, err error) {
	newError := types.Error{Err: err.Error()}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(newError)
}
