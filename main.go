package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// ReverseKGroup reverses the list k nodes at a time; a final group shorter
// than k is left as-is. Each group is reversed in place by repeatedly moving
// the next node to the group front.
func ReverseKGroup(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	dummy := &Node{Next: head}
	groupPrev := dummy
	for n >= k {
		cur := groupPrev.Next // will end as the group's tail
		next := cur.Next
		for i := 1; i < k; i++ {
			cur.Next = next.Next
			next.Next = groupPrev.Next
			groupPrev.Next = next
			next = cur.Next
		}
		groupPrev = cur
		n -= k
	}
	return dummy.Next
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	fmt.Println(toSlice(ReverseKGroup(fromSlice([]int{1, 2, 3, 4, 5}), 2)))
}
