package main

import (
	"fmt"
)

// func sortedSquares(nums []int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		nums[i] *= nums[i]
// 	}

// 	sort.Ints(nums)

// 	return nums
// }

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	left, right := 0, n-1
	pos := n - 1

	for left <= right {

		v_left := nums[left] * nums[left]
		v_right := nums[right] * nums[right]

		if v_left < v_right {
			res[pos] = v_right
			right--
		} else {
			res[pos] = v_left
			left++
		}
		pos--

	}

	return res
}

// time: On, space On

func main() {
	input := []int{-7, -3, 2, 3, 11}
	res := sortedSquares(input)

	fmt.Println("res:::: ", res)
}

// solve: https://leetcode.com/problems/squares-of-a-sorted-array/submissions/2045584811/
