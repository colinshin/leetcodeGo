/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_increasing_subsequence

import "math"

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n^2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

通过次数49,165提交次数111,640

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
动态规划，时间复杂度O(n^2),  空间复杂度O(n)
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)) // dp[i]代表nums[0:i]的最长子序列长度
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 一个元素算递增长度为1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

/*
二分法; 时间复杂度O(nlgn), 空间复杂度O(n)

这是一个纸牌游戏～～
这里有个解释：https://www.itcodemonkey.com/article/15644.html
*/
func lengthOfLIS1(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for _, v := range nums {
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] < v {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles++
		}
		top[left] = v
	}
	return piles
}
