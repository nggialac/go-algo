package main

import "fmt"

func pageCount(n int32, p int32) int32 {
	// Write your code here
	// 6, 2 => 1
	var count int32
	if n-p > p-1 {
		n = p
		p = 1
	}
	for n > 0 {
		if n == p {
			break
		} else if n%2 == 0 && n+1 == p || n%2 == 1 && n-1 == p {
			break
		}
		n = n - 2
		count++
	}

	return count
}

func main() {
	res := pageCount(5, 4)
	fmt.Println(res)
}
