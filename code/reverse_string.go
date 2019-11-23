package main

//func reverseString(s string) string {
//	var new_str string
//	for i := len(s) - 1; i >= 0; i-- {
//		new_str += string(s[i])
//	}
//
//	return new_str
//
//}

func reverseString(s []byte) {
	var new_byte byte
	for i := 0; i < len(s)/2; i++ {
		new_byte = s[i]
		s[i] = s[(len(s) - 1 - i)]
		s[(len(s) - 1 - i)] = new_byte
	}
}
