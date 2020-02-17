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

#### 0. 朴素实现（时间与空间复杂度均为O(m+n)）

```
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
```
#### 1. 朴素实现的改进，不用真的merge
```
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	lastR, currentR := -1, -1
	start1, start2 := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		lastR = currentR
		if start1 < m && (start2 >= n || nums1[start1] < nums2[start2]) {
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
```
#### 2. 时间O(log(m+n))，空间O(1)

原理参考：<br>
https://cloud.tencent.com/developer/article/1483811<br>

```
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	if (n+m)%2 == 1 {
		return getKth(nums1, nums2, (n+m+1)/2)
	}
	return (getKth(nums1, nums2, (n+m)/2) + getKth(nums1, nums2, (n+m)/2+1)) * 0.5
}

func getKth(nums1, nums2 []int, k int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return getKth(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(integer.Min(nums1[0], nums2[0]))
	}

	i := integer.Min(m, k/2) - 1
	j := integer.Min(n, k/2) - 1
	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2[j+1:], k-(j+1))
	}
	return getKth(nums1[i+1:], nums2, k-(i+1))
}
```

#### 3. 时间O(log(min(m,n)))，空间O(1)

原理参考：<br>
https://blog.csdn.net/bjweimengshu/article/details/97717144<br>

```
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}

	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] {
			// i is too smal
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i is too big
			iMax = i - 1
		} else {
			// i is perfect
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
