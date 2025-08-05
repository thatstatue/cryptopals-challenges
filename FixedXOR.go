package main

import (
	"fmt"
)

var hexTable = []byte("0123456789abcdef")

func main() {
	fmt.Printf("input the first hex based string:")
	var input1 string
	fmt.Scanln(&input1)
	fmt.Printf("input the second hex based string:")
	var input2 string
	fmt.Scanln(&input2)
	hex1 := []byte(input1)
	hex2 := []byte(input2)

	if len(hex1) > 0 && len(hex2) > 0 {

		b1 := Hex2Binary(hex1)
		b2 := Hex2Binary(hex2)
		var encoded = Binary2Hex(xor(b1, b2))
		fmt.Println(string(encoded))
	} else {
		fmt.Println("input is empty")
	}
}

func xor(b1, b2 []bool) []bool {
	if len(b1) < len(b2) {
		b1, b2 = b2, b1
	}
	c := make([]bool, len(b1))

	d := len(b1) - len(b2)
	if d > 0 {
		b2 = append(make([]bool, d), b1...)
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			c[i] = true
		}
	}
	return c
}
