package _876_Middle_of_the_Linked_List

import "Leetcode-Go/structure"

type ListNode = structure.ListNode

func middleNode(head *ListNode) *ListNode {
	p1 := head
	p2 := head

	for {
		if p2 == nil || p2.Next == nil {
			break
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
}
