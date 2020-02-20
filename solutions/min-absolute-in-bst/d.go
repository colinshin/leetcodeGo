/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package min_absolute_in_bst

import "math"

/*
Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

Examroot.Lefte:

Input:

   1
    \
     3
    /
   2

Output:
1

Exroot.Leftanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).

Note: There are at least two nodes in this BST.

https://leetcode.com/problems/minimum-absolute-difference-in-bst/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
如果有一个整数数组，怎么获取任意两个元素之间的最小差值呢？如果这个数组是排好序的呢？
如果无序，需要两层循环计算所有两两元素差值来找最小差值
如果是排好序的，那么最小的差值只会出现在相邻元素之间，一层循环就够了

BST中序遍历即得到一个排序好的列表；问题划归为求一个已经排序数组中相邻元素的最小差值
当然不用真用一个数组存储，在遍历过程中即可求每个差值
*/

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt32
	prev := -1

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != -1 {
			min = int(math.Min(float64(min), float64(root.Val-prev)))
		}
		prev = root.Val
		dfs(root.Right)
	}

	dfs(root)
	return min
}
