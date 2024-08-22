package octopus

import (
	"sync"
	"testing"
	"time"
)

type MyStruct struct {
	ID   int
	Name string
}

func TestMultiplexWithInt(t *testing.T) {
	ch1 := make(chan any)
	ch2 := make(chan any)

	go func() {
		defer close(ch1)
		ch1 <- 10
		ch1 <- 20
	}()


	go func() {
		defer close(ch2)
		ch2 <- 30
		ch2 <- 40
	}()

	multiplexedChan := Multiplex(ch1, ch2)
	results := make([]int, 0)

	timeout := time.After(1 * time.Second)

	for i := 0; i < 4; i++ {
		select {
		case val := <-multiplexedChan:
			results = append(results, val.(int))
		case <-timeout:
			t.Fatal("Test timed out")
		}
	}

	expectedResults := map[int]bool{10: true, 20: true, 30: true, 40: true}

	for _, result := range results {
		if !expectedResults[result] {
			t.Fatalf("Unexpected result: %v", result)
		}
	}
}

func TestMultiplexWithStruct(t *testing.T) {
	ch1 := make(chan any)
	ch2 := make(chan any)

	go func() {
		defer close(ch1)
		ch1 <- MyStruct{ID: 1, Name: "A"}
	}()

	go func() {
		defer close(ch2)
		ch2 <- MyStruct{ID: 2, Name: "B"}
	}()

	multiplexedChan := Multiplex(ch1, ch2)
	results := make([]MyStruct, 0)

	timeout := time.After(1 * time.Second)

	for i := 0; i < 2; i++ {
		select {
		case val := <-multiplexedChan:
			results = append(results, val.(MyStruct))
		case <-timeout:
			t.Fatal("Test timed out")
		}
	}

	expectedResults := map[int]MyStruct{
		1: {ID: 1, Name: "A"},
		2: {ID: 2, Name: "B"},
	}

	for _, result := range results {
		expectedResult, exists := expectedResults[result.ID]
		if !exists || result != expectedResult {
			t.Fatalf("Unexpected result: %+v", result)
		}
	}
}

func TestMultiplexWithDifferentTypes(t *testing.T) {
	ch1 := make(chan any)
	ch2 := make(chan any)

	go func() {
		defer close(ch1)
		ch1 <- 42
	}()

	go func() {
		defer close(ch2)
		ch2 <- "hello"
	}()

	multiplexedChan := Multiplex(ch1, ch2)
	results := make([]any, 0)

	var mu sync.Mutex
	timeout := time.After(1 * time.Second)

	for i := 0; i < 2; i++ {
		select {
		case val := <-multiplexedChan:
			mu.Lock()
			results = append(results, val)
			mu.Unlock()
		case <-timeout:
			t.Fatal("Test timed out")
		}
	}

	expectedResults := map[any]bool{
		42:      true,
		"hello": true,
	}

	for _, result := range results {
		if !expectedResults[result] {
			t.Fatalf("Unexpected result: %v", result)
		}
	}
}
