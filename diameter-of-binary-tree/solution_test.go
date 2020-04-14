package solution

import (
	"fmt"
	"strings"
	"testing"
)

const null = -9999999

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{1},
			want: 0,
		},
		{
			args: []int{1, 2},
			want: 1,
		},
		{
			args: []int{1, 2, 3},
			want: 2,
		},
		{
			args: []int{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6},
			want: 4,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7},
			want: 4,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 5,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 5,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			want: 5,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want: 6,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			want: 7,
		},
		{
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			want: 8,
		},
		{
			args: []int{3, 1, null, null, 2},
			want: 2,
		},
		{
			args: []int{4, 2, null, 3, 1, null, null, 5},
			want: 3,
		},
		{
			args: []int{4, -7, -3, null, null, -9, -3, 9, -7, -4, null, 6, null, -6, -6, null, null, 0, 6, 5, null, 9, null, null, -1, -4, null, null, null, -2},
			want: 8,
		},
	}
	for i, tt := range tests {
		title := fmt.Sprintf("testcase: %d", i+1)
		t.Run(title, func(t *testing.T) {
			tree := genBinaryTree(tt.args)
			if got := diameterOfBinaryTree(tree); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	v := []int{4, 2, null, 3, 1, null, null, 5}
	tree := genBinaryTree(v)
	printTree(t, "M", tree, 0)
}

func genBinaryTree(v []int) *TreeNode {
	nodes := make([]*TreeNode, len(v))
	for i := range nodes {
		if v[i] == null {
			continue
		}
		nodes[i] = &TreeNode{
			Val: v[i],
		}
	}
	cursor := 0
	for i := 0; i < len(nodes); i++ {
		vv := v[i]
		if vv == null {
			continue
		}

		left := 2*cursor + 1
		if len(v) > left {
			nodes[i].Left = nodes[left]
		}

		right := 2*cursor + 2
		if len(v) > right {
			nodes[i].Right = nodes[right]
		}
		cursor++
	}
	return nodes[0]
}

func printTree(t *testing.T, s string, tree *TreeNode, indent int) {
	t.Logf("[%s] %s: %d", s, strings.Repeat("-", indent), tree.Val)
	if tree.Left != nil {
		printTree(t, "L", tree.Left, indent+2)
	}
	if tree.Right != nil {
		printTree(t, "R", tree.Right, indent+2)
	}
}
