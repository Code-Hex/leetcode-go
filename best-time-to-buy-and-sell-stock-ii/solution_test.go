package solution

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{7, 1, 5, 3, 6, 4},
			want: 7,
		},
		{
			args: []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			args: []int{7, 6, 4, 3, 1},
			want: 0,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := maxProfit(tt.args); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
