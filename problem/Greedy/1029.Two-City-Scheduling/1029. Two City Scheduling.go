package _1029_Two_City_Scheduling_

import "sort"

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	var sum = 0

	for i := 0; i < n/2; i++ {
		sum = sum + costs[i][0] + costs[n-i-1][1]
	}

	return sum
}
