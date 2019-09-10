package post

//Reader interface for post
type Reader interface {
	Find(ID int) (*Post, error)
	FindByUser(username string) ([]*Post, error)
	Search(query string) ([]*Post, error)
}

//Writer struct post
type Writer interface {
	Insert(post *Post) error
	Update(post *Post) error
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
