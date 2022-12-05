package _876_Middle_of_the_Linked_List

import (
	"Leetcode-Go/structure"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{structure.Arr2List([]int{1, 2, 3, 4, 5})},
			want: structure.Arr2List([]int{3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); got.Val != tt.want.Val {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
