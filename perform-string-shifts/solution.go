package solution

import "log"

// Perform String Shifts
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3299/

// lowercase English letters
// And a matrix shift, where shift[i] = [direction, amount]
//
// Constraints:
// - 1 <= s.length <= 100
// - s only contains lower case English letters.
// - 1 <= shift.length <= 100
// - shift[i].length == 2
// - 0 <= shift[i][0] <= 1
// - 0 <= shift[i][1] <= 100

const (
	left  = 0
	right = 1
)

func stringShift(s string, shift [][]int) string {
	r := []byte(s)

	ln := len(s)
	for i := 0; i < len(shift); i++ {
		direction := shift[i][0]
		amount := shift[i][1]
		if direction == right {
			swap := amount % ln
			for k := 0; k < swap; k++ {
				tmp := k + 1
				r[k], r[tmp] = r[tmp], r[k]
			}
		} else {
			swap := (ln - amount) % ln
			for k := 0; k < swap; k++ {
				tmp := k + 1
				r[k], r[tmp] = r[tmp], r[k]
			}
		}

		log.Println(string(r))

	}
	return string(r)
}
