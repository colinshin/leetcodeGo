/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package binary_tree_postorder_traversal

/*
145. 二叉树的后序遍历 https://leetcode-cn.com/problems/binary-tree-postorder-traversal

给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解决
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var postororder func(root *TreeNode)
	postororder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// left -> right -> root
		postororder(node.Left)
		postororder(node.Right)
		result = append(result, node.Val)
	}
	postororder(root)
	return result
}

/* 迭代，节点标记法
在出入栈的时候，标记节点，具体为：
标记节点的状态，新节点为false，已使用（在这道题里是指将节点值追加到结果数组）的节点true。
如果遇到未标记的节点，则将其标记为true，然后将其自身、右节点、左节点依次入栈；注意到顺序与遍历次序正好相反。
如果遇到的节点标记为true，则使用该节点。

这个方法在前序、中序、后续遍历里的实现代码总体逻辑一致，只是入栈的顺序稍微调整即可
*/
func postorderTraversal2(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{root}
	marked := map[*TreeNode]bool{}
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
		stack = append(stack, node)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return result
}
