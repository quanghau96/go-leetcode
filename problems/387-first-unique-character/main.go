package main

import (
	"fmt"
)

// func firstUniqChar(s string) int {
// 	m := make(map[rune]int)
// 	for _, key := range s {
// 		m[key]++
// 	}

// 	for index, char := range s {
// 		if m[char] == 1 {
// 			return index
// 		}
// 	}

// 	return -1
// } // time On - space On

// func firstUniqChar(s string) int {
// 	var m [26]int
// 	for i := 0; i < len(s); i++ {
// 		m[s[i]-'a']++
// 	}

// 	for i := 0; i < len(s); i++ {
// 		if m[s[i]-'a'] == 1 {
// 			return i
// 		}
// 	}

// 	return -1
// }

// try bruce force

func firstUniqChar(s string) int {
	for i := 0; i < len(s); i++ {
		var count int
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			}
		}

		if count == 1 {
			return i
		}
	}

	return -1
}

func main() {
	ans := firstUniqChar("loveleetcode")
	fmt.Println(ans)
}

// SOLVE: https://leetcode.com/problems/first-unique-character-in-a-string/description/
