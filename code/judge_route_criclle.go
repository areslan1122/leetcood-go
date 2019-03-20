package main

import "fmt"

func judgeCircle(moves string) bool {
	r := 0
	l := 0
	u := 0
	d := 0
	for _, i := range moves {
		if i == 82 {
			r += 1
		}
		if i == 76 {
			l += 1
		}
		if i == 85 {
			u += 1
		}
		if i == 68 {
			d += 1
		}
	}
	if (r == l) && (u == d) {
		return true
	} else {
		return false
	}
}

func main() {
	a := "LL"
	fmt.Println(judgeCircle(a))
}
