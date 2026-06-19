package main

import "fmt"

// func singleNumber(nums []int) int {
// 	m := make(map[int]int)

// 	for _, v := range nums {
// 		m[v]++
// 	}

// 	for key, v := range m {
// 		if v == 1 {
// 			return key
// 		}
// 	}

// 	return -1
// }

// func singleNumber(nums []int) int {
// 	set := make(map[int]bool)

// 	for _, num := range nums {
// 		if set[num] {
// 			delete(set, num)
// 		} else {
// 			set[num] = true
// 		}
// 	}

// 	for k := range set {
// 		return k
// 	}

// 	return -1
// }

// time 0n, space 0n

func singleNumber(nums []int) int {
	var result int

	for _, value := range nums {
		result ^= value
	}

	return result
}

func main() {
	res := singleNumber([]int{2, 2, 1})
	fmt.Println("value::: ", res)
}

// solve: https://leetcode.com/problems/single-number/submissions/2038322586/
