/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package verify_postorder_sequence_inbinary_search_tree

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
思路通判断一个序列是不是BST的先序遍历结果
有递归和借助单调栈的两种解法
*/
// 递归解法
func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n == 0 {
		return true
	}
	root := postorder[n-1]
	i := 0
	for ; i < n-1; i++ {
		if postorder[i] > root {
			break
		}
	}
	// i为右子树根所在位置
	for j := i; j < n-1; j++ {
		if postorder[j] < root {
			return false
		}
	}
	// 截止目前，检查了左子树所有元素小于root，右子树所有元素大于root
	if i > 0 && !verifyPostorder(postorder[:i]) {
		return false
	}
	if i < n-1 && !verifyPostorder(postorder[i:n-1]) {
		return false
	}
	return true
}

// TODO: 单调栈解法
