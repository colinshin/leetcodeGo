/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package n_queen

/*
51. N皇后 https://leetcode-cn.com/problems/n-queens/
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
示例:
输入: 4
输出: [
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
*/
/*
常规回溯
时间复杂度O(n!)；空间复杂度O(n)
*/
func solveNQueens(n int) [][]string {
	var result [][]string
	backtrack(0, makeBoard(n), &result)
	return result
}

func makeBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

// 在行r找到合适的列放置皇后
func backtrack(r int, board [][]byte, result *[][]string) {
	if r == len(board) {
		*result = append(*result, parse(board))
		return
	}
	for c := 0; c < len(board); c++ {
		if !canSetQueen(board, r, c) {
			continue
		}
		board[r][c] = 'Q'
		backtrack(r+1, board, result)
		board[r][c] = '.'
	}
}

func parse(board [][]byte) []string {
	r := make([]string, len(board))
	for i := range r {
		r[i] = string(board[i])
	}
	return r
}

func canSetQueen(board [][]byte, r, c int) bool {
	var i, j int
	for i = 0; i < r; i++ { // top
		if board[i][c] == 'Q' {
			return false
		}
	}
	for i, j = r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // topLeft
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j = r-1, c+1; i >= 0 && j < len(board); i, j = i-1, j+1 { // topRight
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

/*
52. N皇后 II https://leetcode-cn.com/problems/n-queens-ii/

与问题51类似，更简单些，只需要返回不同的解决方案的数量
*/
func totalNQueens(n int) int {
	var total int
	backtrack1(0, makeBoard1(n), &total)
	return total
}

func makeBoard1(n int) [][]bool {
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	return board
}

func backtrack1(r int, board [][]bool, total *int) {
	if r == len(board) {
		*total++
		return
	}
	for c := 0; c < len(board); c++ {
		if !canSetQueen1(board, r, c) {
			continue
		}
		board[r][c] = true
		backtrack1(r+1, board, total)
		board[r][c] = false
	}
}

func canSetQueen1(board [][]bool, r, c int) bool {
	for i := 0; i < r; i++ {
		if board[i][c] {
			return false
		}
	}
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	for i, j := r-1, c+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}
	return true
}
