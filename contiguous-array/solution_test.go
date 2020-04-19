package solution

import (
	"fmt"
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{0, 1},
			want: 2,
		},
		{
			nums: []int{0, 1, 0},
			want: 2,
		},
		{
			nums: []int{0, 1, 0, 1},
			want: 4,
		},
		{
			nums: []int{0, 1, 1, 0},
			want: 4,
		},
		{
			nums: []int{0, 1, 0, 0, 0, 1, 0, 1},
			want: 4,
		},
		{
			nums: []int{0, 1, 1, 0, 0, 1, 1, 0, 1, 0},
			want: 10,
		},
		{
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1},
			want: 0,
		},
		{
			nums: []int{0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			nums: []int{0, 0, 0, 1, 1, 1, 0},
			want: 6,
		},
		{
			nums: []int{0, 0, 1, 1},
			want: 4,
		},
		{
			nums: []int{1, 1, 0, 0},
			want: 4,
		},
		{
			nums: []int{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			want: 4,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := findMaxLength(tt.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
