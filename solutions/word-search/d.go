package search

/*
79. 单词搜索 https://leetcode-cn.com/problems/word-search/
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false


提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
常规dfs
*/
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	i := 0
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r < 0 || c < 0 || r == m || c == n ||
			visited[r][c] || board[r][c] != word[i] {
			return false
		}
		if i == len(word)-1 {
			return true
		}
		i++
		visited[r][c] = true
		if dfs(r+1, c) || dfs(r-1, c) || dfs(r, c+1) || dfs(r, c-1) {
			return true
		}
		i--
		visited[r][c] = false
		return false
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if dfs(r, c) {
				return true
			}
		}
	}
	return false
}
