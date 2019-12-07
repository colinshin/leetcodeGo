# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## 题目
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

长度分别为m和n的两个已排序数组，求其中位数，时间复杂度控制在O(log (m+n))内

Example 1:
```
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
```
Example 2:
```
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
```
## 分析
对于一个有序数组，如果元素个数为奇数，中位数即中间元素的值；若元素个数为偶数，中位数为中间两个元素的平均值。<br>
对于两个或多个有序数组，其合并后的中位数并非每个数组中位数的平均值，如：
```
[1, 3, 5] // 中位数3
[8, 10] // 中位数9
// 合并后的数组
[1, 3, 5, 8, 10] // 中位数5, 并非3和9的平均数
```
所以，必须对两个数组合并，合并后依然有序<br>

1. 朴素实现（时间与空间复杂度均为O(m+n)）

```
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedNums := mergeNums(nums1, nums2)
	return medianOfArray(mergedNums)
}

func mergeNums(nums1, nums2 []int) []int {
	var r []int
	m := len(nums1)
	n := len(nums2)
	for i, j, k := 0, 0, 0; k < m+n; k++ {
		if nums1[i] < nums2[j] {
			r = append(r, nums1[i])
			if i < m-1 {
				i ++
			}
		} else {
			r = append(r, nums2[j])
			if j < n-1 {
				j ++
			}
		}
	}
	return r
}

func medianOfArray(nums []int) float64 {
	length := len(nums)
	if length == 0 {
		return 0.0
	}
	if length%2 == 0 {
		return float64(nums[length/2]+nums[length/2-1]) / 2
	}
	return float64(nums[length/2])
}
```

2. 时间O(log(m+n))，空间O(1)<br>
原理参考：<br>
https://cloud.tencent.com/developer/article/1483811<br>
代码略<br>

3.时间O(log(min(m,n)))，空间O(1)<br>
原理参考：<br>
https://blog.csdn.net/bjweimengshu/article/details/97717144<br>

```
import "github.com/zrcoder/leetcodeGo/util/integer"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = integer.Max(nums1[i-1], nums2[j-1])
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
				minRight = integer.Min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}
```
