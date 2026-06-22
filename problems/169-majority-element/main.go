package main

import "fmt"

// func majorityElement(nums []int) int {
// 	ctn := make(map[int]int)

// 	for _, v := range nums {
// 		ctn[v]++
// 	}

// 	max_key := 0
// 	max_value := 0

// 	for key, value := range ctn {
// 		if value > max_value {
// 			max_value = value
// 			max_key = key
// 		}
// 	}

// 	return max_key
// } time on, space on

func majorityElement(nums []int) int {
	candidate := 0
	sum := 0
	for _, value := range nums {
		if candidate == 0 {
			candidate = value
		}

		if candidate == value {
			sum++
		} else {
			sum--
		}
	}

	return sum
}

func main() {
	res := majorityElement([]int{3, 2, 3})

	fmt.Println("res::: ", res)
}

// https://leetcode.com/problems/majority-element/
