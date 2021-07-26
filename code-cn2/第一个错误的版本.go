package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(bad int) bool {
	if bad >= 10 {
		return true
	} else {
		return false
	}
}

func firstBadVersion(n int) int {

	if n == 1 {
		return 1
	}

	left := 1
	right := n

	for  {

		index := left + (right - left)/2
		if isBadVersion(index) == true {
			if isBadVersion(index -1) == false {
				return index
			} else {
				right = index -1
			}
		} else {
			if isBadVersion(index +1) == true {
				return index+1
			} else {
				left = index + 1
			}
		}
	}
}

func main() {

	fmt.Println(firstBadVersion(15))
}

