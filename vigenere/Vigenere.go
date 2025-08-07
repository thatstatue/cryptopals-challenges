package main

func main() {

}

func HammingDistance(b1, b2 []bool) int {
	d := 0
	if len(b1) != len(b2) {
		return -1
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			d++
		}
	}
	return d
}
