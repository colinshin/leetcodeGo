/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package jump_game

// 贪心，也是最朴素的实现方法；时间复杂度O(n), 空间复杂度O(1)
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		if i > farthest {
			return false
		}
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
	}
	return farthest >= len(nums)-1
}

/*
自底向上动态规划
时间复杂度O(n^2),空间复杂度O(n)
*/
func canJump1(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		end := i + nums[i]
		if end > len(nums)-1 { // 防止索引越界
			end = len(nums) - 1
		}
		for j := end; j > i; j-- { // j从左向右遍历也行
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

/*
自顶向下动态规划，或者理解为带备忘的回溯
时间复杂度O(n^2),空间复杂度O(2n)=O(n)，第一个n是栈空间开销，第二个是dp数组开销
*/
func canJump2(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	const ok, nok = 1, 2 // unknown = 0
	memo := make([]int, len(nums))
	memo[len(nums)-1] = ok
	var canJumpFrom func(pos int) bool
	canJumpFrom = func(pos int) bool {
		if memo[pos] == ok {
			return true
		}
		if memo[pos] == nok {
			return false
		}
		end := pos + nums[pos]
		if end > len(nums)-1 { // 防止索引越界
			end = len(nums) - 1
		}
		for i := pos + 1; i <= end; i++ {
			if canJumpFrom(i) {
				memo[pos] = ok
				return true
			}
		}
		memo[pos] = nok
		return false
	}
	return canJumpFrom(0)
}
