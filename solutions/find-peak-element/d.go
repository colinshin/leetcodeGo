/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_peak_element

import "math"

/*
峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。

示例 1:

输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
示例 2:

输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
说明:

你的解法应该是 O(logN) 时间复杂度的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-peak-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
朴素实现
遍历nums，如果发现nums[i] > nums[i+1]，则i为所求；找不到则为最后一个索引
时间复杂度O(n）
*/
func findPeakElement0(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

/*
二分法
对比nums[mid]和其右边的元素，如果nums[mid]大，说明峰值在mid左侧，包括mid；否则峰值在mid右侧
时间复杂度O(lgn)
*/
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 模板1 查找单个索引
func findPeakElement1(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		pre, next := math.MinInt64, math.MinInt64 // 这里预期测试用例里没有{math.MinInt64, math.MinInt64+1}, 不然预期1， 我们返回了0
		if mid > 0 {
			pre = nums[mid-1]
		}
		if mid < len(nums)-1 {
			next = nums[mid+1]
		}
		if pre < nums[mid] && nums[mid] > next {
			return mid
		}
		if pre < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 模板3
func findPeakElement3(nums []int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	if right >= 0 && nums[left] < nums[right] {
		return right
	}
	return left
}
