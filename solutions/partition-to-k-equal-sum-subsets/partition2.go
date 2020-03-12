/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package partition_to_k_equal_sum_subsets

import "sort"

// 100 / 105 个通过测试用例， 在101用例超时
func canPartition2(nums []int) bool {
	const groups = 2
	if len(nums) < groups {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / groups
	if sum%groups != 0 || max > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking(groups, 0, 0, target, nums, used)
}

// 改进canPartition2，先对nums从大到小排序，极大降低递归次数
// 如果不能修改原数组，可以深拷贝一份
func canPartition(nums []int) bool {
	const groups = 2
	if len(nums) < groups {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum / groups
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if sum%groups != 0 || nums[0] > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking(groups, 0, 0, target, nums, used)
}

// 01 package problem
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
