package sleepsort

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sleepsort(slice []int) chan []int {
	buffer := make(chan int, len(slice))
	result := make(chan []int)

	

	// graseful stop all goroutine
	ctx, stop := context.WithCancel(context.Background()) 

	go func ()  {
		graceful := make(chan os.Signal, 1)
		signal.Notify(graceful, os.Interrupt, syscall.SIGTERM)
		<-graceful
		stop()
	}()
	///

	go func() {
		select {
		case <-ctx.Done():
			close(result)
			return
		default:
			sorted := []int{}

			for i := 0; i < len(slice); i++ {
				number := <-buffer
				sorted = append(sorted, number)
			}
			result <- sorted
			close(result)
		}
	}()

	routine := func(c chan int, num int, sleep time.Duration) {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(sleep)
			c <- num
		}
	}

	for i := 0; i < len(slice); i++ {
		go routine(buffer, slice[i], (time.Second * time.Duration(slice[i])))
	}

	return result
}