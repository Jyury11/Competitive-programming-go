package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	distance := 0

	node := head
	max := 100000
	for {
		node = node.Next
		if node == nil {
			return false
		}
		if distance == max {
			return true
		}
		distance++
	}
}
