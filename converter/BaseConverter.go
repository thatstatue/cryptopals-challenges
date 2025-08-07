package main

import (
	"fmt"
	"slices"
	"strconv"
)

var base64Table = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
var hexTable = []byte("0123456789abcdef")

func main() {
	fmt.Println("welcome to converter\nenter the code of the format you're converting from:",
		"\n1. binary\t2. hexadecimal\t3. decimal\t4. base64\t5. ascii")
	var n1, n2 int
	fmt.Scanln(&n1)
	fmt.Println("\nenter the code of the format you're converting to:",
		"\n1. binary\t2. hexadecimal\t3. decimal\t4. base64\t5. ascii")
	fmt.Scanln(&n2)
	if n1 == n2 {
		fmt.Println("same format")
	} else {
		fmt.Println("\nenter your input: ")
		var input string
		fmt.Scanln(&input)
		bytes := []byte(input)
		var ans string
		if len(bytes) > 0 {
			var temp []bool

			switch n1 {
			case 1:
				temp = Byte2Bool(bytes)
			case 2:
				temp = Hex2Binary(bytes)
			case 3: // temp = IntToBinary(bytes) todo
			case 4:
				temp = Base642Binary(bytes)
			case 5: //temp = ASCII2Binary(bytes) todo
			}
			switch n2 {
			case 1:
				ans = BinaryRepresent(temp)
			case 2:
				ans = string(Binary2Hex(temp))
			case 3:
				ans = strconv.Itoa(BinaryToInt(temp))
			case 4:
				ans = string(Binary2Base64(temp))
			case 5: //ans = Binary2ASCII(temp)
			}

			fmt.Println(ans)
		} else {
			fmt.Println("input is empty")
		}
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
func Byte2Bool(b []byte) []bool {
	ans := make([]bool, len(b))
	for i := 0; i < len(b); i++ {
		if b[i] == 1 || b[i] == '1' {
			ans[i] = true
		} else if b[i] == 0 || b[i] == '0' {
			ans[i] = false
		} else {
			return nil
		}
	}
	return ans
}
func BinaryToInt(binary []bool) int {
	num := 0
	for i := 0; i < len(binary); i++ {
		num *= 2
		if binary[i] {
			num += 1
		}
	}
	return num
}
func IntToBinary(i int) []bool {
	var ans []bool
	for i > 0 {
		ans = append(ans, i%2 == 1)
		i /= 2
	}
	slices.Reverse(ans)
	return ans
}
func BinaryBit2hex(binary []bool) byte {
	return hexTable[BinaryToInt(binary)]
}

func BinaryRepresent(binary []bool) string {
	//fmt.Println(len(binary) / 8)
	temp := ""
	for i := 0; i < len(binary); i++ {
		if binary[i] {
			temp += "1"
		} else {
			temp += "0"
		}
	}
	return temp
}
func Base642Binary(base64 []byte) []bool {
	var ans []bool
	for i := 0; i < len(base64); i++ {
		ans = append(ans, Base64Bit2Binary(base64[i])...)
	}
	return ans
}
func Base64Bit2Binary(b byte) []bool {
	fmt.Println(int(b), " is int of ", b)
	return IntToBinary(int(b))
}

func Hex2Binary(hex []byte) []bool {
	var ans []bool
	for i := 0; i < len(hex); i++ {
		ans = append(ans, HexBit2Binary(hex[i])...)
	}
	return ans

}
func HexBit2Binary(hexBit byte) []bool {
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
