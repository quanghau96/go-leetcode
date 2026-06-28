package main

import "fmt"

// type NumArray struct {
// 	prefix []int
// }

// func Constructor(nums []int) NumArray {
// 	prefix := make([]int, len(nums)+1)

// 	for index, value := range nums {
// 		prefix[index+1] = prefix[index] + value
// 	}

// 	return NumArray{prefix: prefix}
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	return this.prefix[right+1] - this.prefix[left]
// }

// Contructor: O(n)
// SumRange: O(1)
// Space: O(1)

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	return NumArray{prefix: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0

	for i := left; i <= right; i++ {
		sum += this.prefix[i]
	}

	return sum
}

// Constructor: O(1)
// SumRange: O(n)
// Space: O(1)

func main() {
	constructor := Constructor([]int{-2, 0, 3, -5, 2, -1})

	res := constructor.SumRange(0, 2)

	fmt.Println("res:::", res)
}

// https://leetcode.com/problems/range-sum-query-immutable/
