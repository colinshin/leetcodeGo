/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package partition_to_k_equal_sum_subsets

import "sort"

// 100 / 105 个通过测试用例， 在101用例超时
func canPartition0(nums []int) bool {
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

func canPartition00(nums []int) bool {
	const groups = 2
	if len(nums) < groups {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%groups != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return backtrack(0, sum/groups, nums)
}

func backtrack(index int, remain int, nums []int) bool {
	if index < 0 || nums[index] > remain {
		return false
	}
	if nums[index] == remain {
		return true
	}
	for i := 1; i >= 0; i-- {
		remain -= nums[index] * i
		if backtrack(index-1, remain, nums) {
			return true
		}
		remain += nums[index] * i
	}
	return false
}

/*
其实这是一个01背包问题：

给一个可装载重量为 sum / 2 的背包和 N 个物品，
每个物品的重量为 nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满？

时间复杂度O(n * c),其中c背包容量，即所有元素和的一半
空间复杂度O((n+1)*(c+1)) = O(n * c)，dp数组的大小；通过状态压缩，可以将dp压缩为一维数组，空间复杂度降为O(c)
*/
func canPartition1(nums []int) bool {
	n := len(nums)
	if n < 2 {
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
	dp := make([][]bool, n+1) // dp[i][j] 表示对于前 i 个物品，背包的容量为 j 时是否恰好能把背包装满
	for i := range dp {
		dp[i] = make([]bool, c+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = true // 容量为0，相当于装满了
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			if j-nums[i-1] < 0 { // 容量不足，装不了物品i
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]] // 不装或者装
			}
		}
	}
	return dp[n][c]
}

/*
dp空间压缩
通过状态转移方程，可以看到dp[i]只和dp[i-1]有关，可以优化为c+1大小的一维数组
*/
func canPartition2(nums []int) bool {
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

	c := sum / 2
	dp := make([]bool, c+1)
	dp[0] = true // 容量为0相当于装满
	for i := 1; i < len(nums); i++ {
		for j := c; j >= nums[i]; j-- { // j 应该从后往前反向遍历，以免之前的结果当前结果
			dp[j] = dp[j] || dp[j-nums[i]] // 不装或装
		}
	}
	return dp[c]
}
