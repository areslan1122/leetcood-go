package main

import "fmt"

func maxProfit(prices []int) int {

	sum := peerProfit(oldPrices(prices))

	for i := 1; i < len(prices); i++ {
		profit1 := peerProfit(oldPrices(prices)[:i+1])
		profit2 := peerProfit(oldPrices(prices)[i:])
		if sum < profit1+profit2 {
			sum = profit1 + profit2
		}

	}

	return sum
}

func peerProfit(prices_peer []int) int {

	sell := 0

	for i := 1; i < len(prices_peer); i++ {

		if prices_peer[i] > prices_peer[i-1] {
			if sell < (prices_peer[i] - prices_peer[i-1]) {
				sell = prices_peer[i] - prices_peer[i-1]
			}

			prices_peer[i] = prices_peer[i-1]
		}
	}

	return sell

}

func oldPrices(prices []int) []int {

	var old_prices []int
	for i := 0; i < len(prices); i++ {
		old_prices = append(old_prices, prices[i])

	}

	return old_prices
}

func main() {

	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}
