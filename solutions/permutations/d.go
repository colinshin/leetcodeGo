/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package permutations

func permute(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}
	var result [][]int
	for _, v := range permute(nums[:len(nums)-1]) {
		for i := 0; i <= len(v); i++ {
			t := append(append(v[:i:i], nums[len(nums)-1]), v[i:]...)
			result = append(result, t)
		}
	}
	return result
}

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	// 保持start之前的元素固定不变，将其及其之后的元素全排列
	var dfs func(int)
	dfs = func(start int) {
		if start == n-1 {
			r := make([]int, n)
			_ = copy(r, nums)
			result = append(result, r)
			return
		}
		visited := make(map[int]bool, n-start)
		for i := start; i < n; i++ { // 将start及其之后的元素全排列，注意不能漏了start
			if visited[nums[i]] {
				continue
			}
			visited[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start] // 通过交换做排列
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return result
}
