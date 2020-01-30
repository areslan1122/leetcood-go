package main

func romanToInt(s string) int {

	sum := 0
	s = s + string('O')
	for i := 0; i < len(s); i++ {

		if string(s[i]) == string('M') {
			sum = sum + 1000
		}

		if string(s[i]) == string('D') {
			sum = sum + 500
		}

		if string(s[i]) == string('C') {
			if string(s[i+1]) == string('M') || string(s[i+1]) == string("D") {
				sum = sum - 100
			} else {

				sum = sum + 100
			}
		}

		if string(s[i]) == string('L') {
			sum = sum + 50
		}

		if string(s[i]) == string('X') {
			if string(s[i+1]) == string('L') || string(s[i+1]) == string("C") {
				sum = sum - 10
			} else {

				sum = sum + 10
			}
		}

		if string(s[i]) == string('V') {
			sum = sum + 5
		}

		if string(s[i]) == string('I') {
			if string(s[i+1]) == string('V') || string(s[i+1]) == string("X") {
				sum = sum - 1
			} else {

				sum = sum + 1
			}
		}
	}

	return sum

}
