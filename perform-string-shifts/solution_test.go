package solution

import (
	"fmt"
	"testing"
)

func Test_stringShift(t *testing.T) {
	type args struct {
		s     string
		shift [][]int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				s: "abc",
				shift: [][]int{
					{0, 1},
					{1, 2},
				},
			},
			want: "cab",
			// 			[0,1] means shift to left by 1. "abc" -> "bca"
			// [1,2] means shift to right by 2. "bca" -> "cab"
		},
		{
			args: args{
				s: "abcdefg",
				shift: [][]int{
					{1, 1},
					{1, 1},
					{0, 2},
					{1, 3},
				},
			},
			want: "efgabcd",
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := stringShift(tt.args.s, tt.args.shift); got != tt.want {
				t.Errorf("stringShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
