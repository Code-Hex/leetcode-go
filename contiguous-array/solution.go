package solution

// Contiguous Array
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3298/

// Note: The length of the given binary array will not exceed 50,000.
func findMaxLength(nums []int) int {
	// 0: -1
	// 1: +1
	max := -1
	dict := make([]int, len(nums))
	for i, v := range nums {
		for j := i; j >= 0; j-- {
			switch v {
			case 0:
				dict[j]--
			case 1:
				dict[j]++
			}
			if dict[j] == 0 {
				if i-j > max {
					max = i - j
				}
			}
		}
	}
	return max + 1
}
