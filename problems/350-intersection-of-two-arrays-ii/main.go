package main

import (
	"fmt"
)

// func intersect(nums1 []int, nums2 []int) []int {
// 	seens := make(map[int]int)

// 	for _, num := range nums1 {
// 		seens[num]++
// 	}
// 	var res []int
// 	for _, value := range nums2 {
// 		if seens[value] > 0 {

// 			res = append(res, value)
// 			seens[value]--
// 		}
// 	}

// 	return res
// }

// Time: O(n + m)
// Space: O(n)

// func intersect(nums1 []int, nums2 []int) []int {
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)

// 	i, j := 0, 0
// 	res := make([]int, 0)

// 	for i < len(nums1) && j < len(nums2) {
// 		if nums1[i] == nums2[j] {
// 			res = append(res, nums1[i])
// 			i++
// 			j++
// 		} else if nums1[i] < nums2[j] {
// 			i++
// 		} else {
// 			j++
// 		}
// 	}

// 	return res
// }

// time Onlogn -> space O1

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	seens := make(map[int]int)

	for _, value := range nums1 {
		seens[value]++
	}

	res := make([]int, 0)

	for _, value := range nums2 {
		if seens[value] > 0 {
			res = append(res, value)
			seens[value]--
		}
	}

	return res
}

// space O(min(n,m))
// time O(n+m)

func main() {
	input1 := []int{1, 2, 2, 1}
	input2 := []int{2, 2}

	res := intersect(input1, input2)
	fmt.Println("res:::", res)
}

// solve: https://leetcode.com/problems/intersection-of-two-arrays-ii/submissions/2042368969/
