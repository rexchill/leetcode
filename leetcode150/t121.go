package main

// T121:买卖股票的最佳时机
func maxProfit(prices []int) int {
	ans := 0
	premin := prices[0]
	for _, v := range prices {
		ans = max(ans, v-premin)
		premin = min(premin, v)
	}
	return ans
}
