/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package binary_tree_inorder_traversal_

/*
94. 二叉树的中序遍历 https://leetcode-cn.com/problems/binary-tree-inorder-traversal

给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现，时空复杂度O(n), n为节点总数
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// left -> root -> right
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

// 迭代实现，时空复杂度O(n)
func inorderTraversal1(root *TreeNode) []int {
	var (
		result []int
		stack  []*TreeNode
	)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

/* 迭代，节点标记法
在出入栈的时候，标记节点，具体为：
标记节点的状态，新节点为false，已使用（在这道题里是指将节点值追加到结果数组）的节点true。
如果遇到未标记的节点，则将其标记为true，然后将其右节点、自身、左节点依次入栈；注意到顺序与遍历次序正好相反。
如果遇到的节点标记为true，则使用该节点。

这个方法在前序、中序、后续遍历里的实现代码总体逻辑一致，只是入栈的顺序稍微调整即可
*/
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{root}
	marked := make(map[*TreeNode]bool, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		if marked[node] {
			result = append(result, node.Val)
			continue
		}
		marked[node] = true
		stack = append(stack, node.Right)
		stack = append(stack, node)
		stack = append(stack, node.Left)
	}
	return result
}
