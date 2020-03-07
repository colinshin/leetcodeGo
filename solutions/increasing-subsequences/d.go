/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package increasing_subsequences

/*
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:

给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
递归与回溯

查找递增数组实际就是以每个元素为起点，后面每个比它大的元素都可以发散出一条路径，所以可以用dfs。
数组中重复的元素只需要以前面的元素为起点进行dfs，因为后面元素的所有情况在前面都可以考虑到。
所以对需要dfs的nums元素用set去重；
注意在以i为起点dfs时，还要对后面的元素相互去重
*/
func findSubsequences(nums []int) [][]int {
	var cur []int
	var result [][]int

	var dfs func(int)
	dfs = func(start int) {
		cur = append(cur, nums[start])
		if len(cur) >= 2 {
			b := make([]int, len(cur))
			copy(b, cur)
			result = append(result, b)
		}
		visited := make(map[int]struct{}, 0)
		for i := start + 1; i < len(nums); i++ {
			if nums[start] <= nums[i] {
				if _, ok := visited[nums[i]]; ok {
					continue
				}
				visited[nums[i]] = struct{}{}
				dfs(i)
			}
		}
		cur = cur[:len(cur)-1] // 回溯
	}

	visited := make(map[int]struct{}, 0)
	for i := 0; i < len(nums)-1; i++ { // 递增子序列长度至少为2， 没有必要以最后一个元素为起点去dfs遍历了
		// 数组中可能有重复元素，后边的重复元素dfs遍历情况包含在前边的重复元素dfs遍历里边，不用再遍历
		if _, ok := visited[nums[i]]; ok {
			continue
		}
		visited[nums[i]] = struct{}{}
		dfs(i) // 以i为起点深度优先遍历
	}
	return result
}
