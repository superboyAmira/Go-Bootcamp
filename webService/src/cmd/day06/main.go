package main

import (
	"context"
	"day06/configs"
	"day06/internal/api/handlers"
	"day06/internal/api/middlewares"
	"day06/internal/repositories"
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


func main() {
	logo.Create()
	log := initLogger(envProd)
	cfg := configs.GetConfig(log)
	if cfg == nil {
		log.Error("failed read cfg")
		return
	}
	db, err := postgresql.Connect(cfg.DSN.ToString())

	if err != nil {
		log.Error("failed connection to db", "error", err.Error())
		return
	}

	postRepo := repositories.NewPostRepository(db)

	postHandler := handlers.NewPostHandler(postRepo)
	indexHandler := handlers.NewIndexHandler(postRepo)

	router := mux.NewRouter()

	middlewareCombo := func (fn func(w http.ResponseWriter, req *http.Request)) http.Handler {
		return middlewares.RaceLimiter(middlewares.AuthMiddleware(http.HandlerFunc(fn)))
	}

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../web/static/"))))

	router.Handle("/page/{pageNum}", middlewares.RaceLimiter(http.HandlerFunc(indexHandler.MainIndexHandler))).Methods("GET")
	router.Handle("/post/{id}", middlewares.RaceLimiter(http.HandlerFunc(postHandler.PostGet))).Methods("GET")

	router.Handle("/admin", middlewareCombo(handlers.AdminPanelHandler)).
		Methods("GET")
	router.Handle("/admin/create-post", middlewareCombo(postHandler.PostCreate)).
		Methods("POST")
	router.Handle("/admin/delete-post", middlewareCombo(postHandler.PostDelete)).
		Methods("POST")
	router.Handle("/admin/update-post", middlewareCombo(postHandler.PostUpdate)).
		Methods("POST")

	server := http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	serverHandler(&server, log)
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
