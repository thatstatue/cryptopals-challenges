package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func single() {
	fmt.Printf("now, provide the encoded input in 3rd exercise to reveal the decryption:")
	var input1 string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input1 = scanner.Text()
		//fmt.Printf("You entered: %s\n", input1)

		//fmt.Printf("input the second hex based string:")
		for i := 1; i < 2; i++ {

			input2 := strings.Repeat("58", 34) // THE ANSWER FOR THIS QUESTION IS BY XORING AGAINST "58"
			//fmt.Printf("%x\n", input2)
			//fmt.Scanln(&input2)
			hex1 := []byte(input1)
			hex2 := []byte(input2)

			if len(hex1) > 0 && len(hex2) > 0 {

				b1 := Hex2Binary(hex1)
				b2 := Hex2Binary(hex2)
				var encoded = Binary2ASCII(xor(b1, b2))
				fmt.Println(string(encoded))
			} else {
				fmt.Println("input is empty")
			}
		}
	}
}
func Binary2ASCII(binary []bool) string {
	var ans []byte
	for i := 0; i < len(binary); i += 8 {
		var b byte = 0
		for j := 0; j < 8; j++ {
			b = b << 1
			if binary[i+j] {
				b |= 1
			}
		}
		ans = append(ans, b)
	}
	return string(ans)
}
