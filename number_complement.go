package main

import "fmt"
import "math"

func findComplement(num int) int {
	var nums [32]int
	a := 0
	for i := 0; ; i++ {
		nums[i] = num % 2
		num /= 2
		a = i
		if num == 0 {
			break
		}

	}
	res := 0
	for i := 0; i < a; i++ {
		if nums[i] == 0 {
			res += int(math.Pow(float64(2), float64(i)))
		}
	}

	return res
}

func main() {
	fmt.Println(findComplement(4))
}
