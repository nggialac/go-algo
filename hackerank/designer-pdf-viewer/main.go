package main

import "fmt"

func f(in rune) int {
	return int(in - 'a')
}

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var r []rune
	// var a []int32
	for _, v := range word {
		r = append(r, v)
	}

	var max1 int32 = 1
	for _, v := range r {
		if h[f(v)] > max1 {
			max1 = h[f(v)]
		}
	}
	return max1 * int32(len(word))
}

func main() {
	// var str string = "a"
	var h []int32 = []int32{6, 5, 7, 3, 6, 7, 3, 4, 4, 2, 3, 7, 1, 3, 7, 4, 6, 1, 2, 4, 3, 3, 1, 1, 3, 5}

	res := designerPdfViewer(h, "zbkkfhwplj")

	fmt.Println(res)
}
