package solution

import (
	"fmt"
	"testing"
)

func Test_isHappy(t *testing.T) {
	tests := []struct {
		arg  int
		want bool
	}{
		{
			arg:  19,
			want: true,
		},
		{
			arg:  7,
			want: true,
		},
		{
			arg:  97,
			want: true,
		},
		{
			arg:  98,
			want: false,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			if got := isHappy(tt.arg); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
