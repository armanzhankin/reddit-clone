package models

import "time"

type Post struct {
	Id        int
	Title     string
	Content   string
	AuthorID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostUpdate struct {
	Id      int
	Title   *string
	Content *string
}
