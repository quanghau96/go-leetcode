package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	m := make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]]++
// 	}
// 	nums = nums[:0]
// 	for key := range m {
// 		fmt.Printf("value::: %d \n", key)
// 		nums = append(nums, key)
// 	}

// 	return len(nums)
// }

func removeDuplicates(nums []int) int {
	var k int

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	removeDuplicates(nums)

	fmt.Println("Ans:::", nums)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
