package repository

import (
	"context"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/repository/posts/models"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post models.Post) (int, error)
	//Update(ctx context.Context, postUpdate models.PostUpdate) error
	//Delete(ctx context.Context, post_id int) error
}
