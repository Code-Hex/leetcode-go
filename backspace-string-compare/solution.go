package solution

// Backspace String Compare
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3291/

// Note
//
// 1 <= S.length <= 200
// 1 <= T.length <= 200
// S and T only contain lowercase letters and '#' characters.
func backspaceCompare(S string, T string) bool {
	return emulateBackspace(S) == emulateBackspace(T)
}

func emulateBackspace(s string) int {
	ln := len(s)
	sum, bkCnt := 0, 0
	for i := ln - 1; i >= 0; i-- {
		c := s[i]
		switch c {
		case '#':
			bkCnt++
		default:
			n := int(c)
			if bkCnt > 0 {
				bkCnt--
			} else {
				sum += n
			}
		}
	}
	return sum
}
