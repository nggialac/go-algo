package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	//"pwwkew"
	if s == "" {
		return 0
	}
	var tmp map[string]int = make(map[string]int)
	var maxCount, pos int = 1, 1
	for i, v := range s {

		_, ok := tmp[string(v)]
		if ok {
			if maxCount < pos-1 {
				maxCount = pos - 1
			}
			tmp = make(map[string]int)
			pos = 1
			tmp[string(v)] = pos
			fmt.Println(string(v))
		}

		tmp[string(v)] = pos
		if i == len(s)-1 && maxCount < pos {
			maxCount = tmp[string(v)]
		}

		pos++
		fmt.Println(tmp)
	}
	return maxCount
}

func longestCommonPrefix(strs []string) string {
	var firstStr string = strs[0]
	var preNum int = 0
	for {
		for _, s := range strs {
			if preNum == len(strs) {
				return firstStr
			}
			if strings.HasPrefix(s, firstStr) {
				preNum++
				continue
			}
			firstStr = firstStr[:len(firstStr)-1]
			preNum = 0
		}
	}
}

func multiply(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			n1, _ := strconv.Atoi(string(num1[i]))
			n2, _ := strconv.Atoi(string(num2[j]))
			res[i+j+1] += n1 * n2
			res[i+j] += res[i+j+1] / 10
			res[i+j+1] %= 10
		}
	}
	var ans string
	var i int
	for ; i < len(res); i++ {
		// fmt.Println(i)

		if res[i] == 0 {
			continue
		} else {
			break
		}

	}

	for ; i < len(res); i++ {
		ans += fmt.Sprintf("%d", res[i])
	}
	return ans
}

func removeDuplicates(nums []int) int {
	var numLoop int
	var arrTmp map[int]int = make(map[int]int)
	var tmp []int
	for i := 0; i < len(nums); i++ {
		_, ok := arrTmp[nums[i]]
		if !ok {
			arrTmp[nums[i]] = i
			tmp = append(tmp, nums[i])
			continue
		}
		numLoop++
	}
	k := len(nums) - numLoop
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
	return k
}

func maxProfit(prices []int) int {
	var maximumProfit int = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maximumProfit += prices[i] - prices[i-1]
		}
	}
	return maximumProfit
}

func rotate1(nums []int, k int) {
	var n = len(nums)
	// var rot int
	var tmp []int = make([]int, n)
	copy(tmp, nums)
	// for rot < k {
	// 	fmt.Println(tmp, nums)

	// 	tmp = append(tmp, nums[len(nums)-1])
	// 	tmp = append(tmp, nums[:len(nums)-1]...)
	// 	fmt.Println(tmp, nums)
	// 	rot++
	// }
	for i := 0; i < n; i++ {
		if i+k > n-1 {
			tmp[i+k-n] = nums[i]
			continue
		}
		tmp[i+k] = nums[i]
	}
	fmt.Println(tmp, nums)
}

func rotate2(nums []int, k int) []int {
	if k < 0 || len(nums) == 0 {
		return nums
	}

	fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

	r := len(nums) - k%len(nums)
	fmt.Println(nums[r:], nums[:r])
	nums = append(nums[r:], nums[:r]...)

	fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

	return nums
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}
	var res []int
	var flag bool = false
	for i, v := range nums {
		if v == target && !flag {
			res = append(res, i)
			flag = true
		}
		if flag && v != target {
			res = append(res, i-1)
			flag = false
		}
	}
	if res != nil {
		if flag {
			res = append(res, len(nums)-1)
		}
		return res
	}
	return []int{-1, -1}
}

func search(nums []int, target int) int {
	//4,5,6,|7|,0,1,2 -> 0
	var l, r int = 0, len(nums) - 1

	for l <= r {
		var mid = (l + r) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] > target || target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[r] < target || nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func searchMatrix(matrix [][]int, target int) bool {
	mty := len(matrix)    // 3
	mtx := len(matrix[0]) // 4
	var i int = 0
	if mty > 1 {
		for ; i < mty-1; i++ {
			if matrix[i][0] <= target && matrix[i+1][0] > target {
				break
			}
		}
	}

	for j := 0; j < mtx; j++ {
		if matrix[i][j] == target {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pushLinkList(head *ListNode, val int) {
	var node *ListNode = nil
	node.Next = head
	node.Val = val
	head = node
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode = head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

func printLinkList(head *ListNode, n int) {
	for i := 0; i < n-1; i++ {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow, fast *ListNode = head, head.Next
	eHead := fast
	for eHead != nil && eHead.Next != nil {
		slow.Next = head.Next
		slow = slow.Next
		fast.Next = slow.Next
		fast = fast.Next
	}
	slow.Next = eHead

	return nil
}

func main() {
	// nums := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// res := maxProfit(nums)
	// res := rotate2(nums, 3)
	// res := searchMatrix(nums, 11)
	var head *ListNode = &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: nil,
				Val:  2,
			},
			Val: 3,
		},
		Val: 1,
	}
	// pushLinkList()
	// printLinkList(head, 3)
	res := oddEvenList(head)
	// fmt.Println(res)
	printLinkList(res, 3)
}
