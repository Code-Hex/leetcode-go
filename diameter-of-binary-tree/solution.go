package solution

// Diameter of Binary Tree
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3293/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, max := maxDistance(root)
	return max
}

func maxDistance(node *TreeNode) (int, int) {
	max, ldistance, rdistance := 0, 0, 0
	if node.Left != nil {
		ld, dmax := maxDistance(node.Left)
		ldistance, max = ld+1, dmax

	}
	if node.Right != nil {
		rd, dmax := maxDistance(node.Right)
		rdistance = rd + 1
		if dmax > max {
			max = dmax
		}
	}
	if d := ldistance + rdistance; d > max {
		max = d
	}
	if ldistance > rdistance {
		return ldistance, max
	}
	return rdistance, max
}
