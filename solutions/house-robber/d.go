/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package house_robber

import "math"

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
假设f(i)表示有i+1个房子的时候偷到的最大金额(为了和数组索引对应)
那么f(0) = nums[0]
f(1) = max(nums[0], nums[1]
对于一个大于1的i，可以分为两种情况：
偷i房间，收益尽可能大的话就是nums[i]+f(i-2), 不偷i房间，则f(i)=f(i-1);所以f(i) = max(f(i-1), nums[i]+f(i-2))
这就是一个典型的动态规划
时间复杂度O(n)
空间复杂度O(n), 主要为引入dp数组开辟但额外空间
*/
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// 但实际上只需要两个遍变量，不需要一个dp数组
// 这样空间复杂度为O(1)
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// f(i) = max(f(i-1), nums[i]+f(i-2))
	prev, curr := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		prev, curr = curr, max(curr, nums[i]+prev)
	}
	return curr
}

//另外实际上我们发现，一开始让prev和curr均为0，从头遍历，逻辑同样正确，代码可以简化
func rob(nums []int) int {
	// f(i) = max(f(i-1), nums[i]+f(i-2))
	prev, curr := 0, 0
	for _, v := range nums {
		prev, curr = curr, max(curr, prev+v)
	}
	return curr
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
