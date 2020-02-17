package merge_sorted_array

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
如果从前往后遍历，nums1还未处理的元素可能被覆盖，为解决这个问题需要额外m的空间把nums1的元素保存起来

如果从后往前遍历，则规避了覆盖的问题
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for i >= 0 { // j == -1
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 { // i == -1
		nums1[k] = nums2[j]
		j--
		k--
	}
}
