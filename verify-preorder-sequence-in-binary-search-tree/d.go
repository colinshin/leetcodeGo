package verify_preorder_sequence_in_binary_search_tree

import "math"

/*
给定一个整数数组，你需要验证它是否是一个二叉搜索树正确的先序遍历序列。

你可以假定该序列中的数都是不相同的。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [5,2,6,1,3]
输出: false
示例 2：

输入: [5,2,1,3,6]
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 利用单调递减栈
func verifyPreorder1(preorder []int) bool {
	var stack []int
	min := math.MinInt32
	for _, v := range preorder {
		if v < min {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] < v {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	return true
}

// 用原数组模拟单调递减栈
func verifyPreorder(preorder []int) bool {
	min := math.MinInt32
	i := -1
	for _, v := range preorder {
		if v < min {
			return false
		}
		for i >= 0 && preorder[i] < v {
			min = preorder[i]
			i--
		}
		i++
		preorder[i] = v
	}
	return true
}
