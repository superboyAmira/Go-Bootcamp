package main

import (
	"context"
	"day06/configs"
	"day06/internal/api/handlers"
	"day06/internal/api/middlewares"
	"day06/internal/repositories"
	"day06/internal/services"
	postgresql "day06/internal/storage/postgre"
	"day06/pkg/logo"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

const (
	port = ":8888"
	host = "12"
)

func main() {
	logo.Create()
	log := initLogger(envLocal)
	cfg := configs.GetConfig(log)
	if cfg == nil {
		log.Error("failed read cfg")
		return
	}
	db, err := postgresql.Connect(cfg.DSN.ToString())
	if err != nil {
		log.Error("failed connection to db: %v", err)
		return
	}

	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)
	server := http.Server {
		Addr: ":8888",
		Handler: setupHandlers(postHandler),
	}
	serverHandler(&server, log)
}

func setupHandlers(postHandler *handlers.PostHandler) *mux.Router {
	router := mux.NewRouter()	

	router.HandleFunc("/", handlers.MainIndexHandler).Methods("GET")
	router.Handle("/admin", middlewares.
		AuthMiddleware(http.HandlerFunc(handlers.AdminPanelHandler))).
		Methods("GET")
	// router.Handle("/admin/post/{id}", ).Methods("POST")
	// router.Handle("/admin/post/{id}").Methods("PUT")
	// router.Handle("/admin/post/{id}").Methods("DELETE")
	return router
}

func serverHandler(server *http.Server, log *slog.Logger) {
	quitSIG := make(chan os.Signal, 1)
	signal.Notify(quitSIG, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTSTP)

	go func() {
		log.Info("Server is running at http://127.0.0.1:8888")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("Error starting server", "error", err.Error())
		}
	}()

	<-quitSIG
	log.Info("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error("Shut down Err server", "error", err.Error())
	}
	log.Info("Server shut down!")
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