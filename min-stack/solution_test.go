package solution

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		op   []string
		args [][]int
		want []string
	}{
		{
			op:   []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			args: [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			want: []string{"null", "null", "null", "null", "-3", "null", "0", "-2"},
		},
		{
			op:   []string{"MinStack", "push", "push", "push", "getMin", "top", "pop", "getMin"},
			args: [][]int{{}, {-2}, {0}, {-1}, {}, {}, {}, {}},
			want: []string{"null", "null", "null", "null", "-2", "-1", "null", "-2"},
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			var stack MinStack
			got := make([]string, len(tt.op))
			for i := 0; i < len(tt.op); i++ {
				switch tt.op[i] {
				case "MinStack":
					stack = Constructor()
					got[i] = "null"
				case "push":
					stack.Push(tt.args[i][0])
					got[i] = "null"
				case "getMin":
					v := stack.GetMin()
					got[i] = strconv.Itoa(v)
				case "top":
					v := stack.Top()
					got[i] = strconv.Itoa(v)
				case "pop":
					stack.Pop()
					got[i] = "null"
				default:
					t.Fatalf("unexpected operation: `%s`", tt.op[i])
				}
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
