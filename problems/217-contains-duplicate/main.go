package main

import (
	"fmt"
	"sort"
)

// func containsDuplicate(nums []int) bool {
// 	arr := make([]int, 0)

// 	for _, value := range nums {
// 		for _, v := range arr {
// 			if v == value {
// 				return true
// 			}
// 		}
// 		arr = append(arr, value)
// 	}

// 	return false
// }

// time complexity: O(n^2) because we have a nested loop where the outer loop iterates through each element in the input array nums and the inner loop iterates through the arr array to check for duplicates. In the worst case, if all elements in nums are unique, the inner loop will iterate through all elements in arr for each element in nums, resulting in n * n = n^2 iterations.
// space complexity: O(n) because in the worst case, if all elements in nums are unique, we will store all n elements in the arr array, resulting in O(n) additional space usage.

// func containsDuplicate(nums []int) bool {
// 	seen := make(map[int]bool)
// 	for _, value := range nums {
// 		if seen[value] {
// 			return true
// 		}
// 		seen[value] = true
// 	}
// 	return false
// }

// time complexity: O(n) because we iterate through the nums array once to check for duplicates, where n is the length of the input array nums. The map operations (checking if a value exists and adding a value) are on average O(1), so the overall time complexity is O(n).
// space complexity: O(n) because in the worst case, if all elements in nums are unique, we will store all n elements in the seen map, resulting in O(n) additional space usage.

// func containsDuplicate(nums []int) bool {
// 	seen := make(map[int]struct{})
// 	for _, value := range nums {
// 		if _, exists := seen[value]; exists {
// 			return true
// 		}
// 		seen[value] = struct{}{}
// 	}
// 	return false
// }

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for index, value := range nums {
		if index > 0 && value == nums[index-1] {
			return true
		}
	}
	return false
}

// time complexity: O(n log n) because we sort the input array nums, which takes O(n log n) time. After sorting, we iterate through the sorted array once to check for duplicates, which takes O(n) time. Therefore, the overall time complexity is dominated by the sorting step, resulting in O(n log n).
// space complexity: O(1) if we ignore the space used by the sorting algorithm (which is typically O(log n) for in-place sorting algorithms like quicksort). If we consider the space used by the sorting algorithm, then the space complexity would be O(log n) due to the recursive stack space used by the sorting algorithm. However, since we are not using any additional data structures that grow with the input size, we can say that the space complexity is O(1) for the purpose of this problem.

func main() {
	ans := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(ans)

	// arr := make([]int, 0)
}

// solve: https://leetcode.com/problems/contains-duplicate/description/
