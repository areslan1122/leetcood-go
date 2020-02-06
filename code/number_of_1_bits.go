package main

import "fmt"

func hammingWeight(num uint32) int {

	weight := 0
	for {
		if num == 1 || num == 0 {
			weight += int(num)
			return weight
		}
		tmp := num % 2
		num = num / 2
		weight += int(tmp)
	}
}

func main() {

	var a uint32
	a = 7
	fmt.Println(hammingWeight(a))
}
