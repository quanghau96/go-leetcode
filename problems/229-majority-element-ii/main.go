package main

import "fmt"

// func majorityElement(nums []int) []int {
// 	if len(nums) <= 1 {
// 		return nums
// 	}

// 	var ans []int

// 	m := make(map[int]int)

// 	for _, value := range nums {
// 		m[value]++
// 	}

// 	for key, value := range m {
// 		if value > (len(nums) / 3) {
// 			ans = append(ans, key)
// 			delete(m, key)
// 		}
// 	}

// 	return ans
// }

// func majorityElement(nums []int) []int {
// 	counts := make(map[int]int)
// 	threshold := len(nums) / 3

// 	for _, value := range nums {
// 		counts[value]++
// 	}

// 	var res []int

// 	for key, count := range counts {
// 		if count > threshold {
// 			res = append(res, key)
// 		}
// 	}

// 	return res
// }

func majorityElement(nums []int) []int {
	thresold := len(nums) / 3
	var candidate1, candidate2 int
	count1, count2 := 0, 0

	for _, num := range nums {
		switch {
		case candidate1 == num:
			count1++
		case candidate2 == num:
			count2++
		case count1 == 0:
			candidate1 = num
			count1 = 1
		case count2 == 0:
			candidate2 = num
			count2 = 1
		default:
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	res := []int{}
	if count1 > thresold {
		res = append(res, candidate1)
	}

	if count2 > thresold {
		res = append(res, candidate2)
	}

	return res
}

func main() {
	input := []int{3, 3}
	res := majorityElement(input)
	fmt.Println("res::: ", res)
}
