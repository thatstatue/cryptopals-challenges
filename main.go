package main

import (
	"fmt"
	"slices"
)

var base64Table = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func main() {
	fmt.Printf("input a hex based string and receive it in Base 64\n:")
	var input string
	fmt.Scanln(&input)
	hex := []byte(input)
	if len(hex) > 0 {
		var ans = binary2base64(hex2binary(hex))
		//slices.Reverse(ans)
		fmt.Println(string(ans))
	} else {
		fmt.Println("input is empty")
	}
}
func hex2binary(hex []byte) []bool {
	var ans []bool
	for i := 0; i < len(hex); i += 2 {
		ans = append(ans, hexBit2binary(hex[i:i+2])...)
	}
	return ans

}
func hexBit2binary(hexBit []byte) []bool {
	var ans []bool
	b := hexBit[1]
	for b > 0 {
		if b >= 'a' {
			b = b - 'a' + 10
		} else if b >= '0' {
			b = b - '0'
		}
		ans = append(ans, b%2 == 1)
		b /= 2
	}
	b = hexBit[0]
	for b > 0 {
		if b >= 'a' {
			b = b - 'a' + 10
		} else if b >= '0' {
			b = b - '0'
		}
		ans = append(ans, b%2 == 1)
		b /= 2
	}
	slices.Reverse(ans)
	return ans
}
func binary2base64(binary []bool) []byte {
	ans := []byte{}
	for i := 0; i < len(binary); i += 6 {
		end := i + 6
		if end > len(binary) {
			end = len(binary)
		}
		ans = append(ans, binaryBit2base64(binary[i:end]))

	}
	for len(ans)%4 != 0 {
		ans = append(ans, '=')
	}
	return ans
}
func binaryBit2base64(binary []bool) byte {
	num := 0
	for i := 0; i < len(binary); i++ {
		num *= 2
		if binary[i] {
			num += 1
		}
	}
	return base64Table[num]
}
