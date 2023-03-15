package main

import (
	"fmt"
)

// func updateRanking(rankedMap map[int32]int, val int32) {
// 	for k, _ := range rankedMap {
// 		if val < k {
// 			rankedMap[k]++
// 		}
// 	}
// }

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	var res []int32
	var rankedMap map[int32]int32 = make(map[int32]int32)
	var loopCount int32
	for _, v := range ranked {
		_, ok := rankedMap[v]
		if !ok {
			loopCount++
			rankedMap[v] = loopCount
		}
	}

	var rankedMax int32 = ranked[0]
	var rankedMin int32 = ranked[len(ranked)-1]
	for _, v := range player {
		if v >= rankedMax {
			res = append(res, 1)
			continue
		}
		if v < rankedMin {
			res = append(res, rankedMap[rankedMin]+1)
			continue
		}
		for i := 1; i < len(ranked)-1; i++ {
			if v == ranked[i] {
				res = append(res, rankedMap[ranked[i]])
				break
			}
			if v > ranked[i+1] && v < ranked[i] {
				res = append(res, rankedMap[ranked[i]]+1)
			}
		}
	}
	return res
}

func main() {
	res := climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102})
	fmt.Println(res)
}
