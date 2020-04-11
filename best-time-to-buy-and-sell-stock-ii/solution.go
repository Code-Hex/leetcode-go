package solution

// Best Time to Buy and Sell Stock II
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3287/

func maxProfit(prices []int) int {
	result := 0
	ln := len(prices)
	if ln == 0 {
		return 0
	}
	v := prices[0]
	p := prices[0]

	i := 0
	for i < ln-1 {
		for i < ln-1 && prices[i] >= prices[i+1] {
			i++
		}
		v = prices[i]
		for i < ln-1 && prices[i] <= prices[i+1] {
			i++
		}
		p = prices[i]
		result += p - v
	}
	return result
}
