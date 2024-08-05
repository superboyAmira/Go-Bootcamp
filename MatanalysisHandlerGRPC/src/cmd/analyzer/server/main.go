package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"team00/internal/cfg"
	"team00/internal/transmitter"
	"team00/pkg/transmitter_v1"
	"time"

	"google.golang.org/grpc"
)

const (
	xml_configuration string = "cfg.xml"
)

func main() {
	config, err := cfg.LoadCfg(xml_configuration)
	if err != nil || config == nil {
		log.Fatalf("Erorr reading config: %v", err.Error())
	}

	// Использование http.Server предполагает работу с HTTP/1.1 а для 2 требует доп настройки как я понял
	// поэтому использование net предпочтительнее
	listen, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	transmitter_v1.RegisterNewConnectionServiceServer(server, &transmitter.ConnServer{})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTSTP)

	log.Printf("Server is running on port %s", config.Port)
	go func ()  {
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	<-quit
	log.Println("Shutting down server...")
	_, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	server.GracefulStop()
	log.Println("Server stopped gracefully")
}
