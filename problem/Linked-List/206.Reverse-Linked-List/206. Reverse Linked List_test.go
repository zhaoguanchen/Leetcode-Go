package _206_Reverse_Linked_List

import (
	"Leetcode-Go/structure"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "first",
			args: args{
				[]int{1, 2, 3, 4}},
			want: []int{4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(structure.Arr2List(tt.args.nums))

			for i, val := range structure.List2Array(got) {
				if val != tt.want[i] {
					t.Errorf("reverseList() = %v, want %v", structure.List2Array(got), tt.want)
					break
				}
			}

			{

			}
		})
	}
}
