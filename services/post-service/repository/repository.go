package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/armanzhankin/reddit-clone/services/post-service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostRepository struct {
	pg *pgxpool.Pool
}

type PostRepositoryI interface {
	CreatePost(ctx context.Context, post models.Post) (int, error)
	Update(ctx context.Context, postUpdate models.PostUpdate) error
	Delete(ctx context.Context, post_id int) error
}

func (r PostRepository) CreatePost(ctx context.Context, post models.Post) (int, error) {
	query, args, err := sq.Insert("posts").
		Columns("title", "content", "author_id").
		Values(post.Title, post.Content, post.AuthorID).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	var postId int

	if err := r.pg.QueryRow(ctx, query, args...).Scan(&postId); err != nil {
		return 0, err
	}

	return postId, nil
}
