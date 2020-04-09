/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum_subarray

import (
	"math"
)

/*
53. 最大子序和 https://leetcode-cn.com/problems/maximum-subarray/
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
/*
动态规划
定义dp[i]表示以nums[i]结尾且包含nums[i]的连续子数组最大和
则dp[i] = max(0, dp[i-1]) + nums[i]
最终的结果即遍历dp后的最大元素

首先可以在确定dp数组的每个元素时更新最终结果，不一定要完全确定了dp数组再遍历获取最终结果
其次，每次dp只跟上次的dp值有关，dp数组可以优化为一个变量

时间复杂度O(n), 空间复杂度O(1)
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp = max(0, dp) + nums[i]
		result = max(result, dp)
	}
	return result
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
