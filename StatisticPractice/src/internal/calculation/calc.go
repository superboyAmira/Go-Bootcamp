package calculation

import (
	"goday00/internal/input"
	"math"
	"sort"
)

type ResultMetrics struct {
	Mean   *float64
	Median *float64
	Mode   *float64
	SD     *float64
}

// В Go срезы ([]float64), структуры и тп являются ссылочными типами данных, поэтому при
// передаче среза в функцию копирование всей коллекции не происходит.
// Вместо этого передается указатель на базовый массив данных и информация о длине и емкости среза
func GetMetrics(used input.UsedMetrics, bunch []float64) (res ResultMetrics) {
	if len(bunch) == 0 {
		return
	}
	sort.Float64s(bunch)
	if used.NeedMean {
		mean := calcMean(bunch)
		res.Mean = &mean
	}
	if used.NeedMedian {
		median := calcMedian(bunch)
		res.Median = &median
	}
	if used.NeedMode {
		mode := calcMode(bunch)
		res.Mode = &mode
	}
	if used.NeedSD {
		var sd float64
		if res.Mean == nil {
			mean := calcMean(bunch)
			sd = calcSD(bunch, mean)
		} else {
			sd = calcSD(bunch, *res.Mean)
		}
		res.SD = &sd
	}
	return
}

func calcMean(bunch []float64) float64 {
	sum := 0.0
	cnt := 0
	for _, num := range bunch {
		sum += num
		cnt++
	}
	return (sum / float64(cnt))
}

func calcMedian(bunch []float64) float64 {
	cnt := 0
	for _, num := range bunch {
		_ = num
		cnt++
	}
	if cnt%2 == 0 {
		return (bunch[cnt/2-1] + bunch[cnt/2]) / 2.0
	} else {
		return bunch[cnt/2]
	}
}

func calcMode(bunch []float64) float64 {
	frequency := make(map[float64]int)
	for _, num := range bunch {
		if _, ok := frequency[num]; !ok {
			frequency[num] = 1
		} else {
			frequency[num] += 1
		}
	}
	max_cnt := 0
	max_num := 0.0

	for num, cnt := range frequency {
		if max_cnt < cnt || (max_cnt == cnt && max_num > num) {
			max_num = num
			max_cnt = cnt
		}
	}
	return max_num
}

func calcSD(bunch []float64, avg float64) float64 {
	another_bunch := bunch
	for i, num := range bunch {
		another_bunch[i] = math.Pow(num-avg, 2.0)
	}
	return math.Sqrt(calcMean(bunch))
}
