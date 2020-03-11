package main

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	dict := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !dict[i] {
			count++
			for j := 2; i*j < n; j++ {
				dict[i*j] = true
			}
		}
	}
	return count
}
