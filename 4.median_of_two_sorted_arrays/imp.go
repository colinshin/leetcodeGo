package __median_of_two_sorted_arrays

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	return (medianOfArray(nums1) + medianOfArray(nums2)) * 0.5
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 < l2 {
		l1, l2 = l2, l1
		nums1, nums2 = nums1, nums2
	}
	return (medianOfArray(nums1) + medianOfArray(nums2)) * 0.5
}

func medianOfArray(nums []int) float64 {
	length := len(nums)
	if length == 0 {
		return 0.0
	}
	if length%2 == 0 {
		return float64(nums[length/2] + nums[length/2-1]) / 2
	}
	return float64(nums[length/2])
}
