package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func reverseStr(s string, k int) string {
	strArr := []byte(s)

	for i := 0; i <= len(strArr); i += 2 * k {

		left := i
		right := min(i+k-1, len(strArr)-1)

		for left < right {
			strArr[left], strArr[right] = strArr[right], strArr[left]
			left++
			right--
		}
	}

	return string(strArr)
}

func main() {
	input := "abcdefg"
	res := reverseStr(input, 2)

	fmt.Println("ress::: ", res)
}
