package present

import (
	"container/heap"
	"day05/internal/support"
	"errors"
	"sort"
)


func getNCoolestPresents(presents []support.Present, n int) ([]support.Present, error) {
    if n > len(presents) || n <= 0 {
        return nil, errors.New("n is nil or greater than size of heap")
    }
    ph := make(support.Presents, len(presents))
	for i, j := range presents {
		ph[i] = &support.Present{
			Value: j.Value,
			Size:  j.Size,
		}
	}

	heap.Init(&ph)

	coolestPresent := make([]support.Present, n)
	for i := 0; i < n; i++ {
		coolestPresent[i] = *(heap.Pop(&ph).(*support.Present))
	}

	return coolestPresent, nil
}

func grabPresents(presents []support.Present, memory int) ([]support.Present) {
    if memory <= 0 {
        return []support.Present{}
    }

    TemporaryMatrix := make([][]int, len(presents)+1)

	for i := range TemporaryMatrix {
		TemporaryMatrix[i] = make([]int, memory+1)
	}

	// Заполняем таблицу dp
	for numberOfElement := 1; numberOfElement <= len(presents); numberOfElement++ {
		for MaxMemoryCapacityKnapsack := 1; MaxMemoryCapacityKnapsack <= memory; MaxMemoryCapacityKnapsack++ {
			if presents[numberOfElement-1].Size <= MaxMemoryCapacityKnapsack {
				TemporaryMatrix[numberOfElement][MaxMemoryCapacityKnapsack] = max(TemporaryMatrix[numberOfElement-1][MaxMemoryCapacityKnapsack], TemporaryMatrix[numberOfElement-1][MaxMemoryCapacityKnapsack-presents[numberOfElement-1].Size]+presents[numberOfElement-1].Value)
			} else {
				TemporaryMatrix[numberOfElement][MaxMemoryCapacityKnapsack] = TemporaryMatrix[numberOfElement-1][MaxMemoryCapacityKnapsack]
			}
		}
	}

    w := memory
	selected := []support.Present{}

	for i := len(presents); i > 0 && w > 0; i-- {
		if TemporaryMatrix[i][w] != TemporaryMatrix[i-1][w] {
			selected = append(selected, presents[i-1])
			w -= presents[i-1].Size
		}
	}

    sort.Slice(selected, func(i, j int) bool {
        if selected[i].Value == selected[j].Value {
            return selected[i].Size < selected[j].Size
        }
        return selected[i].Value > selected[j].Value
    })

    return selected
}