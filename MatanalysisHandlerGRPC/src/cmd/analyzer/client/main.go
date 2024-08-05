package main

import (
	"context"
	"io"
	"log"
	"team00/internal/cfg"
	"team00/pkg/transmitter_v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config, err := cfg.LoadCfg("../server/cfg.xml")
	if err != nil || config == nil {
		log.Fatalf("Erorr reading config: %v", err.Error())
	}

	conn, err := grpc.NewClient("localhost"+config.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	client := transmitter_v1.NewNewConnectionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.RandomDataGen(ctx, &transmitter_v1.VoidRequest{})
	if err != nil {
		log.Fatalf("Error on calling RandomDataGen: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error on receiving: %v", err)
		}
		log.Printf("Received: %v", response)
	}
}
