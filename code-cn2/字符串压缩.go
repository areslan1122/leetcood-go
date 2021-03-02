package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {

	if len(S) == 0 {
		return S
	}

	c := S[0]
	count := 0
	res := ""
	for i:=0 ; i<len(S); i++ {
		if S[i] == c {
			count++
		} else {
			res = (res + string(c) + strconv.Itoa(count))
			count = 1
			c = S[i]
		}
		if i == len(S)-1 {
			res = (res + string(c) + strconv.Itoa(count))
		}
	}
	if len(res) < len(S) {
		return res
	} else {
		return S
	}

}

func main() {

	fmt.Println(compressString("abcd"))
}