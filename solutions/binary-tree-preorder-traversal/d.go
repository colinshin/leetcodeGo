package binary_tree_preorder_traversal

/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]

进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现，时空复杂度O(n), n为节点总数
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

// 迭代实现，时空复杂度O(n), n为节点总数
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

/* 迭代，节点标记法
在出入栈的时候，标记节点，具体为：
标记节点的状态，新节点为false，已使用（在这道题里是指将节点值追加到结果数组）的节点true。
如果遇到未标记的节点，则将其标记为true，然后将其右节点、左节点、自身依次入栈；注意到顺序与遍历次序正好相反。
如果遇到的节点标记为true，则使用该节点。

这个方法在前序、中序、后续遍历里的实现代码总体逻辑一致，只是入栈的顺序稍微调整即可
*/
func preorderTraversal2(root *TreeNode) []int {
	type item struct {
		node   *TreeNode
		marked bool
	}
	var result []int
	stack := []*item{{node: root}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.node == nil {
			continue
		}
		if top.marked {
			result = append(result, top.node.Val)
			continue
		}
		top.marked = true
		stack = append(stack, &item{node: top.node.Right})
		stack = append(stack, &item{node: top.node.Left})
		stack = append(stack, top)
	}
	return result
}
