package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	n := len(numbers)

	left, right := 0, n-1

	for left < right {

		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("res::: ", res)

	res = twoSum([]int{2, 7, 11, 15}, 18)
	fmt.Println("res::: ", res)

	res = twoSum([]int{-1, 0}, -1)
	fmt.Println("res::: ", res)

	res = twoSum([]int{5, 25, 75}, 100)
	fmt.Println("res::: ", res)
}

// SOLVE: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
