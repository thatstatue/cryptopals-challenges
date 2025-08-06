package main

import (
	"fmt"
	"slices"
)

var base64Table = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
var hexTable = []byte("0123456789abcdef")

func main() {
	fmt.Printf("input a hex based string and receive it in Base 64\n:")
	var input string
	fmt.Scanln(&input)
	hex := []byte(input)
	if len(hex) > 0 {
		var ans = Binary2Base64(Hex2Binary(hex))
		fmt.Println(string(ans))
	} else {
		fmt.Println("input is empty")
	}
}
func Binary2Base64(binary []bool) []byte {
	ans := []byte{}
	for i := 0; i < len(binary); i += 6 {
		end := i + 6
		if end > len(binary) {
			end = len(binary)
		}
		ans = append(ans, BinaryBit2base64(binary[i:end]))

	}
	for len(ans)%4 != 0 {
		ans = append(ans, '=')
	}
	return ans
}
func BinaryBit2base64(binary []bool) byte {
	num := 0
	for i := 0; i < len(binary); i++ {
		num *= 2
		if binary[i] {
			num += 1
		}
	}
	return base64Table[num]
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

func BinaryRepresent(binary []bool) []int {
	ans := make([]int, len(binary))
	for i := 0; i < len(binary); i++ {
		if binary[i] {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}
	return ans
}
