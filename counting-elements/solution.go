package solution

// Counting Elements
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3289/

// Constraints
// - 1 <= arr.length <= 1000
// - 0 <= arr[i] <= 1000
func countElements(arr []int) int {
	tmp := make([]int, 1001)

	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		tmp[v]++
	}

	count := 0
	for i := 0; i < max; i++ {
		if tmp[i] > 0 && tmp[i+1] > 0 {
			count += tmp[i]
		}
	}
	return count
}
