package user

//Reader interface
type Reader interface {
	Find(username string) (*User, error)
}

//Writer bookmark writer
type Writer interface {
	Insert(username, password string) error
	ChangePassword(username, password string) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
