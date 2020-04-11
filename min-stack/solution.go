package solution

// Min Stack
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3292/

type MinStack struct {
	top *node
}

type node struct {
	val  int
	next *node
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	newNode := &node{
		val: x,
	}
	newNode.next = m.top
	if m.top != nil && m.top.min < x {
		newNode.min = m.top.min
	} else {
		newNode.min = x
	}
	m.top = newNode
}

func (m *MinStack) Pop() {
	if m.top != nil {
		m.top = m.top.next
	}
}

func (m *MinStack) Top() int {
	if m.top != nil {
		return m.top.val
	}
	return 0
}

func (m *MinStack) GetMin() int {
	if m.top != nil {
		return m.top.min
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
