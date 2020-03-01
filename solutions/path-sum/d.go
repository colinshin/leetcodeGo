/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package path_sum

import (
	"bytes"
	"sort"
	"strconv"
)

/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 常规dfs
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

/* 变体 https://leetcode-cn.com/problems/path-sum-ii

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/
/*
用一个切片path记录遍历的路径，到达叶子节点发现path内元素和为sum则将当期path添加到结果里，
注意切片底层是同一个数组，添加到结果时要深拷贝一份
*/
func prefixSum(root *TreeNode, sum int) [][]int {
	var count [][]int
	var path []int
	prefixSum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		prefixSum += node.Val
		if node.Left == nil && node.Right == nil {
			if prefixSum == sum {
				tmp := make([]int, len(path))
				_ = copy(tmp, path)
				count = append(count, tmp)
			}
		}
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
		prefixSum -= node.Val
	}
	dfs(root)
	return count
}

/*变体，类似前缀树Trie的实现，用数组表示树，且树是多叉树，应该怎么解？
假设有k个节点，每个节点从0到k-1编号，编号即为其id
caps数组表示每个节点的值
哈希表relations，是个邻接表，键为节点id，值为节点的孩子节点组成的数组
给定sum，返回每条从根节点（id为0）出发到叶子节点，值相加和为sum的路径组成的集合
路径处理成字符串，前最终结果按照字符串非递增排序
*/
func getPath(caps []int, relations map[int][]int, sum int) []string {
	var count []string
	var path []int
	prefixSum := 0
	var dfs func(nodeId int)

	dfs = func(nodeId int) {
		path = append(path, caps[nodeId])
		prefixSum += caps[nodeId]
		if len(relations[nodeId]) == 0 {
			if prefixSum == sum {
				count = append(count, parsePath(path))
			}
		}
		for _, c := range relations[nodeId] {
			dfs(c)
		}
		path = path[:len(path)-1]
		prefixSum -= caps[nodeId]
	}
	dfs(0)

	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})
	return count
}
func parsePath(path []int) string {
	buf := bytes.NewBuffer(nil)
	for _, v := range path {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString(" ")
	}
	count := buf.String()
	return count[:len(count)-1]
}

/* 变体 假设不一定要从根节点开始，也不需要走到叶子节点，来查找和为定值的路径呢？
https://leetcode-cn.com/problems/path-sum-iii

给定一个二叉树，它的每个结点都存放着一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/
// 递归解法
func pathSumCount(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := countPrefix(root, sum)
	count += pathSumCount(root.Left, sum)
	count += pathSumCount(root.Right, sum)
	return count
}

// 返回前缀和为sum的路径个数
func countPrefix(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val == sum {
		count = 1
	}
	count += countPrefix(root.Left, sum-root.Val)
	count += countPrefix(root.Right, sum-root.Val)
	return count
}

/*
如果某个节点x的前缀和等于其某个子孙节点y的前缀和减去sum，
即prefixSum(x) = prefixSum(y)-sum ，说明x到y这条路径的和是sum
*/
func pathSumCount0(root *TreeNode, sum int) int {
	prefixSumCount := make(map[int]int, 0) // 记录前缀和，key为前缀和，value为前缀和的个数
	prefixSumCount[0] = 1                  // 前缀和为0的一条路径，方便边界处理，即节点值就是sum这种情况
	count := 0
	prefixSum := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		prefixSum += node.Val                  // 当前节点node的前缀和（即从root到当前节点这条路径的和）
		count += prefixSumCount[prefixSum-sum] // 如果当前节点之前已经有前缀和为urrPrefixSum-sum的节点，说明那些节点到当前节点的和就是sum
		prefixSumCount[prefixSum]++
		dfs(node.Left)
		dfs(node.Right)
		// 回溯
		prefixSumCount[prefixSum]--
		prefixSum -= node.Val
	}

	dfs(root)
	return count
}
