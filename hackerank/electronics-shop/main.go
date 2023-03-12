package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var diff int32
	var count int
	for _, k := range keyboards {
		for _, d := range drives {
			res := b - d - k
			fmt.Println(b, k, d, res, diff)
			if res >= 0 {
				if count > 0 && res < diff {
					diff = res
				}
				if count == 0 {
					diff = res
				}
				count++
			}
		}
	}
	fmt.Println(diff, count)
	if count > 0 {
		return b - diff
	}
	return -1
}

func main() {
	kb := []int32{1, 2, 3}
	d := []int32{1, 2, 2}
	var b int32 = 4
	res := getMoneySpent(kb, d, b)
	fmt.Println(res)
}
