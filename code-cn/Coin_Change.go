package main

import "fmt"

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	arr := make([]int, amount+1)

	for i := 0; i < len(arr); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			count := arr[i-coin] + 1
			if (count < arr[i] || arr[i] == 0) && (arr[i-coin] != 0 || i == coin) {
				arr[i] = count
			}
		}
	}
	if arr[amount] == 0 {
		return -1
	}
	return arr[amount]
}

func main() {

	fmt.Println(coinChange([]int{2}, 3))
}
