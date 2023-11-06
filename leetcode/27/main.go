package main

import "fmt"

func removeElement(nums []int, val int) int {
	var count int
	var n int = len(nums)

	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	var nums = []int{3, 2, 2, 3}
	var val int = 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
