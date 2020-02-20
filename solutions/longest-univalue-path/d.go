/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_univalue_path

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:

2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:

2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-univalue-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
对于一个节点node， 以其为根节点的路径就是分别向左下和右下延伸形成的箭头

设计一个递归函数，返回当前节点的最大同值路径
如果该节点的值与左右子树的值都不相等，则其最大同值路径为0
如果该节点的值与左右子树的值都相等，则其最大同值路径是其左子树最大同值路径+1 + 右子树最大同值路径+1
如果该节点的值等于其左子树的值但不等于右子树的值，则最长同值路径为左子树的最长同值路径+1
如果该节点的值等于其右子树的值，则最长同值路径为右子树的最长同值路径+1

我们用一个全局变量记录这个最大值，不断更新
*/
func longestUnivaluePath(root *TreeNode) int {
	result := 0

	var caculate func(root *TreeNode) int
	caculate = func(root *TreeNode) int { // 返回与root值相同的左子树最大路径和与root值相同的右子树最大路径里边较大的路径，方便递归
		if root == nil {
			return 0
		}
		left := caculate(root.Left)
		right := caculate(root.Right)
		if root.Left != nil && root.Left.Val == root.Val {
			left++
		} else {
			left = 0
		}
		if root.Right != nil && root.Right.Val == root.Val {
			right++
		} else {
			right = 0
		}
		if left+right > result {
			result = left + right
		}
		return max(left, right)
	}

	_ = caculate(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
