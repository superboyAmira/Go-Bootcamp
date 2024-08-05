package transmitter

import (
	"log"
	"math/rand"
	"team00/pkg/transmitter_v1"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ConnServer struct {
	transmitter_v1.UnimplementedNewConnectionServiceServer
}

// реализовываем нашу ручку
func (ser *ConnServer) RandomDataGen(req *transmitter_v1.VoidRequest, stream grpc.ServerStreamingServer[transmitter_v1.NewConnectionResponse]) error {
	ev, std := randValues()
	normDist := distuv.Normal{
		Mu: ev,
		Sigma: std,
		Src: nil,
	}
	UUIDcon := uuid.New().String()
	log.Printf("{UUID: %s} New Conn generation [EV: %f], [STD: %f]", UUIDcon, ev, std)
	
	// не очень понял понятие: 'поток записей с случайным нормальным распределением',
	// пока сделал, чтобы отправлялся стрим из 3 сообщений с одним UUID 
	// и 3мя разными числами из распределения

	// используется массив указателей по причине использования протобафом (поле state protoimpl.MessageState) мютексов (sync.Mutex)
	resp := []*transmitter_v1.NewConnectionResponse{
		{
			SessionId: UUIDcon,
			Frequency: normDist.Rand(),
			TimeUtc: timestamppb.Now(),
		},
		{
			SessionId: UUIDcon,
			Frequency: normDist.Rand(),
			TimeUtc: timestamppb.Now(),
		},
		{
			SessionId: UUIDcon,
			Frequency: normDist.Rand(),
			TimeUtc: timestamppb.Now(),
		},
	}

	for _, data := range resp {
		if err := stream.Send(data); err != nil {
			return err
		}
	}

	return nil
}

func randValues() (ev float64, std float64) {
	ev = (rand.Float64() * 20) - 10
	std = 0.3 + rand.Float64() * (1.5 - 0.3)
	return
}
