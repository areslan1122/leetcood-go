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



//func lengthOfLongestSubstring1(s string) int {
//	start, end := 0, 0
//	for i := 0; i < len(s); i++ {
//		index := strings.Index(s[start:i], string(s[i]))
//		if index == -1 {
//			if i+1 > end {
//				end = i + 1
//			}
//		} else {
//			start += index + 1
//			end += index + 1
//		}
//	}
//	return end - start
//}