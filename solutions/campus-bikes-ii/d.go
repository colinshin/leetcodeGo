/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package campus_bikes_ii

import "math"

/*
在由 2D 网格表示的校园里有 n 位工人（worker）和 m 辆自行车（bike），n <= m。所有工人和自行车的位置都用网格上的 2D 坐标表示。

我们为每一位工人分配一辆专属自行车，使每个工人与其分配到的自行车之间的曼哈顿距离最小化。

p1 和 p2 之间的曼哈顿距离为 Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|。

返回每个工人与分配到的自行车之间的曼哈顿距离的最小可能总和。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/campus-bikes-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
0-1背包的解法：
工人数等价于物品数
自行车数等价于容量
距离等价于价值
*/
func assignBikes(workers [][]int, bikes [][]int) int {
	dp := make(map[int]int, 0)
	dp[0] = 0
	for _, worker := range workers {
		newDp := make(map[int]int, 0)
		for j, bike := range bikes {
			for key, value := range dp {
				if key&(1<<uint(j)) != 0 {
					continue
				}
				lastKey := key | (1 << uint(j))
				newV := value + dist(worker[0], worker[1], bike[0], bike[1])
				if _, ok := newDp[lastKey]; !ok || newDp[lastKey] > newV {
					newDp[lastKey] = newV
				}
			}
		}
		dp = newDp
	}
	min := math.MaxInt32
	for _, v := range dp {
		if v < min {
			min = v
		}
	}
	return min
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
