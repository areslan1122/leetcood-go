package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	for i:=0 ; i < len(s1) ; i++ {
		for j:=0; j< len(s2); j++ {
			if s1[i] == s2[j] {
				s2 = s2[:j] + s2[j+1:]
				fmt.Println(s2)
				break
			} else if j == len(s2) - 1 {
				return false
			}
		}
	}

	return true
}

func main() {

	s1 := "cuhv"
	s2 := "cuvs"
	fmt.Println(CheckPermutation(s1,s2))
}
