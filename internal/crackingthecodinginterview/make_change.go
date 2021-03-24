package crackingthecodinginterview

import "fmt"

// MakeChange return number of ways to make money with given coins
func MakeChange(coins []int, money int) int {
	return makeChange(coins, money, 0, make(map[string]int))
}

func makeChange(coins []int, money, index int, memo map[string]int) int {
	if money == 0 {
		return 1
	}

	if index >= len(coins) {
		return 0
	}

	key := fmt.Sprintf("%d-%d", money, index)

	if _, found := memo[key]; found {
		return memo[key]
	}

	amountWithCoin, ways := 0, 0

	for amountWithCoin <= money {
		remaining := money - amountWithCoin
		ways += makeChange(coins, remaining, index+1, memo)
		amountWithCoin += coins[index]
	}

	memo[key] = ways

	return memo[key]
}
