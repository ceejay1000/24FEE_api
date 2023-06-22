package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ceejay1000/go-rest-24FEE-api/internal/comment"
	uuid "github.com/google/uuid"
)

type CommentRow struct {
	ID   string
	Slug sql.NullString
	Body sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID: c.ID,
		Slug: c.Slug.String,
		Author: c.Author.String,
		Body: c.Body.String,
	}
}

func (db *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {

	var cmtRow CommentRow

	row := db.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author 
		FROM comments WHERE id $1`,
		uuid,
	)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment)(comment.Comment, error){

	id, err := uuid.NewUUID()

	if err != nil {
		fmt.Println("Unable to generate UUID")
		return comment.Comment{}, err
	}

	cmt.ID = id.String()

	postRow := CommentRow{
		ID: cmt.ID,
		Slug: sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Slug, Valid: true},
		Body: sql.NullString{String: cmt.Slug, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comment
		(id, slug, author, body)
		VALUES
		(:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows %w", err)
	}

	return cmt, nil
}