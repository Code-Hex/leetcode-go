package solution

import "sort"

// Last Stone Weight
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3297/

// Note:
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 1000
func lastStoneWeight(stones []int) int {
	for {
		sort.Ints(stones)
		ln := len(stones)
		switch ln {
		case 0:
			return 0
		case 1:
			return stones[0]
		}

		diff := stones[ln-1] - stones[ln-2]
		stones = append(stones[:ln-2], diff)
	}
}
