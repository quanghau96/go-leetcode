package main

import (
	"fmt"
)

// func isPalindrome(s string) bool {
// 	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
// 	if err != nil {
// 		panic(err)
// 	}

// 	str := reg.ReplaceAllString(s, "")
// 	str = strings.ToLower(str)

// 	strArr := []byte(str)

// 	left, right := 0, len(strArr)-1

// 	res := true

// 	for left < right {

// 		if strArr[left] != strArr[right] {
// 			res = false
// 			break
// 		}

// 		left++
// 		right--

// 	}

// 	return res
// }

// func isPalindrome(s string) bool {
// 	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

// 	str := strings.ToLower(reg.ReplaceAllLiteralString(s, ""))

// 	left, right := 0, len(str)-1

// 	for left < right {

// 		if str[left] != str[right] {
// 			return false
// 		}

// 		left++
// 		right--

// 	}

// 	return true
// }

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}
		for left < right && !isAlphaNum(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--

	}

	return true
}

func isAlphaNum(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func toLower(c byte) byte {
	fmt.Println("c::: ", c)
	if 'A' <= c && c <= 'Z' {
		return c + ('a' - 'A')
	}

	return c
}

func main() {
	input := "A man, a plan, a canal: Panama"
	res := isPalindrome(input)

	fmt.Println("value::: ", res)

	abcd := 'A'

	fmt.Println("aaa", abcd)
}

// solve: https://leetcode.com/problems/valid-palindrome/description/
