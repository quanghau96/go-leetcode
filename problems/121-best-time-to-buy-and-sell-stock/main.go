package main

import "fmt"

// func maxProfit(prices []int) int {
// 	minPrice := prices[0]
// 	maxProfit := 0

// 	for _, price := range prices {
// 		if minPrice > price {
// 			minPrice = price
// 		}

// 		profit := price - minPrice

// 		if maxProfit < profit {
// 			maxProfit = profit
// 		}

// 	}

// 	return maxProfit
// }

func maxProfit(prices []int) int {
	ans := 0

	for buy := 0; buy < len(prices); buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			profit := prices[sell] - prices[buy]
			if profit > ans {
				ans = profit
			}

		}
	}

	return ans
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(input)

	fmt.Println("result::: ", res)
}
