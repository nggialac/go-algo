package main

import "fmt"

func plusOne(digits []int) []int {
	var len int = len(digits)
	var loopCount int

	if digits[len-1] == 9 {
		for i := len - 1; i > 0; i-- {
			if digits[i] == 9 {
				loopCount += 1
				digits[i] = 0
			} else {
				break
			}
		}
	}

	digits[len-loopCount-1] += 1
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	var a = []int{9, 9, 9}
	fmt.Println(plusOne(a))
}
