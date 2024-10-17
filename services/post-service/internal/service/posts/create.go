package posts

import (
	"context"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/models"
)

func (s Service) CreatePost(ctx context.Context, post models.Post) (int, error) {
	return s.repo.CreatePost(ctx, post)
}
