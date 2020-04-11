package solution

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		args []string
		want [][]string
	}{
		{
			args: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			args: []string{"aab", "aba", "baa", "abbbccc"},
			want: [][]string{
				{"aab", "aba", "baa"},
				{"abbbccc"},
			},
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			got := groupAnagrams(tt.args)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
