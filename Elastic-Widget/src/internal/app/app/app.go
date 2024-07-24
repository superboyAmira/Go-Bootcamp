package app

import (
	"context"
	"goday03/src/internal/app/repository"
	"goday03/src/internal/app/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Exec() {
	// loader.LoadData()
	rep := repository.NewPlaceRepository()
	serv := service.NewPlaceService(rep)
	http.HandleFunc("/", serv.BigStorePageHandler)

	server := &http.Server{
		Addr:    ":8888",
		Handler: nil,
	}
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

