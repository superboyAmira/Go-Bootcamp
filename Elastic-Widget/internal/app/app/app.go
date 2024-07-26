package app

import (
	"context"
	"goday03/src/internal/app/repository"
	"goday03/src/internal/app/service"
	"goday03/src/internal/app/service/api"
	"goday03/src/internal/app/service/recommend"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serverHandler(server *http.Server) {
	quitSIG := make(chan os.Signal, 1)
	signal.Notify(quitSIG, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTSTP)

	go func() {
		log.Println("Server is running at http://127.0.0.1:8888")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Error starting server: %v", err)
		}
	}()

	<-quitSIG
	log.Println("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Shut down Err server: %s", err.Error())
	}
	log.Println("Server shut down!")
}

func Exec() {
	// ex00
	// loader.LoadData()
	// ex01
	rep := repository.NewPlaceRepository()
	serv := service.NewPlaceService(rep)
	http.HandleFunc("/", serv.BigStorePageHandler)
	// ex02
	apiService := api.NewApi(rep)
	http.HandleFunc("/api/places", apiService.GetPlacesApiHandler)
	// ex03
	rec := recommend.NewRecommendation(rep)
	http.HandleFunc("/api/recommend", rec.GetRecommendations)
	// ex04
	http.HandleFunc("/api/get_token", func(w http.ResponseWriter, req *http.Request) {
		token := &api.TokenJWT{}
		token.Generate(w, req)
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: nil,
	}
	serverHandler(server)
}
