package user

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Find a bookmark
func (s *Service) Find(id int) (*User, error) {
	return s.repo.Find(id)
}
