package solution

// Maximum Subarray
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3285/

func maxSubArray(nums []int) int {
	ln := len(nums)
	max := 0
	if ln > 0 {
		max = nums[0]
	}
	for i := 0; i < ln; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		for j := i - 1; j >= 0; j-- {
			nums[j] += nums[i]
			if nums[j] > max {
				max = nums[j]
			}
		}
	}
	return max
}
