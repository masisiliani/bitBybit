package user

//Reader interface for user
type Reader interface {
	Find(username string) (*User, error)
}

//Writer user
type Writer interface {
	Insert(username, password string) error
	ChangePassword(username, password string) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase service interface
type UseCase interface {
	Reader
	Writer
}
