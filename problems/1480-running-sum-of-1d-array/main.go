package main

import "fmt"

// func runningSum(nums []int) []int {
// 	ans := make([]int, len(nums))

// 	for index, value := range nums {
// 		if index == 0 {
// 			ans[index] = value
// 		} else {
// 			ans[index] = ans[index-1] + value
// 		}
// 	}
// 	return ans
// }

// time complexity: O(n) because we iterate through the nums once to fill the ans array, where n is the length of the input array nums
// space complexity: O(n) because we create a new aray of size n to store the result

// func runningSum(nums []int) []int {
// 	for i := 1; i < len(nums); i++ {
// 		nums[i] += nums[i-1]
// 	}
// 	return nums
// }

// //time complexity: O(n) because we iterate through the nums once to update the values, where n is the length of the input array nums
// // space complexity: O(1) because we are modifying the input array in place and not using any additional data structures that grow with the input size

func runningSum(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}

// time complexity: O(n) because we iterate through the nums once to update the values, where n is the length of the input array nums
// space complexity: O(1) because we are modifying the input array in place and not using any additional data structures that grow with the input size

func main() {
	ans := runningSum([]int{1, 2, 3, 4})
	fmt.Println(ans)
}

// solve: https://leetcode.com/problems/running-sum-of-1d-array/description/
