## [46. 全排列](https://leetcode-cn.com/problems/permutations)
```text
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
   [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```
递归，在已有n-1大小的排列的每个空隙插入最后一个元素
```go
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
```
深度优先搜索，有一个自然的递归，先固定前边几个元素，然后开始尝试排列后边的。参见dfs函数:
```go
func permute(nums []int) [][]int {
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
		for i := start; i < n; i++ { // 将i及其i之后的元素全排列，注意不能漏了i
			nums[start], nums[i] = nums[i], nums[start] // 通过交换做排列
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return result
}
```
## [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii)
```text
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```
问题与46相似，只是加了元素可能重复的情况，结果不能有重复；<br>
解法与46的解法也相似，递归时用set去重,<br>
具体在交换start处元素与后边元素的时候，看看是否已有相同的元素参与过排列，已经参与过的跳过<br>
[参考实现](d.go)