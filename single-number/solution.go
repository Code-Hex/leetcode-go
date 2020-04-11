package solution

// Single Number
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/

func singleNumber(nums []int) int {
	n := nums[0]
	for _, v := range nums[1:] {
		n = n ^ v
	}
	return n
}
