package solution

// Move Zeroes
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3286/

func moveZeroes(nums []int) {
	l := len(nums)
	cursor := 0
	for i := 0; i < l; i++ {
		v := nums[i]
		if v != 0 {
			if cursor != i {
				nums[cursor] = v
				nums[i] = 0
			}
			cursor++
		}
	}
}
