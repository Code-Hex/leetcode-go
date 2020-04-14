package solution

import (
	"fmt"
	"testing"
)

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		stones []int
		want   int
	}{
		{
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			stones: []int{2, 7, 4, 1, 8, 1, 3, 4, 100, 304, 1000},
			want:   566,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
