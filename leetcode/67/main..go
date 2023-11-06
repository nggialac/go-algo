package main

import (
	"fmt"
	"sort"
)

func addBinary(a string, b string) string {
	var carry byte = 0
	i := len(a) - 1
	j := len(b) - 1
	var temp []byte

	for i >= 0 || j >= 0 || carry == 1 {
		if i >= 0 {
			carry += a[i] - '0'
			i--
		}
		if j >= 0 {
			carry += b[j] - '0'
			j--
		}
		temp = append(temp, carry%2)
		carry /= 2
	}

	sort.SliceStable(temp, func(i, j int) bool {
		return i > j
	})

	return fmt.Sprintf("%v", temp)
}

func main() {
	a := "01011"
	b := "11010"
	fmt.Println(a[4] - '1')
	fmt.Println(addBinary(a, b))
}
