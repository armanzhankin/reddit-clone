package posts

import (
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/repository"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/service"
)

type Service struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) service.PostService {
	return &Service{
		repo: repo,
	}
}
