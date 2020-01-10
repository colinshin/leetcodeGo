# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

给定一个整型数组，返回和为定值（target）的两个元素的索引。

假定每个输入有且仅有一组解， 且同一个元素不会出现两次。

Example:

```text
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 分析

1. 朴素实现， O(n^2)复杂度<br>
```
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j ++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```
2. 时间O(n), 空间O(n)的实现<br>
在遍历中，如果i在结果中，target-nums[i]必在nums中，否则i不满足要求<br>
为了迅速查找target-nums[i]是否在nums中，我们可以构造一个map，其键为nums里的元素，值为元素的索引<br>
构造需要遍历一遍nums，注意构造过程中，考虑有nums元素重复的情况<br>
查找另需一遍遍历<br>
```
func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, element := range nums {
		index[element] = i
	}
	for i, element := range nums {
		if j, found := index[target-element]; found {
			return []int{i, j}
		}
	}
	return nil
}
```
3. 进一步优化<br>
实际上可以边构建map，边做查找，总体只需遍历一遍<br>
```
func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, element := range nums {
		if j, found := index[target-element]; found {
			return []int{j, i} // not {i, j}, but {j, i}; let's think, j < i
		}
		index[element] = i
	}
	return nil
}

```
## 拓展
如果数组是已经排好序的呢？<br>
可以从两边往中间凑， 时间O(n), 不用额外空间~
```
func twoSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		}
		if sum < target {
			i ++
		} else {
			j --
		}
	}
	return nil
}
```
