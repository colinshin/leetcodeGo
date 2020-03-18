/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package median_of_two_sorted_arrays

import (
	"math"
)

/*
寻找两个已排序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 朴素实现，先merge再找中间的
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	return medianOf(merge(nums1, nums2))
}

func merge(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	r := make([]int, m+n)
	for i, j, k := 0, 0, 0; i < m || j < n; k++ {
		if j == n {
			r[k] = nums1[i]
			i++
		} else if i == m {
			r[k] = nums2[j]
			j++
		} else if nums1[i] < nums2[j] {
			r[k] = nums1[i]
			i++
		} else {
			r[k] = nums2[j]
			j++
		}
	}
	return r
}

func medianOf(nums []int) float64 {
	length := len(nums)
	if length == 0 {
		return 0.0
	}
	if length%2 == 0 {
		return float64(nums[length/2]+nums[length/2-1]) * 0.5
	}
	return float64(nums[length/2])
}

// 朴素实现的改进，不用真的merge
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	lastR, currentR := -1, -1
	start1, start2 := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		lastR = currentR
		if start1 < m && (start2 == n || nums1[start1] <= nums2[start2]) {
			currentR = nums1[start1]
			start1++
		} else {
			currentR = nums2[start2]
			start2++
		}
	}
	if (m+n)%2 == 1 {
		return float64(currentR)
	}
	return float64(lastR+currentR) * 0.5
}

// 划归，寻找第K个数。参考：
// https://cloud.tencent.com/developer/article/1483811
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	if size == 0 {
		return 0.0
	}
	if size%2 == 1 {
		return getKth(nums1, nums2, size/2+1)
	}
	return (getKth(nums1, nums2, size/2) + getKth(nums1, nums2, size/2+1)) * 0.5
}
func getKth(nums1, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return getKth(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	i, j := min(m-1, k/2-1), min(n-1, k/2-1)
	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2[j+1:], k-(j+1))
	}
	return getKth(nums1[i+1:], nums2, k-(i+1))
}

// 参考
// https://blog.csdn.net/bjweimengshu/article/details/97717144
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays1(nums2, nums1) // prevent edge conditions
	}

	left, right := 0, m
	for left <= right {
		i := left + (right-left)/2
		j := (m+n+1)/2 - i
		if i < m && j > 0 && nums2[j-1] > nums1[i] { // i is too smal
			left = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] { // i is too big
			right = i - 1
		} else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
