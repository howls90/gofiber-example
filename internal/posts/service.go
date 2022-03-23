package posts

import "example/fiber/pkg"

type PostService struct {
	repository	*PostRepository
}

type ServiceInterface interface {
	GetPosts(offset int, limit int) (*[]Post, *pkg.MyError)
	GetPost(id string) (*Post, *pkg.MyError)
	CreatePost(post Post) (*Post, *pkg.MyError)
}

func (s *PostService) GetPosts(offset int, limit int) (*[]Post, *pkg.MyError) {
	return s.repository.GetAll(offset, limit)
}

func (s *PostService) GetPost(id string) (*Post, *pkg.MyError) {
	return s.repository.FindById(id)
}

func (s *PostService) CreatePost(post Post) (*Post, *pkg.MyError) {
	return s.repository.Create(post)
}