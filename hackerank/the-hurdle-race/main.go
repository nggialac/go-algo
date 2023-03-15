package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var max int32
	for _, v := range height {
		if v > max {
			max = v
		}
	}
	if k > max {
		return 0
	}

	return max - k
}
func main() {
	var height []int32 = []int32{1, 6, 3, 5, 2}
	res := hurdleRace(4, height)
	fmt.Println(res)
}
