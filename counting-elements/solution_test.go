package solution

import (
	"fmt"
	"testing"
)

func Test_countElements(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{1, 2, 3},
			want: 2,
		},
		{
			args: []int{1, 1, 3, 3, 5, 5, 7, 7},
			want: 0,
		},
		{
			args: []int{1, 3, 2, 3, 5, 0},
			want: 3,
		},
		{
			args: []int{1, 1, 2, 2},
			want: 2,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := countElements(tt.args); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
