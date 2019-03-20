package main

import "fmt"
import "strconv"

func fizzBuzz(n int) []string {
	var nums []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			nums = append(nums, "FizzBuzz")
		} else if i%5 == 0 {
			nums = append(nums, "Buzz")
		} else if i%3 == 0 {
			nums = append(nums, "Fizz")
		} else {
			nums = append(nums, strconv.Itoa(i))
		}
	}
	return nums

}

func main() {
	n := 16
	fmt.Println(fizzBuzz(n))

}
