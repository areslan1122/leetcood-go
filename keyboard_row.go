package main

import "fmt"

func findWords(words []string) []string {
	keyboard1 := "qwertyuiopQWERTYUIOP"
	keyboard2 := "asdfghjklASDFGHJKL"
	keyboard3 := "zxcvbnmZXCVBNM"
	keyboard := make(map[string]int)
	for _, i := range keyboard1 {
		keyboard[string(i)] = 1
	}
	for _, i := range keyboard2 {
		keyboard[string(i)] = 2
	}
	for _, i := range keyboard3 {
		keyboard[string(i)] = 3
	}

	var new_words []string
	for _, i := range words {
		sum := 0
		for _, j := range i {
			sum += keyboard[string(j)]
		}
		if sum == len(i) || sum == len(i)*2 || sum == len(i)*3 {
			new_words = append(new_words, i)
		}
	}

	return new_words
}

func main() {
	var a []string
	a = append(a, "Hello", "Dad", "alska")
	fmt.Println(findWords(a))
}
