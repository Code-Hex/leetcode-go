package solution

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			args: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{4, 5, 6, 7},
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			args := serialize(tt.args)
			result := middleNode(args)
			got := deserialize(result)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}

func Test_middleNode2(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			args: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{4, 5, 6, 7},
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			args := serialize(tt.args)
			result := middleNode2(args)
			got := deserialize(result)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}

func serialize(n []int) *ListNode {
	head := &ListNode{
		Val: n[0],
	}
	root := head
	for _, v := range n[1:] {
		node := &ListNode{
			Val: v,
		}
		root.Next = node
		root = node
	}
	return head
}

func deserialize(head *ListNode) []int {
	// Note: The number of nodes in the given list will be between 1 and 100.
	ret := make([]int, 0, 100)
	root := head
	for {
		ret = append(ret, root.Val)
		if root.Next == nil {
			break
		}
		root = root.Next
	}
	return ret
}
