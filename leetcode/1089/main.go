package main

import "fmt"

func sort(arr []int, pos int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if i == pos {
			arr[i+1] = arr[i]
		}
	}
	fmt.Println(arr)
	return arr
}

func duplicateZeros(arr []int) {
	arrLength := len(arr)
	var i int
	// var tmp int
	for ; i < arrLength-1; i++ {
		if arr[i] == 0 { //0
			// tmp = arr[i+1]
			// arr[i+1] = 0
			sort(arr, i)
		}
	}
	// fmt.Println(arr)
}

func main() {
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
}
