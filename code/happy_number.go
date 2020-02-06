package main

import "fmt"

func isHappy(n int) bool {

	tmp := []int{n}
	for {
		n = calculate(n)
		if n == 1 {
			return true
		}
		for _, v := range tmp {
			if n == v {
				return false
			}
		}
		tmp = append(tmp, n)
	}
}

func calculate(i int) int {

	sum := 0
	for {
		if i < 10 {
			sum += i * i
			break
		}

		j := i % 10
		sum += j * j
		i = i / 10
	}

	return sum
}

func main() {

	fmt.Println(isHappy(2))
}
