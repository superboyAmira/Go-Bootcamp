package transmitter

import (
	"context"
	"log"
	"math/rand"
	"team00/pkg/transmitter_v1"
	"time"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ConnServer struct {
	transmitter_v1.UnimplementedConnectionServiceServer
}

// реализовываем нашу ручку
func (ser *ConnServer) RandomDataGen(req *transmitter_v1.VoidRequest, stream grpc.ServerStreamingServer[transmitter_v1.ConnectionResponse]) error {
	ev, std := randValues()
	normDist := distuv.Normal{
		Mu: ev,
		Sigma: std,
		Src: nil,
	}
	UUIDcon := uuid.New().String()
	
	stopStream, cancel := context.WithTimeout(context.Background(), time.Second * 11) 
	defer cancel()

	log.Printf("{UUID: %s} {NEW CONNECTION} [EV: %f], [STD: %f]", UUIDcon, ev, std)

	for {
		select {
		case <-stopStream.Done() :
			log.Println("Stream finished (context Duration)")
			return nil

		case <-stream.Context().Done() :
			log.Println("Stream finished")
			return nil
		default : 
			if err := stream.Send(&transmitter_v1.ConnectionResponse{
				SessionId: UUIDcon,
				Frequency: normDist.Rand(),
				TimeUtc: timestamppb.Now(),
			}); err != nil {
				return err
			}
		}
	}
}

func randValues() (ev float64, std float64) {
	ev = (rand.Float64() * 20) - 10
	std = 0.3 + rand.Float64() * (1.5 - 0.3)
	return
}
