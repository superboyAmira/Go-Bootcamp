// Package minCoins provides an optimized solution to the minimum coins problem.
// This implementation uses a dynamic programming approach, which differs from
// the naive recursive approach by storing intermediate results to avoid
// redundant calculations, resulting in a significant performance improvement.
package mincoins

import (
	"slices"
	"sort"
)

// Old version 
func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}


// New verision to with unsorted and compact slice
// we using a std lib quck sort and compact, whith removed all dublicates in slice
// In this relization we walk for end to start - is is more efficient, we first add the largest denominations that fit into the price tag
func MinCoins2(val int, coins []int) []int {
	coins = slices.Compact(coins)
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	res := []int{}
	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] <= val {
			val -= coins[i]
			res = append(res, coins[i])
			i++
		}
		if val == 0 {
			break
		}
	}
	return res
}

// Support func to testing adn pprof
func GenerateSlice(size int, value int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = value
	}
	return slice
}
