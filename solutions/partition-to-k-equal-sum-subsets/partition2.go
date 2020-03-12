/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package partition_to_k_equal_sum_subsets

import "sort"

func canPartition1(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}

	// 从所有数字里选出一部分，其和为sum / 2
	c := sum / 2
	dp := make([]bool, c+1)
	for j := 0; j <= c; j++ {
		dp[j] = nums[0] == j
	}
	for i := 1; i < len(nums); i++ {
		for j := c; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[c]
}

/*
贪心+DFS回溯，注意对nums降序排列，可大大减少递归次数
*/
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	halfSum, mod := sum/2, sum%2
	if mod != 0 {
		return false
	}
	targets := []int{halfSum, halfSum}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var search func(index int) bool
	search = func(index int) bool {
		if index == len(nums) {
			return true
		}
		for i, v := range targets {
			if v >= nums[index] {
				targets[i] -= nums[index]
				if search(index + 1) {
					return true
				}
				targets[i] += nums[index]
			}
		}
		return false
	}
	return search(0)
}
