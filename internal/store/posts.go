package store

import (
	"context"
	"database/sql"
	"encoding/json"
)

type Post struct {
	ID        int64      `json:"id"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	UserID    int64      `json:"user_id"`
	Tags      []string   `json:"tags"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	Comments  []*Comment `json:"comments"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, posts *Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	// Convert tags slice to JSON string
	tagsJSON, err := json.Marshal(posts.Tags)
	if err != nil {
		return err
	}

	err = s.db.QueryRowContext(ctx, query, posts.Content, posts.Title, posts.UserID, string(tagsJSON)).Scan(&posts.ID, &posts.CreatedAt, &posts.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT id, content, title, user_id, tags, created_at, updated_at FROM posts WHERE id = $1`
	post := &Post{}

	var tagsJSON string
	err := s.db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.Content, &post.Title, &post.UserID, &tagsJSON, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}

	// Convert JSON string back to slice
	if err := json.Unmarshal([]byte(tagsJSON), &post.Tags); err != nil {
		return nil, err
	}

	return post, nil
}
