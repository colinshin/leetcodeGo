/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package matrix_cells_in_distance_order

import (
	"math"
	"sort"
)

/*
给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。

另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。

返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，
其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）


示例 1：

输入：R = 1, C = 2, r0 = 0, c0 = 0
输出：[[0,0],[0,1]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1]
示例 2：

输入：R = 2, C = 2, r0 = 0, c0 = 1
输出：[[0,1],[0,0],[1,1],[1,0]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2]
[[0,1],[1,1],[0,0],[1,0]] 也会被视作正确答案。
示例 3：

输入：R = 2, C = 3, r0 = 1, c0 = 2
输出：[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2,2,3]
其他满足题目要求的答案也会被视为正确，例如 [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/matrix-cells-in-distance-order
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
一、最直观的做法，起个名字：方形水波遍历~
从(r0, c0)出发，按照曼哈顿距离向外扩展，可以看到对于曼哈顿距离为定值的点，连接起来是一个旋转了45°的正方形，当然正方形的中心是（r0，c0）
可以将这些正方形形象地称作方形水波~
实际就是从(r0, c0)出发，从内向外层层遍历那些方形水波
对于每个方形水波，有上下左右4个顶点，
每次可以先从上顶点开始顺时针遍历，先遍历到右顶点，再遍历到下顶点，再到左顶点，最后回归上顶点；
当然这道题对顺序没有要求，可以从任意顶点开始，顺、逆时针遍历都行
注意方形水波的某些点可能超出矩阵范围，处理好边界即可

时空复杂度都是O(R*C)。
时间上，会有一些超出矩阵的点去遍历，但总的量级在R*C；
空间复杂度主要在最后的结果数组上，这是没法省的，除去结果数组，方形水波遍历的解法显然是常数级的空间复杂度
*/
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, R*C)
	result[0] = []int{r0, c0}
	step := 1
	r, c := r0, c0
	for step < R*C {
		r--          // 从上一次遍历后的方形水波上顶点再上一步到达这次需要遍历的方形水波上顶点
		for r < r0 { // 上顶点 -> 右顶点
			if r >= 0 && c <= C-1 {
				result[step] = []int{r, c}
				step++
			}
			r++
			c++
		}
		for c > c0 { // 右顶点 -> 下顶点
			if r <= R-1 && c <= C-1 {
				result[step] = []int{r, c}
				step++
			}
			r++
			c--
		}
		for r > r0 { // 下顶点 -> 左顶点
			if r <= R-1 && c >= 0 {
				result[step] = []int{r, c}
				step++
			}
			r--
			c--
		}
		for c < c0 { // 左顶点 -> 上顶点 // 这里的判断是c<c0, 保证了不会将上顶点重复放入结果
			if r >= 0 && c >= 0 {
				result[step] = []int{r, c}
				step++
			}
			r--
			c++
		}
	}
	return result
}

/*
二、借助哈希表
1.先用一个哈希表记录每个点到(r0, c0)的曼哈顿距离，以曼哈顿距离为key， 这些点的坐标组成的数组为value
2.再按照曼哈顿距离从小到大遍历哈希表即可
虽然Go里没有可排序的map，但对于这道题，其实比较容易实现上面的第2步;因为可以确定，曼哈顿距离是从0到最大曼哈顿距离的连续序列
距离(r0,c0)最大的曼哈顿距离很好求，就是矩阵4个顶点距离(r0,c0)的最大值；当然也可以在第1步确定这个最大值
时间复杂度O(R*C*2)=O(R*C)，空间复杂度O(R*C*2)=O(R*C)
*/
func allCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	m := make(map[int][][]int, 0)
	maxDist := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			dist := abs(r-r0) + abs(c-c0)
			if dist > maxDist {
				maxDist = dist
			}
			m[dist] = append(m[dist], []int{r, c})
		}
	}

	result := make([][]int, R*C)
	k := 0
	for i := 0; i <= maxDist; i++ {
		for _, v := range m[i] {
			result[k] = v
			k++
		}
	}
	return result
}

/*
三、先将所有点放入结果，再对结果排序
时间复杂度O(R*C*lg(R*C))， 空间复杂度O(R*C)
*/
func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, R*C)
	k := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			result[k] = []int{r, c}
			k++
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return abs(result[i][0]-r0)+abs(result[i][1]-c0) < abs(result[j][0]-r0)+abs(result[j][1]-c0)
	})
	return result
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
