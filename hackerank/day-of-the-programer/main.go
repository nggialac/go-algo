package main

import "fmt"

func dayOfProgrammer(year int32) string {
	// Write your code here
	if year == 1918 {
		return "26.09." + fmt.Sprint(year)
	}
	if year%4 == 0 && year%100 != 0 || year%400 == 0 || year%100 == 0 && year < 1918 {
		return "12.09." + fmt.Sprint(year)
	}

	return "13.09." + fmt.Sprint(year)
}

func main() {
	res := dayOfProgrammer(1800)
	fmt.Println(res)
}
