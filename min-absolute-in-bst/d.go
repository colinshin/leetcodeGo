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
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}

	minL := getMinimumDifference(root.Left)
	minR := getMinimumDifference(root.Right)

	if root.Left != nil {
		max := max(root.Left)
		abs := absVal(max, root.Val)
		minL = minVal(minL, abs)
	}

	if root.Right != nil {
		min := min(root.Right)
		abs := absVal(root.Val, min)
		minR = minVal(minR, abs)
	}

	return minVal(minL, minR)
}

func max(t *TreeNode) int {
	if t.Right == nil {
		return t.Val
	}
	return max(t.Right)
}

func min(t *TreeNode) int {
	if t.Left == nil {
		return t.Val
	}
	return min(t.Left)
}

func absVal(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
