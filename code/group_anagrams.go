package main

import (
	"fmt"
	"sort"
	"strings"
)

//Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
func groupAnagrams(strs []string) [][]string {

	res := make([][]string, 0)
	tmp := make(map[string][]string, 0)

	for _, str := range strs {
		org := str
		ss := strings.Split(str, "")
		sort.Strings(ss)
		str = strings.Join(ss, "")
		tmp[str] = append(tmp[str], org)
	}

	for _, s := range tmp {
		res = append(res, s)
	}

	return res
}

func main() {

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
