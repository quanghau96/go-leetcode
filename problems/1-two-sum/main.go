package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 	}
// 	return []int{} // This line is added to handle the case where no solution is found
// }

// time complexity: O(n^2) because we have a nested loop where the outer loop iterates through each element in the input array nums and the inner loop iterates through the remaining elements to find a pair that adds up to the target. In the worst case, if no two numbers add up to the target, we will have n * (n-1) / 2 iterations, which simplifies to O(n^2).
// space complexity: O(1) because we are not using any additional data structures that grow with the input size. We are only using a constant amount of space to store the indices of the two numbers that add up to the target.

// best option
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for index, value := range nums {
		complement := target - value
		if complementIndex, ok := seen[complement]; ok {
			return []int{complementIndex, index}
		}
		seen[value] = index
	}
	return []int{} // This line is added to handle the case where no solution is found
}

// time complexity: O(n) because we iterate through the nums array once to find the two numbers that add up to the target, where n is the length of the input array nums. The map operations (checking if a value exists and adding a value) are on average O(1), so the overall time complexity is O(n).
// space complexity: O(n) because in the worst case, if all elements in nums are unique, we will store all n elements in the seen map, resulting in O(n) additional space usage.

func main() {
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(ans)
}

// Solve: https://leetcode.com/problems/two-sum/description/

// TODO: we can also solve this problem using sorting and two pointers technique, but it will not return the original indices of the numbers, so we need to store the original indices in a separate array before sorting. The time complexity of this approach will be O(n log n) due to the sorting step, and the space complexity will be O(n) due to the additional array to store the original indices.
