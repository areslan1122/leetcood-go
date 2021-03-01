package main

import "fmt"

func oneEditAway(first string, second string) bool {

	if (len(first) - len(second)) > 1 || (len(second) - len(first)) >1 {
		return false
	}



	if len(first) == len(second) {
		i := 0
		tag := 0
		count := 0
		for i<len(first) {
			if first[i] == second[i] {
				i++
				continue
			} else {
				count += 1
				tag = i
				i++

			}

		}
		if count < 2 {
			return true
 		} else if count > 2 {
 			return false
		} else {
			if (first[tag] == second[tag-1]) && (first[tag-1] == second[tag]) {
				return true
			} else {
				return false
			}
		}
	}

	if len(first) - len(second) ==1  {
		i := 0
		j := 0
		count := 0
		fmt.Println(i,j)
		for i < len(first) && j < len(second){
			fmt.Println(i,j)
			if first[i] == second[j] {
				i++
				j++
				continue
				fmt.Println(i,j)
			} else {
				count+=1
				i++
				fmt.Println(i,j)
			}
		}
		if count > 1 {
			return false
		} else {
			return true
		}
	}
	if len(first) - len(second) == -1  {
		i := 0
		j := 0
		count := 0
		for j <len(second) && i < len(first){
			if first[i] == second[j] {
				i++
				j++
				continue
			} else {
				count+=1
				j++
			}
		}
		if count > 1 {
			return false
		} else {
			return true
		}
	}

	return false

}

func main() {
	fmt.Println(oneEditAway("ple", "pale"))
}