package main

import (
	"fmt"
	"sort"
)

// bruce force
// func intersection(nums1 []int, nums2 []int) []int {
// 	m := make(map[int]bool)

// 	for _, v := range nums1 {
// 		m[v] = true
// 	}

// 	resultMap := make(map[int]bool)

// 	for _, v := range nums2 {
// 		if m[v] {
// 			resultMap[v] = true
// 		}
// 	}

// 	var ans []int
// 	for key := range resultMap {
// 		ans = append(ans, key)
// 	}

// 	return ans
// }

// only use 1 map and delete

// func intersection(nums1 []int, nums2 []int) []int {
// 	m := make(map[int]bool)
// 	// set := make(map[int]struct{})
// 	// can use struct cuz it saves the memeory

// 	for _, v := range nums1 {
// 		m[v] = true
// 	}

// 	var ans []int
// 	for _, v := range nums2 {
// 		if m[v] {
// 			ans = append(ans, v)
// 			delete(m, v)
// 		}
// 	}

// 	return ans
// }

// binary search

// func intersection(nums1 []int, nums2 []int) []int {
// 	set := make(map[int]struct{})

// 	var ans []int

// 	for _,v := range nums1 {
// 		set[v] = struct{}{}

// 	for _, v  := range nums2 {
// 		if _, ok:=m[v]; ok {
// 			ans = append(ans, v)
// 			delete(m,v )
// 		}
// 	}

// 	return ans
// }

// import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums2)
	m := make(map[int]bool)
	var ans []int

	for _, v := range nums1 {
		if m[v] {
			continue
		}

		idx := sort.SearchInts(nums2, v)

		if idx < len(nums2) && nums2[idx] == v {
			ans = append(ans, v)
			m[v] = true
		}
	}

	return ans
}

func main() {
	input1 := []int{1, 2, 2, 1}
	input2 := []int{2, 2}

	ans := intersection(input1, input2)
	fmt.Println("ans:::, ", ans)
}

// solve: https://leetcode.com/problems/intersection-of-two-arrays/
