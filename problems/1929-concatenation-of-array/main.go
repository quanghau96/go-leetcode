package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	for index, value := range nums {
		ans[index] = value
		ans[index+n] = value
	}

	// or we can do this way
	// copy(ans, nums)
	// copy(ans[n:], nums)

	// for i := 0; i < n; i++ {
	// 	ans[i] = nums[i]
	// 	ans[i+n] = nums[i]
	// }

	return ans
}

// time complexity: O(n) because we iterate through the nums once to fill the ans array, where n is the length of the input array nums
// space complexity: O(n) because we create a new aray of size 2n to store the result

func main() {
	ans := getConcatenation([]int{1, 2, 1})
	fmt.Println(ans)
}

// problem https://leetcode.com/problems/concatenation-of-array/
