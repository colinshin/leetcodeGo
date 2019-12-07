package __median_of_two_sorted_arrays

import "github.com/zrcoder/leetcodeGo/util/integer"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	middle := (n + m + 1) / 2

	middleVal := getKth(nums1, nums2, 0, n-1, 0, m-1, middle)
	if n+m%2 == 1 {
		return middleVal
	}

	nextMiddle := (n + m + 2) / 2
	nextMiddleVal := getKth(nums1, nums2, 0, n-1, 0, m-1, nextMiddle)
	return (middleVal + nextMiddleVal) * 0.5
}

func getKth(nums1, nums2 []int, start1, end1, start2, end2, k int) float64 {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 > len2 {
		return getKth(nums2, nums1, start2, end2, start1, end1, k)
	}
	if len1 == 0 {
		return float64(nums2[start2+k-1])
	}
	if k == 1 {
		return float64(integer.Min(nums1[start1], nums2[start2]))
	}

	i := start1 + integer.Min(len1, k/2) - 1
	j := start2 + integer.Min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2, start1, end1, j+1, end2, k-(j-start2+1))
	}

	return getKth(nums1, nums2, i+1, end1, start2, end2, k-(i-start1+1))
}

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
