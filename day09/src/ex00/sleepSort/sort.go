package sleepsort

import (
	"time"
)

// func sleepsort(slice []int) chan []int {
// 	result := make(chan []int)
// 	go func() {
// 		buffer := make(chan int, len(slice))
// 		sorted := []int{}

// 		routine := func(c chan int, num int, sleep time.Duration) {
// 			time.Sleep(sleep)
// 			c <- num
// 		}

// 		for i := 0; i < len(slice); i++ {
// 			go routine(buffer, slice[i], (time.Second * time.Duration(slice[i])))
// 		}
// 		for i := 0; i < len(slice); i++ {
// 			number := <-buffer
// 			sorted = append(sorted, number)
// 		}
// 		result <- sorted
// 	}()
// 	<-result
// 	return result
// }

func sleepsort(slice []int) chan []int {
	buffer := make(chan int, len(slice))
	result := make(chan []int)

	go func() {
		sorted := []int{}

		for i := 0; i < len(slice); i++ {
			number := <-buffer
			sorted = append(sorted, number)
		}
		result <- sorted
		close(result)
	}()

	routine := func(c chan int, num int, sleep time.Duration) {
		time.Sleep(sleep)
		c <- num
	}

	for i := 0; i < len(slice); i++ {
		go routine(buffer, slice[i], (time.Second * time.Duration(slice[i])))
	}

	return result
}