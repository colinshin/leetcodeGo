/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum_depth_of_binary_tree

import "math"

/*
104. 二叉树的最大深度 https://leetcode-cn.com/problems/maximum-depth-of-binary-tree

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 时间复杂度O(n): 每个节点遍历一次 —— n为节点数量
// 空间复杂度，当树完全不平衡，退化为链表，最坏为O(n)；当树平衡时，为O(lgn)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

/*
变体： 如果求二叉树的最大直径呢？

543. 二叉树的直径 https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

总节点数 = 左子树最大深度 + 右子树最大深度 + 1（node本身）
直径需要-1，故直径 = leftDepth + rightDepth
*/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := maxDepth(root.Left) + maxDepth(root.Right)
	sum = max(diameterOfBinaryTree(root.Left), sum)
	sum = max(diameterOfBinaryTree(root.Right), sum)
	return sum
}

// 可以进一步优化，在求最大深度的过程中更新结果
func diameterOfBinaryTree1(root *TreeNode) int {
	result := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int { // 计算node的最大深度
		if node == nil {
			return 0
		}
		leftDepth, rightDepth := depth(node.Left), depth(node.Right)
		result = max(result, leftDepth+rightDepth)
		return 1 + max(leftDepth, rightDepth)
	}
	_ = depth(root)
	return result
}
