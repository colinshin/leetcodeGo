/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package validate_binary_search_tree

import (
	"math"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true

示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
按直觉递归：
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
    if root.Left == nil {
        return root.Val < root.Right.Val && isValidBST(root.Right)
    }
    if root.Right == nil {
        return root.Val > root.Left.Val && isValidBST(root.Left)
    }
    return root.Left.Val < root.Val &&
    root.Val < root.Right.Val &&
    isValidBST(root.Left) &&
    isValidBST(root.Right)
}
这样的解法是错的，不能简单比较节点的左右子节点的值和其自身值；实际BST要求左子树所有节点值<根节点值<所有右子树节点值
实际应该对每个节点的值与上下界做判断
时空复杂度均为O(n), n为节点总数
*/
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

/*
如果一个二叉树式BST， 那么所有元素的值都在开区间(min, max)里
时空复杂度均为O(n), n为节点总数
*/
func helper(t *TreeNode, min, max int) bool {
	switch {
	case t == nil:
		return true
	case t.Val <= min, t.Val >= max:
		return false
	case !helper(t.Left, min, t.Val):
		return false
	case !helper(t.Right, t.Val, max):
		return false
	default:
		return true
	}
}

/* 递归式中序遍历
可以按照中序遍历的顺序，将所有节点的值存入一个数组，再检查数组是否升序排序的即可（遍历数组，判断每个元素是否大于其前的一个元素）
实际上空间可以优化，并不需要一个数组，只需要一个int变量记录前一个元素的值即可
时空复杂度均为O(n), n为节点总数
*/
func isValidBST0(root *TreeNode) bool {
	prev := math.MinInt64
	var bfs func(*TreeNode) bool
	bfs = func(t *TreeNode) bool {
		if t == nil {
			return true
		}
		if !bfs(t.Left) {
			return false
		}
		if t.Val <= prev {
			return false
		}
		prev = t.Val
		if !bfs(t.Right) {
			return false
		}
		return true
	}
	return bfs(root)
}

/* 借助栈，迭代式中序遍历
时空复杂度均为O(n), n为节点总数
*/
func isValidBST1(root *TreeNode) bool {
	prev := math.MinInt64
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= prev {
			return false
		}
		prev = root.Val
		root = root.Right
	}
	return true
}
