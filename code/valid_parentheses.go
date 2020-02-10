package main

import "fmt"

func isValid(s string) bool {
	myStack := []string{}
	for _, v := range s {
		fmt.Println(string(v))
		if len(myStack) == 0 {
			myStack = append(myStack, string(v))
			continue
		}
		if string(v) == ")" && myStack[len(myStack)-1] == "(" ||
			string(v) == "]" && myStack[len(myStack)-1] == "[" ||
			string(v) == "}" && myStack[len(myStack)-1] == "{" {
			myStack = myStack[:len(myStack)-1]
		} else {
			myStack = append(myStack, string(v))
		}

	}

	if len(myStack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("()"))
}
