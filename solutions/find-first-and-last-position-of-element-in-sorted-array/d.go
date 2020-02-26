/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_first_and_last_position_of_element_in_sorted_array

/*	在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func searchRange(nums []int, target int) []int {
	leftIndex := search(nums, target) // 或用标准库api： leftIndex := sort.SearchInts(nums, target)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}
	return []int{leftIndex, searchFromRight(nums, target) - 1}
}

/*
在nums里搜索target，返回新target应该插入的位置；如果nums里已经有target，则在第一个已有target元素之前插入
nums已经排序，但可能有重复元素

功能同标准库里的sort.SearchInts(nums, target)
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] >= target:
			right = mid
		}
	}
	return left
}

/*
在nums里从右向左搜索target，返回新target应该插入但位置；如果nums里已经有target，则在最后一个target元素之后插入
nums已经排序，但可能有重复元素
*/
func searchFromRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] <= target:
			left = mid + 1
		case nums[mid] > target:
			right = mid
		}
	}
	return left
}
