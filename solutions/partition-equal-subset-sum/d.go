/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package partition_equal_subset_sum

import "sort"

/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
先求所有数字和，如果是奇数直接返回false
如果是偶数，问题转化为：找出一部分数字，其和为sum/2

这是一个01背包问题:
原数组里面的每个元素都可以看作是一种物品，而这件物品的重量和价值都为元素值；
原数组和的一半可看作背包的最大承重量，而当背包能放下物品的最大价值为原数组和的一半时，就返回真，否则返回假
*/
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
