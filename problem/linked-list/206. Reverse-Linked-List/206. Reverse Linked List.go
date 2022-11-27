package _206_Reverse_Linked_List

import "Leetcode-Go/structure"

type ListNode = structure.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	node := head
	for node != nil {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}
