package main

import (
	"fmt"
	"slices"
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
		b2 = append(make([]bool, d), b2...)
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			c[i] = true
		}
	}
	return c
}
func Hex2Binary(hex []byte) []bool {
	var ans []bool
	for i := 0; i < len(hex); i++ {
		ans = append(ans, HexBit2binary(hex[i])...)
	}
	return ans

}
func HexBit2binary(hexBit byte) []bool {
	var ans []bool
	b := hexBit
	for b > 0 {
		if b >= 'a' {
			b = b - 'a' + 10
		} else if b >= '0' {
			b = b - '0'
		}
		ans = append(ans, b%2 == 1)
		b /= 2
	}
	for len(ans) < 4 {
		ans = append(ans, false)
	}
	slices.Reverse(ans)
	return ans
}
func BinaryBit2hex(binary []bool) byte {
	num := 0
	for i := 0; i < len(binary); i++ {
		num *= 2
		if binary[i] {
			num += 1
		}
	}
	return hexTable[num]
}
func Binary2Hex(binary []bool) []byte {
	ans := []byte{}
	for i := 0; i < len(binary); i += 4 {
		end := i + 4
		if end > len(binary) {
			end = len(binary)
		}
		ans = append(ans, BinaryBit2hex(binary[i:end]))

	}
	return ans
}
