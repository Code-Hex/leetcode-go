package solution

// Middle of the Linked List
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3290/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	reserved := make([]*ListNode, 100)
	root := head
	i := 0
	for {
		reserved[i] = root
		if root.Next == nil {
			break
		}
		root = root.Next
		i++
	}
	len := i + 1
	mid := len / 2
	return reserved[mid]
}

func middleNode2(head *ListNode) *ListNode {
	r1, r2 := head, head
	for r2 != nil && r2.Next != nil {
		r1 = r1.Next
		r2 = r2.Next.Next
	}
	return r1
}
