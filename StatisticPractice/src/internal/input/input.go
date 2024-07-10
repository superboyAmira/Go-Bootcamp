package input

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// нужно начинать поля с большой буквы для доступа к полям вне модуля!
type UsedMetrics struct {
	NeedMean   bool
	NeedMedian bool
	NeedMode   bool
	NeedSD     bool
}

func StartListening() (ret UsedMetrics, bunch []float64) {
	var flagsPtrs [4]*bool
	flagsPtrs[0] = flag.Bool("mean", true, "calculate a mean num")
	flagsPtrs[1] = flag.Bool("median", true, "calculate a median num")
	flagsPtrs[2] = flag.Bool("mode", true, "calculate a mode num")
	flagsPtrs[3] = flag.Bool("sd", true, "calculate a SD num")
	flag.CommandLine.Init("didn`t stop executing a app", flag.ContinueOnError)

	flag.Parse()

	for i, ptr := range flagsPtrs {
		if ptr != nil {
			switch i {
			case 0:
				ret.NeedMean = *ptr
			case 1:
				ret.NeedMedian = *ptr
			case 2:
				ret.NeedMode = *ptr
			case 3:
				ret.NeedSD = *ptr
			}
		}
	}

	fmt.Println("Insert please a bunch of nums:")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil || num > 100000 || num < -100000 {
			continue
		}
		bunch = append(bunch, num)
	}

	return
}
