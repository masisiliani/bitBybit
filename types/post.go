package types

type Post struct {
	ID          int `json:"id"`
	UserName    string `json:"username"`
	Description string `json:"description"`
	Date        string `json:"date"`
}