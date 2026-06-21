package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] != val {
			left++
		} else {
			nums[left], nums[right-1] = nums[right-1], nums[left]
			right--
		}
	}

	return right
}

// func removeElement(nums []int, val int) int {
// 	k := 0

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != val {
// 			nums[k] = nums[i]
// 			k++
// 		}
// 	}

// 	return k
// }

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	ans := removeElement(nums, val)

	fmt.Println("Ans:::", ans)
}

// Solve: https://leetcode.com/problems/remove-element/submissions/2040787606/
