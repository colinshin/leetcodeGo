package check_if_a_number_is_majority_element_in_a_sorted_array

/*
给出一个按 非递减 顺序排列的数组 nums，和一个目标数值 target。假如数组 nums 中绝大多数元素的数值都等于 target，则返回 True，否则请返回 False。
所谓占绝大多数，是指在长度为 N 的数组中出现必须 超过 N/2 次。

示例 1：

输入：nums = [2,4,5,5,5,5,5,6,6], target = 5
输出：true
解释：
数字 5 出现了 5 次，而数组的长度为 9。
所以，5 在数组中占绝大多数，因为 5 次 > 9/2。
示例 2：

输入：nums = [10,100,101,101], target = 101
输出：false
解释：
数字 101 出现了 2 次，而数组的长度是 4。
所以，101 不是 数组占绝大多数的元素，因为 2 次 = 4/2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
方法1：
用两次二分法找到数组中最左边target的索引和最右边的索引，两个索引距离+1即target出现的个数
时间复杂度O(lgn)，空间复杂度O(1)
*/
func isMajorityElement(nums []int, target int) bool {
	left := searchLeft(nums, target)
	if left == -1 {
		return false
	}
	right := searchRight(nums, target)
	return right-left+1 > len(nums)/2
}

// 二分法找到nums里最左边的target元素的索引
func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}

// 二分法找到nums里最右边的target的索引
func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}

/* searchLeft 可直接用标准库：
func searchLeft(nums []int, target int) int {
	i := sort.SearchInts(nums, target)
	if i == len(nums) {
		return -1
	}
	return i
}
*/

// 以下方法都没有借助数组已经排序的特性

/* 方法2：借助一个map，key为每个元素，值为对应元素出现的个数
时间空间复杂度都是O(n)
*/
func isMajorityElement1(nums []int, target int) bool {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	return m[target] > len(nums)/2
}

/* 方法3：直接遍历数组得到个数
时间复杂度O(n)，空间复杂度O(1)
*/
func isMajorityElement2(nums []int, target int) bool {
	count := 0
	for _, v := range nums {
		if v == target {
			count++
		}
	}
	return count > len(nums)/2
}
