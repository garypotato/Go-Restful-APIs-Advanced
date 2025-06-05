package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `INSERT INTO comments (post_id, user_id, content) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, comment.PostID, comment.UserID, comment.Content).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetByID(ctx context.Context, id int64) (*Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at, updated_at FROM comments WHERE id = $1`
	comment := &Comment{}

	err := s.db.QueryRowContext(ctx, query, id).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Comment not found
		}
		return nil, err
	}

	return comment, nil
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) ([]*Comment, error) {
	query := `SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.id, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id = $1 ORDER BY c.created_at DESC`
	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		comment := &Comment{}
		comment.User = User{}
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.User.ID, &comment.User.Username); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
