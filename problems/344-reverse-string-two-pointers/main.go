package main

import "fmt"

// Solve: https://leetcode.com/problems/reverse-string/description/

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// time complexity: O(n) because we need to iterate through the input array s once to reverse the string, where n is the length of the input array s. In the worst case, if the string is of length n, we will have n/2 iterations to swap the characters, which simplifies to O(n).
// space complexity: O(1) because we are not using any additional data structures that grow with the input size. We are only using a constant amount of space to store the left and right pointers.

// using recursion

// func reverseString(s []byte) {
// 	reverse(s, 0, len(s)-1)
// }

// func reverse(s []byte, left, right int) {
// 	if left >= right {
// 		return
// 	}
// 	s[left], s[right] = s[right], s[left]
// 	reverse(s, left+1, right-1)
// }

// time complexity: O(n) because we need to iterate through the input array s once to reverse the string, where n is the length of the input array s. In the worst case, if the string is of length n, we will have n/2 recursive calls to swap the characters, which simplifies to O(n).
// space complexity: O(n) because in the worst case, if the string is of length n, we will have n/2 recursive calls on the call stack, which simplifies to O(n).

// using another array
// func reverseString(s []byte) {
// 	ans := make([]byte, len(s))
// 	for index, value := range s {
// 		ansIndex := len(s) - 1 - index
// 		ans[ansIndex] = value
// 	}
// 	copy(s, ans)
// }

// time complexity: O(n) because we need to iterate through the input array s once to fill the ans array, where n is the length of the input array s. After filling the ans array, we also need to copy the contents of ans back to s, which takes O(n) time. Therefore, the overall time complexity is O(n).
// space complexity: O(n) because we create a new array ans of the same length as the input array s to store the reversed string, resulting in O(n) additional space usage.


// func reverseString(s []byte) {
// 	n := len(s)
// 	for i := 0; i < n/2; i++ {
// 		s[i], s[n-1-i] = s[n-1-i], s[i]
// 	}
// }


// steps
// 1. We initialize two pointers, left and right, to the start and end of the input array s, respectively.
// 2. We enter a loop that continues until the left pointer is less than the right pointer.


// time complexity: O(n) because we need to iterate through the first half of the input array s to reverse the string, where n is the length of the input array s. In the worst case, if the string is of length n, we will have n/2 iterations to swap the characters, which simplifies to O(n).
// space complexity: O(1) because we are not using any additional data structures that grow with the input size. We are only using a constant amount of space to store the loop variable i and the length of the input array n.

func main() {
	input := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(input)
	fmt.Println(string(input)) // Output: "olleh"
}
