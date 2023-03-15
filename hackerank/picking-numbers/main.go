package main

import "fmt"

func pickingNumbers(a []int32) int32 {
	// Write your code here
	var maxLength int = 0
	// var lenArr int = len(a)
	var mapArr map[int32]int = make(map[int32]int)

	for _, v := range a {
		_, ok := mapArr[v]
		if ok {
			mapArr[v] += 1
		} else {
			mapArr[v] = 1
		}
	}

	for k, v := range mapArr {
		fmt.Println(k, v, maxLength)

		if mapArr[k+1] != 0 && maxLength < v+mapArr[k+1] {
			maxLength = v + mapArr[k+1]
		}
	}

	return int32(maxLength)
}

func main() {
	var a []int32 = []int32{4, 6, 5, 3, 3, 1}
	res := pickingNumbers(a)
	fmt.Println(res)
}
