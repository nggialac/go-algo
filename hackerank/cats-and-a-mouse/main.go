package main

import (
	"fmt"
	"math"
)

func catAndMouse(x int32, y int32, z int32) string {
	var res string
	if x > z || y > z || x == z && y == z {
		res = "Mouse C"
	}

	var disA, disB float64
	disA = math.Abs(float64(z) - float64(x))
	disB = math.Abs(float64(z) - float64(y))
	if disA < disB {
		res = "Cat A"
	} else if disB < disA {
		res = "Cat B"
	} else {
		res = "Mouse C"
	}

	return res
}

func main() {
	res := catAndMouse(1, 3, 2)
	fmt.Println(res)
}
