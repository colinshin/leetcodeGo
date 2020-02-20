/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_minimum_in_rotated_sorted_array

import "math"

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
二分法，每次将mid和right处的值比较，以判断mid落在来旋转点左侧还是右侧
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] { // mid落在旋转点左侧
			left = mid + 1
		} else { // mid和right在旋转点同侧，但因为一开始right在整个数组最右，所以当前只可能同在旋转点右侧
			right = mid
		}
	}
	return nums[left]
}

/*
题目变体：
如果nums里有重复元素呢？如：
[3, 1, 2, 3]
[3, 1, 2, 2, 3]
*/
func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] > nums[right]: // mid落在旋转点左侧
			left = mid + 1
		case nums[mid] < nums[right]: // mid和right在旋转点同侧，但因为一开始right在整个数组最右，所以当前只可能同在旋转点右侧
			right = mid
		default: // TODO: what's this?
			right--
		}
	}
	return nums[left]
}

// 朴素实现
func findMin2(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt32
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}
