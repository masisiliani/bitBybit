package router

import (
	"net/http"
	"encoding/json"
	"github.com/masisiliani/bitBybit/types"
	"github.com/masisiliani/bitBybit/pkg/user"
	"github.com/masisiliani/bitBybit/pkg/post"

)

type Router struct{
	UserController user.UserController
	PostController post.PostController
}

func (r *Router) Login(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
    var parameters types.User
    err := decoder.Decode(&parameters)
    if err != nil {
		newError := types.Error{Err: "invalid format"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
	}

	err = r.UserController.Login(parameters)
	if err != nil{
		newError := types.Error{Err: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
	}

	w.Header().Set("Session", newSession(parameters.UserName))	
	w.WriteHeader(http.StatusOK)
}

func (r *Router) Register(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
    var parameters types.User
    err := decoder.Decode(&parameters)
    if err != nil {
		newError := types.Error{Err: "invalid format"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
	}



	err = r.UserController.Insert(parameters)
	if err != nil{
		newError := types.Error{Err: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (r *Router) ChangePassword(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
    var parameters types.User
    err := decoder.Decode(&parameters)
    if err != nil {
		newError := types.Error{Err: "invalid format"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
	}



	err = r.UserController.ChangePassword(parameters)
	if err != nil{
		newError := types.Error{Err: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(newError)
	}

	w.WriteHeader(http.StatusOK)
}