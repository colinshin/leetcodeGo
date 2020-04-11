package elimination

import (
	"container/list"
	"math"
)

/*
1293. 网格中的最短路径 https://leetcode-cn.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。
如果找不到这样的路径，则返回 -1。

示例 1：
输入：
grid =
[[0,0,0],
 [1,1,0],
 [0,0,0],
 [0,1,1],
 [0,0,0]],
k = 1
输出：6
解释：
不消除任何障碍的最短路径是 10。
消除位置 (3,2) 处的障碍后，最短路径是 6 。该路径是 (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).


示例 2：
输入：
grid =
[[0,1,1],
 [1,1,1],
 [1,0,0]],
k = 1
输出：-1
解释：
我们至少需要消除两个障碍才能找到这样的路径。

提示：

grid.length == m
grid[0].length == n
1 <= m, n <= 40
1 <= k <= m*n
grid[i][j] == 0 or 1
grid[0][0] == grid[m-1][n-1] == 0
*/

/*
DFS

*/
func shortestPath(grid [][]int, k int) int {
	const maxSize = 40
	m, n := len(grid), len(grid[0])
	// 设想一种情况，只向右或向下走，将会是最短路径，共（m-1）+（n-1）步，已知起点终点都不是障碍，那么这条最短路径上最多有障碍m+n-3个
	// 如果k不小于m+n-3就可以走这条最短路径，这里不判断都话，会导致无谓都遍历尝试
	if k >= m+n-3 {
		return m + n - 2
	}
	passed := 0 // paased代表之前已经走了多少步
	visited := [maxSize][maxSize]bool{}
	var dfs func(r, c, k int) int
	dfs = func(r, c, k int) int { // 返回从起点走到（r，c）处，再走到终点最少需要多少步，k代表还有多少次清理路障的机会
		if r < 0 || r >= m || c < 0 || c >= n ||
			visited[r][c] || grid[r][c] == 1 && k == 0 {
			return math.MaxInt64
		}
		if r == m-1 && c == n-1 {
			return passed
		}
		visited[r][c] = true
		if grid[r][c] == 1 {
			k--
		}
		passed++
		left := dfs(r, c-1, k)
		right := dfs(r, c+1, k)
		up := dfs(r-1, c, k)
		down := dfs(r+1, c, k)
		// 回溯，非常重要；别的路径尝试需要能走到（r，c）处
		passed--
		visited[r][c] = false

		return min(min(left, right), min(up, down))
	}
	r := dfs(0, 0, k)
	if r == math.MaxInt64 {
		return -1
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
BFS，除了横坐标、纵坐标，需要把剩余能清理的障碍数量作为另一个状态传递给bfs的队列，方便判断
时空复杂度都是O(MN∗min(M+N,K))
*/
func shortestPath0(grid [][]int, k int) int {
	const maxSize = 40
	const maxK = maxSize + maxSize - 3
	m, n := len(grid), len(grid[0])
	// 设想一种情况，只向右或向下走，将会是最短路径，共（m-1）+（n-1）步，已知起点终点都不是障碍，那么这条最短路径上最多有障碍m+n-3个
	// 如果k不小于m+n-3就可以走这条最短路径，这里不判断都话，会导致无谓都遍历尝试
	if k >= m+n-3 {
		return m + n - 2
	}
	visited := [maxSize][maxSize][maxK + 1]bool{}
	queue := list.New()
	queue.PushBack([]int{0, 0, k})
	visited[0][0][k] = true
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	passed := 0 // paased代表之前已经走了多少步
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ { // 当前层
			info := queue.Remove(queue.Front()).([]int)
			r, c, k := info[0], info[1], info[2]
			if r == m-1 && c == n-1 {
				return passed
			}
			if grid[r][c] == 1 {
				k--
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n ||
					visited[nr][nc][k] || grid[nr][nc] == 1 && k == 0 {
					continue
				}
				queue.PushBack([]int{nr, nc, k})
				visited[nr][nc][k] = true
			}
		}
		passed++ // 一层处理完，步数加1
	}
	return -1
}
