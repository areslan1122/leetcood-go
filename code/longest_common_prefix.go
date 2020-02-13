package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	prefix := []byte{}

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for v, k := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if v == len(strs[i]) || byte(k) != strs[i][v] {
				return string(prefix)
			}
		}
		if byte(k) == strs[1][v] {
			prefix = append(prefix, byte(k))
		}

	}

	return string(prefix)
}

func main() {

	fmt.Println(longestCommonPrefix([]string{"fowa", "flowb", "flowcd"}))
}
