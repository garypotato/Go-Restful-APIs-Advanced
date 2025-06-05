package main

import (
	"another-restful-api/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags,omitempty"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  1, // Assuming a static user ID for simplicity
		Tags:    payload.Tags,
	}

	if err := app.store.Posts.Create(r.Context(), post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post, err := app.store.Posts.GetByID(ctx, postID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if post == nil {
		app.notFoundHandler(w, r)
		return
	}

	comments, err := app.store.Comments.GetByPostID(ctx, postID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	post.Comments = comments

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
