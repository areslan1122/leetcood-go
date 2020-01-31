package main

import "fmt"

func maxProfit(prices []int) int {

	sell := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] > prices[i-1] {
			if sell < (prices[i] - prices[i-1]) {
				sell = prices[i] - prices[i-1]
			}

			prices[i] = prices[i-1]
		}
	}

	return sell
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
