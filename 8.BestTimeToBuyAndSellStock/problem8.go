/*
Best time to buy and sell stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/

package main

import "fmt"

func maxProfit(prices []int) int {
	max_profit := 0
	n := len(prices)
	min_price := prices[0]
	for i := 1; i < n; i++ {
		if prices[i]-min_price > max_profit {
			max_profit = prices[i] - min_price
		}
		if min_price > prices[i] {
			min_price = prices[i]
		}
	}
	return max_profit

}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}
