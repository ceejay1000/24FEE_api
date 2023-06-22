package main

import (
	"context"
	"fmt"

	"github.com/ceejay1000/go-rest-24FEE-api/internal/comment"
	db "github.com/ceejay1000/go-rest-24FEE-api/internal/database"
	"github.com/google/uuid"
)

func Run() error {
	fmt.Println("Starting up the application")

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(), 
		comment.Comment{
			ID: uuid.NewString(),
			Slug: "manual-test",
			Author: "Elliot",
			Body: "Hello World",
		},
	)

	fmt.Println(cmtService.GetComment(context.Background(), ""))

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("Connected to databse succesfully")
	return nil
}

func main() {
	fmt.Println("Hello API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}