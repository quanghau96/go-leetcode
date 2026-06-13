package main

import "fmt"

// func isAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	counts := make(map[rune]int)

// 	for _, value := range s {
// 		counts[value]++
// 	}

// 	for _, value := range t {
// 		counts[value]--
// 	}

// 	for _, count := range counts {
// 		if count != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// time: 3 O(n)
// space: O(n), have to store

// using bruce force

// func isAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	used := make([]bool, len(t))

// 	for i := 0; i < len(s); i++ {
// 		found := false
// 		for j := 0; j < len(t); j++ {
// 			if !used[j] && s[i] == t[j] {
// 				used[j] = true
// 				found = true
// 				break
// 			}
// 		}

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// time O(n^2)
// space O(n)

// best time On - space O1
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	ans := isAnagram("aacc", "ccac")
	fmt.Println(ans)
}
