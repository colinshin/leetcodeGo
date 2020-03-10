/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package flood_fill

import "container/list"

/*
颜色填充
编写函数，实现许多图片编辑软件都支持的“颜色填充”功能。
给定一个屏幕（以二维数组表示，元素为颜色值）、一个点和一个新的颜色值，将新颜色值填入这个点的周围区域，直到原来的颜色值全都改变。

示例1:

 输入：
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
 输出：[[2,2,2],[2,2,0],[2,0,1]]
 解释:
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
说明：

image 和 image[0] 的长度在范围 [1, 50] 内。
给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/color-fill-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
意思是将与给定点相连且颜色相同的点都改变颜色

DFS递归，时空复杂度都是O(N)
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return nil
	}
	var fill func(r, c, newColor, oldColor int)
	fill = func(r, c, newColor, oldColor int) { // DFS
		if r < 0 || r == len(image) || c < 0 || c == len(image[0]) ||
			newColor == image[r][c] || image[r][c] != oldColor {
			return
		}
		image[r][c] = newColor
		fill(r-1, c, newColor, oldColor)
		fill(r+1, c, newColor, oldColor)
		fill(r, c-1, newColor, oldColor)
		fill(r, c+1, newColor, oldColor)
	}
	fill(sr, sc, newColor, image[sr][sc])
	return image
}

/*
BFS迭代。时空复杂度都是O(N)
*/
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return nil
	}
	oldColor := image[sr][sc]
	if newColor == oldColor {
		return image
	}
	direct := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := list.New() // 记录坐标
	queue.PushBack([]int{sr, sc})
	for queue.Len() > 0 {
		rc := queue.Remove(queue.Front()).([]int)
		image[rc[0]][rc[1]] = newColor // 着色
		for _, v := range direct {
			r, c := rc[0]+v[0], rc[1]+v[1] // 四个方向
			if r >= 0 && r < len(image) && c >= 0 && c < len(image[0]) &&
				newColor != image[r][c] && image[r][c] == oldColor {
				queue.PushBack([]int{r, c})
			}
		}
	}
	return image
}
