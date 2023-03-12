package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	count := len(arr)
	var max int
	var val int32
	var setArr map[int32]int = make(map[int32]int)
	for i := 0; i < count; i++ {
		_, ok := setArr[arr[i]]
		if ok {
			setArr[arr[i]]++
		} else {
			setArr[arr[i]] = 1
		}
	}
	fmt.Println(setArr)
	for i, v := range setArr {
		if max < v {
			max = v
			val = i
		}
	}
	return val
}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	res := migratoryBirds(arr)
	fmt.Println(res)
}
