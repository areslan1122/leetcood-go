package main

import (
	"fmt"
	"strings"
)

func replaceSpaces(S string, length int) string {

	return strings.ReplaceAll(S[:length], " ", "%20")
}

func main() {

	s := "Mr John Smith  "
	fmt.Println(replaceSpaces(s, 13))

}
