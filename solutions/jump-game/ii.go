/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package jump_game

import "math"

/*
自顶向下动态规划，或者理解为带备忘的回溯
时间复杂度O(n^2),空间复杂度O(n)
*/
func jump1(nums []int) int {
	n := len(nums)
	memo := make([]int, n) // memo[i]表示从位置i跳到最后位置的最小步数
	memo[n-1] = 0
	for i := 0; i < n-1; i++ {
		memo[i] = n // n相当于max
	}
	var helper func(pos int) int
	helper = func(pos int) int {
		if memo[pos] < n {
			return memo[pos]
		}
		end := min(pos+nums[pos], n-1)
		for i := end; i > pos; i-- {
			memo[pos] = min(memo[pos], helper(i)+1)
		}
		return memo[pos]
	}
	return helper(0)
}

/*
自底向上动态规划
时间复杂度O(n^2),空间复杂度O(n)
*/
func jump2(nums []int) int {
	n := len(nums)
	dp := make([]int, n) // dp[i]表示从位置i跳到最后位置的最小步数
	dp[n-1] = 0
	for i := 0; i < n-1; i++ {
		dp[i] = n // n相当于max
	}
	for i := n - 2; i >= 0; i-- {
		end := min(nums[i]+i, n-1)
		for j := end; j > i; j-- {
			dp[i] = min(dp[i], dp[j]+1)
		}
	}
	return dp[0]
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

/*
逆向思考。要到达最后一个位置，前一个位置在哪？找到后，再继续寻找上上个位置，直到找到第0个位置；
为了使最终步数最少，每次需要找到距离当前位置最远的距离，从左到右遍历数组，第一个满足的位置就是了。
时间复杂度， 最坏情况下是O(n^2), 空间复杂度O(1)
*/
func jump3(nums []int) int {
	pos := len(nums) - 1
	result := 0
	for pos > 0 {
		i := 0
		for i < pos && i+nums[i] < pos { // 从左到右找到第一个能跳到pos的位置i即为最优的i
			i++
		}
		pos = i
		result++
	}
	return result
}

/*
贪心策略
每次在可跳范围内选择可以跳得更远的位置
遍历时对i+nums[i]使用贪心策略做选择
时间复杂度O(n), 空间复杂度O(1)
*/
func jump(nums []int) int {
	var (
		end      = 0 // 代表当前可以调到的边界
		farthest = 0 // 代表下一跳可以跳到的最远位置
		result   = 0
	)
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end {
			end = farthest
			result++
		}
	}
	return result
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
