package posts

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/repository"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/repository/posts/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	pg *pgxpool.Pool
}

func NewPostRepository(pg *pgxpool.Pool) repository.PostRepository {
	return &PostRepository{
		pg: pg,
	}
}

func (r PostRepository) CreatePost(ctx context.Context, post models.Post) (int, error) {
	query, args, err := sq.Insert("posts").
		Columns("title", "content", "author_id").
		Values(post.Info.Title, post.Info.Content, post.Info.AuthorID).
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

func (r PostRepository) GetPost(ctx context.Context, id int) (models.Post, error) {
	query, args, err := sq.Select("id", "title", "content", "author_id").
		From("posts").
}
