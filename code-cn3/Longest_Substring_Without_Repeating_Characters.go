package main

func lengthOfLongestSubstring(s string) int {

	rk, ans := 0,0
	m := map[byte]int{}

	for i:=0; i< len(s); i++ {

		if i != 0 {
			delete(m, s[i-1])
		}

		for rk < len(s) && m[s[rk]] ==0 {
			m[s[rk]]++
			rk++
		}

		ans = max(ans, rk-i)
	}
	return ans
}




func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}