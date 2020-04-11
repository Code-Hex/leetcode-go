package solution

// Happy Number
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3284/

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	var memo = map[int]int{}
	for {
		c := calc(n)
		if c == 1 {
			return true
		}
		if _, ok := memo[c]; ok {
			return false
		}
		memo[n] = c
		n = c
	}
}

func calc(n int) int {
	sum := 0
	for v := n; v > 0; {
		balance := v % 10
		sum += balance * balance
		v = v / 10
	}
	return sum
}
