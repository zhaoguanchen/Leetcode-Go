package _001_Two_Sum

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// first 是参数
// second 代表第一个参数
type para1 struct {
	first  []int
	second int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	first []int
}

func Test_Problem66(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{1, 2, 3}, 1},
			ans1{[]int{1, 2, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum(p.first, p.second))
	}
	fmt.Printf("\n\n\n")
}
