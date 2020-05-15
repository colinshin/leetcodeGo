/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package subset

import (
	"strconv"
)

/*
78. 子集 https://leetcode-cn.com/problems/subsets/
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

/*
朴素实现
时空复杂度均为O(n*2^n)
*/
func subsets(nums []int) [][]int {
	result := [][]int{{}} // 空集也是子集之一
	for _, num := range nums {
		for _, r := range result {
			tmp := make([]int, len(r)+1)
			_ = copy(tmp, r)
			tmp[len(tmp)-1] = num
			result = append(result, tmp)
		}
	}
	return result
}

/*
回溯
时空复杂度均为O(n*2^n)
*/
func subsets1(nums []int) [][]int {
	var result [][]int
	var curr []int
	var backtrack func(start, size int)
	backtrack = func(start, size int) {
		if size == len(curr) {
			tmp := make([]int, len(curr))
			_ = copy(tmp, curr)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i+1, size)
			curr = curr[:len(curr)-1]
		}
	}
	for size := 0; size <= len(nums); size++ {
		backtrack(0, size)
	}
	return result
}

/*
数学法：字典排序（二进制排序）
nums里的每个元素，要么在结果中，要么不在结果中
用一个n位的bitset来表示各个元素在不在结果中，
如000...000表示所有元素都不在结果中，000..011表示后边两个元素在结果中
借助标准库，也可以把bitset改成一个二进制的字符串，称为mask，
需要注意若干0开头的mask不能把0丢弃，可以给所有mask最前边补个1来解决
时空复杂度均为O(n*2^n)
*/
func subsets3(nums []int) [][]int {
	var result [][]int
	dummy := 1 << uint(len(nums))
	end := dummy
	for i := 0; i < end; i++ {
		mask := strconv.FormatUint(uint64(i|dummy), 2)[1:]
		var curr []int
		for j := range mask {
			if mask[j] == '1' {
				curr = append(curr, nums[j])
			}
		}
		result = append(result, curr)
	}
	return result
}
