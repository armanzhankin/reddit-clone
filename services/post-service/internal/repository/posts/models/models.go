package models

import "time"

type Post struct {
	ID        int64     `db:"id"`
	Info      *Info     `db:""`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Info struct {
	Title    string `db:"title"`
	Content  string `db:"content"`
	AuthorID int64  `db:"author_id"`
}
