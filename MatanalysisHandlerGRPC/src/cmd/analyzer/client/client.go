package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"team00/db/model/anomaly"
	"team00/db/postgresql"
	"team00/internal/cfg"
	"team00/internal/clientServices/detector"
	"team00/internal/clientServices/receiver"
	"team00/pkg/transmitter_v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initK() float64 {
	k := flag.Float64("k", 0.0, "STD coefficient")
	flag.Parse()
	return *k
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	closeClient := make(chan os.Signal, 1)
	signal.Notify(closeClient, syscall.SIGTSTP)
	go func() {
		<-closeClient
		cancel()
	}()

	// cfg
	config, err := cfg.LoadCfg("../server/cfg.xml")
	if err != nil || config == nil {
		log.Fatalf("Erorr reading config: %v", err.Error())
	}

	// connect + global ctx
	conn, err := grpc.NewClient("localhost"+config.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer conn.Close()
	client := transmitter_v1.NewConnectionServiceClient(conn)

	// db connect
	db, err := postgresql.Connect()
	if err != nil {
		log.Fatalf("Bad db connection: %v", err)
	}

	// get stream
	stream, err := client.RandomDataGen(ctx, &transmitter_v1.VoidRequest{})
	if err != nil {
		log.Fatalf("Error on calling RandomDataGen: %v", err)
	}

	diagnostic := make(chan os.Signal, 1)
	signal.Notify(diagnostic, syscall.SIGINT)
	statistic := receiver.NewReceivedData()

	anomalyBool := false
	go func() {
		<-diagnostic
		log.Println("Start anomaly detection...")
		anomalyBool = true
	}()
	pool := sync.Pool{
		New: func() any {
			return &transmitter_v1.ConnectionResponse{}
		},
	}
	k := initK()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stoping client...")
			err = stream.CloseSend()
			if err != nil {
				log.Fatalf("Err stopping server: %v", err)
			}
			return

		default:
			// здержка на получение ответа от сервера
			time.Sleep(time.Millisecond * 100)

			response := pool.Get().(*transmitter_v1.ConnectionResponse)
			response, err := stream.Recv()

			if err == io.EOF {
				log.Println("End of stream...")
				return
			}
			if err != nil {
				log.Fatalf("Error on receiving: %v", err)
			}

			if anomalyBool {
				if detector.AnomalyDetect(&response.Frequency, &k, statistic) {
					object := anomaly.NewAnomalyModel(response.SessionId, response.Frequency, statistic.Mean, statistic.STD, response.TimeUtc.AsTime())
					postgresql.TxSaveExecutor(db, object.LoadDb)
				}
			} else {
				log.Printf("{RESPONSE}: %v", response)
				statistic.Append(&response.Frequency)
			}

			response.Reset()
			pool.Put(response)
		}
	}
}
