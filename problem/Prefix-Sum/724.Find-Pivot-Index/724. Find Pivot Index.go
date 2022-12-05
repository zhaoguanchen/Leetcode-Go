package _724_Find_Pivot_Index

func pivotIndex(nums []int) int {
	n := len(nums)
	prefixLeft := make([]int, n)
	prefixLeft[0] = nums[0]
	prefixRight := make([]int, n)
	prefixRight[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		prefixLeft[i] = prefixLeft[i-1] + nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		prefixRight[i] = prefixRight[i+1] + nums[i]
	}

	for i := 0; i < n; i++ {
		if prefixLeft[i] == prefixRight[i] {
			return i
		}
	}

	return -1
}
