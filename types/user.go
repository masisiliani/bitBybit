package types

//User is a struct with user properties
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PasswordConfirm string `json:"passwordConf"`
}
