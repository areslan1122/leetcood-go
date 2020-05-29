package main

func reverseString(s []byte) {

	if len(s) > 1 {
		var b byte
		for i := 0; i <= (len(s)-1)/2; i++ {
			b = s[i]
			s[i] = s[len(s)-1-i]
			s[len(s)-1-i] = b
		}
	}
}
