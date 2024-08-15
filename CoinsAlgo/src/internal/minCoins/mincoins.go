package mincoins

import (
	"slices"
	"sort"
)

func minCoins(val int, coins []int) []int {
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

func minCoins2(val int, coins []int) []int {
    // if len(coins) == 0 {
    //     return coins
    // }

    sort.Slice(coins, func(i, j int) bool {
        return coins[i] < coins[j]
    })

    coins = slices.Compact(coins)

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