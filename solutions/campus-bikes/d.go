/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package campus_bikes

import (
	"sort"
)

/*
在由 2D 网格表示的校园里有 n 位工人（worker）和 m 辆自行车（bike），n <= m。所有工人和自行车的位置都用网格上的 2D 坐标表示。

我们需要为每位工人分配一辆自行车。在所有可用的自行车和工人中，我们选取彼此之间曼哈顿距离最短的工人自行车对  (worker, bike) ，并将其中的自行车分配給工人。
如果有多个 (worker, bike) 对之间的曼哈顿距离相同，那么我们选择工人索引最小的那对。类似地，如果有多种不同的分配方法，则选择自行车索引最小的一对。
不断重复这一过程，直到所有工人都分配到自行车为止。

给定两点 p1 和 p2 之间的曼哈顿距离为 Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|。

返回长度为 n 的向量 ans，其中 a[i] 是第 i 位工人分配到的自行车的索引（从 0 开始）。

链接：https://leetcode-cn.com/problems/campus-bikes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func assignBikes(workers [][]int, bikes [][]int) []int {
	type item struct {
		workerId int
		bikeId   int
		dist     int
	}
	var items []item
	for i, worker := range workers {
		for j, bike := range bikes {
			distance := dist(worker[0], worker[1], bike[0], bike[1])
			items = append(items, item{workerId: i, bikeId: j, dist: distance})
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].dist == items[j].dist {
			if items[i].bikeId == items[j].bikeId {
				return items[i].workerId < items[j].workerId
			}
			return items[i].bikeId < items[j].bikeId
		}
		return items[i].dist < items[j].dist
	})

	workerUsed := make([]bool, len(workers))
	bikeUsed := make([]bool, len(bikes))
	result := make([]int, len(workers))
	for _, item := range items {
		if workerUsed[item.workerId] || bikeUsed[item.bikeId] {
			continue
		}
		result[item.workerId] = item.bikeId
		workerUsed[item.workerId] = true
		bikeUsed[item.bikeId] = true
	}
	return result
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
