package service

import (
	"context"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/models"
)

type PostService interface {
	CreatePost(ctx context.Context, post models.Post) (int, error)
}
