package main

func mySqrt(x int) int {

	i := 0
	for {
		if i*i > x {
			return i - 1
		}
		i++
	}
}
