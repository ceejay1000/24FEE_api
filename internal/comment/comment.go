package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("failed to fetch comment bu id")
	ErrorNotImplemented = errors.New("not implemented")
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Defines all the methods the service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) ( error)
}

// Service - is the struct on which all logic will be build
type Service struct{
	Store Store
}

func NewService(store Store) *Service{
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving comment")

	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}

	return cmt, nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) error {
	return ErrorNotImplemented
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, id, cmt)

	if err != nil {
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) ( error) {
	return  ErrorNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {

	insertedCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}