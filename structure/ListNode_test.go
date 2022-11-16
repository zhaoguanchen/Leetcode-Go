package structure

import (
	"fmt"
	"testing"
)

func Test_Arr2List(t *testing.T) {

	data := [][]int{
		{
			1, 2, 3, 4,
		},

		{
			1, 2, 3, 4, 5, 6, 7,
		},
	}

	fmt.Printf("------------------------Structure: ListNode------------------------\n")

	for _, nums := range data {
		listNode := Arr2List(nums)
		listStr := List2Array(listNode)
		fmt.Printf("【input】:%v    【ListNode】:%v    【output】:%v\n", nums, listNode, listStr)
	}
	fmt.Printf("\n\n\n")
}
