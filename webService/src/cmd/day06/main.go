package main

import (
	"day06/internal/api/handlers"
	"day06/internal/api/middlewares"
	"day06/internal/repositories"
	"day06/internal/services"
	postgresql "day06/internal/storage/postgre"
	"day06/pkg/logo"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Exec() {
	logo.Create()
	log := initLogger(envLocal)
	db, err := postgresql.Connect()
	if err != nil {
		log.Error("failed connection to db: %v", err)
		return
	}

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)
	setupHandlers(postHandler)
}

func setupHandlers(postHandler *handlers.PostHandler) {
	router := mux.NewRouter()	

	router.HandleFunc("/", handlers.MainIndexHandler).Methods("GET")
	router.Handle("/admin", middlewares.
		AuthMiddleware(http.HandlerFunc(handlers.AdminPanelHandler))).
		Methods("GET")
	router.Handle("/admin/post/{id}", ).Methods("POST")
	router.Handle("/admin/post/{id}").Methods("PUT")
	router.Handle("/admin/post/{id}").Methods("DELETE")
}

func initLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}