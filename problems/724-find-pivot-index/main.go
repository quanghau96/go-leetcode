package main

import "fmt"

// func pivotIndex(nums []int) int {
// 	n := len(nums)

// 	for i := 0; i < n; i++ {

// 		leftSum := 0
// 		rightSum := 0

// 		for j := 0; j < i; j++ {
// 			leftSum += nums[j]
// 		}

// 		for j := i + 1; j < n; j++ {
// 			rightSum += nums[j]
// 		}

// 		if leftSum == rightSum {
// 			return i
// 		}
// 	}

// 	return -1
// } // time On2

// func pivotIndex(nums []int) int {
// 	var total, leftSum, rightSum int

// 	for _, num := range nums {
// 		total += num
// 	}

// 	for index, num := range nums {
// 		rightSum = total - leftSum - num

// 		if leftSum == rightSum {
// 			return index
// 		}

// 		leftSum += num
// 	}

// 	return -1
// } // best

func pivotIndex(nums []int) int {
	n := len(nums)

	prefix := make([]int, n)

	prefix[0] = nums[0]

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	total := prefix[n-1]

	for i := 0; i < n; i++ {
		left := 0

		if i > 0 {
			left = prefix[i-1]
		}

		right := total - prefix[i]

		if left == right {
			return i
		}

	}

	return 1
}

func main() {
	input := []int{1, 7, 3, 6, 5, 6}
	// [1,2,3]
	// [2,1,-1]
	res := pivotIndex(input)

	fmt.Println("res::: ", res)
}

// https://leetcode.com/problems/find-pivot-index
