package structure

type ListNode struct {
	val  int
	next *ListNode
}

// List2Array
// Convert LinkedList to Array.
func List2Array(node *ListNode) []int {
	arr := make([]int, 0)
	for {
		if node == nil {
			break
		}
		arr = append(arr, node.val)
		node = node.next
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
			val: num, next: nil,
		}
		base.next = &cur
		base = base.next
	}

	return vHead.next
}
