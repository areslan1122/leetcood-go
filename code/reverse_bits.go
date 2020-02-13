package main

import "fmt"

func reverseBits(num uint32) uint32 {

	i := [32]uint32{}
	index := 31

	for num != 0 {
		i[index] = num % 2
		index--
		num /= 2
	}
	fmt.Println(i)

	var num1 uint32 = 0
	num1 += 1 * i[0]
	var add uint32 = 2
	for j := 1; j < 32; j++ {

		num1 += (add * i[j])
		add *= 2
	}

	return num1
}

func main() {

	fmt.Println(reverseBits(4294967293))
}
