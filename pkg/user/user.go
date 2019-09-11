package user

import (
	db "github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/types"

	"errors"

)

type UserController struct {
	Repository db.Repository
}

//Find an user
func (uc *UserController) Login(u types.User) ( error) {
	storedUser, err := uc.Repository.FindUser(u.UserName)
	if err != nil{
		return  errors.New("user not found")
	}
	if storedUser.Password != u.Password{
		return errors.New("wrong password")
	}
	return nil
}

//Insert a user
func (uc *UserController) Insert(u types.User) error {
	if u.Password != u.PasswordConfirm{
		return errors.New("passwords don't match")
	}
	if _, err := uc.Repository.FindUser(u.UserName); err != nil{
		return errors.New("user already exists")
	}
	return uc.Repository.InsertUser(u)
}

//ChangePassword update the user's password
func (uc *UserController) ChangePassword(u types.User) error {
	if u.Password != u.PasswordConfirm{
		return errors.New("passwords don't match")
	}
	return uc.Repository.ChangePassword(u.UserName, u.Password)
}
