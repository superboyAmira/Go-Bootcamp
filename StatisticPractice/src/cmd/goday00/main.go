package main

import (
	"fmt"
	"goday00/internal/calculation"
	"goday00/internal/input"
)

func main() {
	metrics, bunch := input.StartListening()
	result := calculation.GetMetrics(metrics, bunch)
	if result.Mean != nil {
		fmt.Println("Mean: ", *result.Mean)
	}
	if result.Median != nil {
		fmt.Println("Median: ", *result.Median)
	}
	if result.Mode != nil {
		fmt.Println("Mode: ", *result.Mode)
	}
	if result.SD != nil {
		fmt.Println("SD: ", *result.SD)
	}
}
