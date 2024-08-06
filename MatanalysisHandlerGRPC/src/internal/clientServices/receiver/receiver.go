package receiver

import (
	"log"

	"github.com/montanaflynn/stats"
)

/*
	здесь я так понимаю, мы храним все отправленные сервером данные на клиенте
	что очень странно! (аномалия же детектиться будет на клиенте,
	и промежуточные рультаты так же должны выводиться в логи клиента)
*/

type ReceivedData struct {
	count int64
	seq   []float64

	Mean float64
	STD  float64
}

func NewReceivedData() (r *ReceivedData) {
	r = &ReceivedData{
		count: 0,
		Mean:  0.0,
		STD:   0.0,
		seq:   make([]float64, 0),
	}
	return
}

// обертка добавления для промеуточного вычисления mean and STD
func (r *ReceivedData) Append(value *float64) {
	r.seq = append(r.seq, *value)
	r.count++
	r.Mean, _ = stats.Mean(r.seq)
	r.STD, _ = stats.StandardDeviation(r.seq)
	log.Printf("[%v] ---> predicted values of mean [%v] and STD [%v]", r.count, r.Mean, r.STD)
}
