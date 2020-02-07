package main

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	return 150094635296999121%n == 0
}
