/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rotate_image

/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// solution with a help matrix
func rotate1(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}
	helper := make([][]int, n)
	for i := 0; i < n; i++ {
		helper[i] = make([]int, n)
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			//for a point (r, c), will be rotated to (c, n-1-r)
			helper[c][n-1-r] = matrix[r][c]
		}
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			matrix[r][c] = helper[r][c]
		}
	}
}

// transpose matrix and then reverse each row
func rotate2(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}
	// transpose matrix
	for r := 0; r < n; r++ {
		for c := r; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
	// reverse each row
	for r := 0; r < n; r++ {
		row := matrix[r]
		for i := 0; i < n/2; i++ {
			row[i], row[n-1-i] = row[n-1-i], row[i]
		}
	}
}

/*
when we rotate points, for egxample:
p1 -> p2 ; then p2 -> p3; then p3 -> p4, and actually p4 will be rotated to p1：

p1 → p2
↑     ↓
p4 ← p3

there is a circle!

for a point (r, c), will be rotated to (c, n-1,r)
if p1 is point (r, c):
p1	(r, c)			→ p2
p2	(c, n-1-r)		→ p3
p3	(n-1-r, n-1-c)	→ p4
p4	(n-1-c, r)		→ p1
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}
	for r := 0; r < (n+1)/2; r++ {
		for c := 0; c < n/2; c++ {
			// p1, p2, p3, p4
			matrix[r][c], matrix[c][n-1-r], matrix[n-1-r][n-1-c], matrix[n-1-c][r] =
				// p4, p1, p2, p3
				matrix[n-1-c][r], matrix[r][c], matrix[c][n-1-r], matrix[n-1-r][n-1-c]
		}
	}
}

/*
when rotate anticlockwise, for egxample:
p1 -> p2 ; then p2 -> p3; then p3 -> p4, and actually p4 will be rotated to p1：

p2 ← p1
↓     ↑
p3 → p4

there is a circle!

for a point(r, c), will be rotated to (n-1-c, r)
if p1 is point (i, j):
p1	(r, c)			→ p2
p2	(n-1-c, r)		→ p3
p3	(n-1-r, n-1-c)	→ p4
p4	(c, n-1-r)		→ p1
*/
func rotateAnticlockwise(s [][]int) {
	n := len(s)
	if n < 2 {
		return
	}
	for r := 0; r < (n+1)/2; r++ {
		for c := 0; c < n/2; c++ {
			s[r][c], s[n-1-c][r], s[n-1-r][n-1-c], s[c][n-1-r] =
				s[c][n-1-r], s[r][c], s[n-1-c][r], s[n-1-r][n-1-c]

		}
	}
}

func rotateAnticlockwise1(s [][]int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 0; i < 3; i++ {
		rotate(s)
	}
}
