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

//Find an user
func (s *Service) Find(id int) (*User, error) {
	return s.repo.Find(id)
}

//Insert a user
func (s *Service) Insert(username, password string) error {
	return s.repo.Insert(username, password)
}

//ChangePassword update the user's password
func (s *Service) ChangePassword(username, password string) error {
	return s.repo.ChangePassword(username, password)
}
