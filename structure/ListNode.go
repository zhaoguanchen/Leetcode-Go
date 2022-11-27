package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Array
// Convert LinkedList to Array.
func List2Array(node *ListNode) []int {
	arr := make([]int, 0)
	for {
		if node == nil {
			break
		}
		arr = append(arr, node.Val)
		node = node.Next
	}

	return arr
}

// Arr2List
// build LinkedList
func Arr2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	vHead := &ListNode{}
	base := vHead

	for _, num := range arr {
		cur := ListNode{
			Val: num, Next: nil,
		}
		base.Next = &cur
		base = base.Next
	}

	return vHead.Next
}
