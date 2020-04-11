package solution

// Group Anagrams
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3288/

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0, len(strs))
	for _, v := range strs {
		for i := 0; ; i++ {
			if i == len(ret) {
				ret = append(ret, make([]string, 0))
			}

			if len(ret[i]) == 0 {
				ret[i] = append(ret[i], v)
				break
			} else {
				if isSimilar(v, ret[i][0]) {
					ret[i] = append(ret[i], v)
					break
				}
			}
		}
	}
	return ret
}

func isSimilar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	boxA := make([]int, 26)
	boxB := make([]int, 26)
	for i := 0; i < len(a); i++ {
		ai := a[i] - 'a'
		bi := b[i] - 'a'
		boxA[ai]++
		boxB[bi]++
	}

	for i := 0; i < len(boxA); i++ {
		if boxA[i] != boxB[i] {
			return false
		}
	}
	return true
}
