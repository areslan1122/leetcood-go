package main

import "fmt"

func plusOne(digits []int) []int {

	tmp := []int{0}
	tmp = append(tmp, digits...)
	tmp[len(tmp)-1] += 1

	for i := len(tmp) - 1; i > 0; i-- {

		if tmp[i] == 10 {
			tmp[i] = 0
			tmp[i-1] += 1
		} else {
			break
		}
	}

	if tmp[0] != 0 {
		return tmp
	} else {
		return tmp[1:]
	}

}

func main() {

	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}
