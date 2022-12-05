package _1480_Running_Sum_of_1d_Array

func runningSum(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = nums[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + nums[i]
	}

	return ans
}
