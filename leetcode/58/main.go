package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	temp := strings.TrimSpace(s)
	var count int
	i := len(temp)
	fmt.Println(temp, len(temp))
	for i > 0 {
		if string(temp[i-1]) == " " {
			break
		}
		count++
		i--
	}
	return count
}

func main() {
	res := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Print(res)
}
