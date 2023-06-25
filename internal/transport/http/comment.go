package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ceejay1000/go-rest-24FEE-api/internal/comment"
	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CommentService interface{
	GetComment(context.Context, string) (comment.Comment, error)
	PostComment(context.Context, PostCommentRequest) (comment.Comment, error)
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
	DeleteComment(context.Context, string) ( error)
}

type Response struct {
	Message string
}

type PostCommentRequest struct {
	Slug string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body string `json:"body" validate:"required"`
}

func convertPostCommentToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug: c.Slug,
		Author: c.Author,
		Body: c.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()

	err := validate.Struct(cmt)

	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	convertedComment := convertPostCommentToComment(cmt)

	cmt, err = h.Service.PostComment(r.Context(), convertedComment)

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id) 

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt) 

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id) 

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}

}