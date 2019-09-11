package router

import (
	"net/http"
	"encoding/json"
)


type NewUserParams struct{
	Username string `json:"username"`
	Password  string`json:"password"`
	PasswordConfirmation string `json:"password2"`

}
func NewUser(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
    var parameters NewUserParams
    err := decoder.Decode(&parameters)
    if err != nil {
        panic(err)
    }
    User
}

// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	
// })