package http

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)


type Handler struct {
	Router *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler{
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes(){
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World!")
	})

	h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.PostComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) Serve() error {

	go func() error {
		if err := h.Server.ListenAndServe(); err != nil {
			return err
		}

		return nil
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<- c

	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	return nil
}