package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByID(context.Context, int64) (*Comment, error)
		GetByPostID(context.Context, int64) ([]*Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UsersStore{db},
		Comments: &CommentStore{db},
	}
}
