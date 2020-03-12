/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_k_th_smallest_pair_distance

import "sort"

/*
给定一个整数数组，返回所有数对之间的第 k 个最小距离。
一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。

示例 1:
输入：
nums = [1,3,1]
k = 1
输出：0
解释：
所有数对如下：
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
因此第 1 个最小距离的数对是 (1,1)，它们之间的距离为 0。

提示:
2 <= len(nums) <= 10000.
0 <= nums[i] < 1000000.
1 <= k <= len(nums) * (len(nums) - 1) / 2.


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
排序后方便计算最小距离，二分搜索
*/
func smallestDistancePair(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := nums[len(nums)-1] - nums[0] // 数对距离最大为max
	left, right := 0, max+1
	for left < right {
		mid := left + (right-left)/2
		if countLowers(nums, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 使用标准库，减少代码量
func smallestDistancePair1(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := nums[len(nums)-1] - nums[0] // 数对距离最大为max
	return sort.Search(max+1, func(i int) bool {
		return countLowers(nums, i) >= k
	})
}

// 返回nums中数对距离不大于v的个数；nums已经排序
func countLowers(nums []int, v int) int {
	count := 0
	for left, right := 0, 0; right < len(nums); right++ {
		for nums[right]-nums[left] > v {
			left++
		}
		count += right - left
	}
	return count
}
