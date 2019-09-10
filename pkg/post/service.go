package post

import "time"

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

//Find a post
func (s *Service) Find(ID int) (*Post, error) {
	return s.repo.Find(ID)
}

//FindByUser a post
func (s *Service) FindByUser(username string) ([]*Post, error) {
	return s.repo.FindByUser(username)
}

//Insert a post
func (s *Service) Insert(post *Post) error {
	post.Date = time.Now().String()

	return s.repo.Insert(post)
}

//Update a post
func (s *Service) Update(post *Post) error {
	return s.repo.Update(post)
}

//Delete a post
func (s *Service) Delete(ID int) error {
	return s.repo.Delete(ID)
}
