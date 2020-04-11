package solution

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{2, 2, 1},
			want: 1,
		},
		{
			args: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			args: []int{1, 2, 1, 2, 5, -8, 9, 9, -8, 5, 9},
			want: 9,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := singleNumber(tt.args); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
