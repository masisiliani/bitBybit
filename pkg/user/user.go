package user

import (
	db "github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/types"
)

type UserController struct {
	Repository db.Repository
}

//Find an user
func (uc *UserController) Find(username string) (types.User, error) {
	return uc.Repository.FindUser(username)
}

//Insert a user
func (uc *UserController) Insert(u types.User) error {
	return uc.Repository.InsertUser(u)
}

//ChangePassword update the user's password
func (uc *UserController) ChangePassword(username, password string) error {
	return uc.Repository.ChangePassword(username, password)
}
