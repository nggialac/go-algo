package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32
	for i, v := range bill {
		if int32(i) == k {
			continue
		}
		sum += v
	}
	res := b - (sum / 2)
	if res > 0 {
		fmt.Println(res)
	} else if res == 0 {
		fmt.Println("Bon Appetit")
	}
}

func main() {
	bill := []int32{3, 10, 2, 9}
	var k, b int32 = 1, 12

	bonAppetit(bill, k, b)
}
