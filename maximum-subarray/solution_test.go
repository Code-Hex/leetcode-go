package solution

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			args: []int{4, 8000, -12341, 299, 1000, 0, 3049, 10, -2939},
			want: 8004,
		},
		{
			args: []int{1},
			want: 1,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := maxSubArray(tt.args); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
