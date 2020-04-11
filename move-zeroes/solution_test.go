package solution

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			args: []int{1, 2, 0, 4, 3, 0, 5, 0},
			want: []int{1, 2, 4, 3, 5, 0, 0, 0},
		},
		{
			args: []int{1, 2, 0, 0, 0, 3, 6},
			want: []int{1, 2, 3, 6, 0, 0, 0},
		},
		{
			args: []int{1, 2, 3, 0, 0, 0},
			want: []int{1, 2, 3, 0, 0, 0},
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			moveZeroes(tt.args)
			got := tt.args
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
