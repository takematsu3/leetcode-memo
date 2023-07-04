package main

// MaxProfit
/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
*/
func MaxProfit(prices []int) int {
	var profit int
	var minPrice int

	profit = 0
	minPrice = prices[0]

	for _, price := range prices[1:] {
		if minPrice > price {
			minPrice = price
		}

		_profit := price - minPrice
		if _profit > profit {
			profit = _profit
		}
	}

	return profit
}
