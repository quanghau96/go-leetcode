package main

import (
	"fmt"
)

// func moveZeroes(nums []int) {
// 	var k int

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != 0 {
// 			nums[k], nums[i] = nums[i], nums[k]
// 			k++
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	var k int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for k < len(nums) {
		nums[k] = 0
	}
}

func main() {
	nums := []int{3, 2, 2, 0, 1, 3, 0}

	moveZeroes(nums)

	fmt.Println("Ans:::", nums)
}

// Solve: https://leetcode.com/problems/move-zeroes/
