package children_weight_order

import (
	"fmt"
	"sort"
)

/*
假设有搅乱顺序的一群儿童成一个队列。 每个儿童由一个整数对(w, k)表示，
其中 w 是这个儿童 的体重，k 是排在这个儿童前面且体重大于或等于 w 的儿童数量。
编写一个算法来重建这个队 列。

注意:

总儿童数量少于 1100 人。
输入的数据是合理的，只是顺序给打乱了。

示例

输入:
[[8,0], [4,4], [8,1], [5,0], [6,1], [5,2]]

输出:
[[5,0], [8,0], [5,2], [6,1], [4,4], [8,1]]
解释:[5,2]前面两个儿童的体重分别是 5 和 8，且只有两个儿童;[6,1]前面只有[8,0]儿童的体重 大于他/她，并且不能和[5,2]换位置，否则会导致[5,2]的 2 不对。
*/

func reconstructQueue(people [][]int) [][]int {
	// 先根据k从小到大排序
	sort.SliceStable(people, func(i, j int) bool {
		return people[i][1] < people[j][1]
	})
	// 由w、k微调顺序
	for i := 1; i < len(people); i++ {
		actW := people[i][0]
		actK := people[i][1]
		countK := 0
		for j := 0; j < i; j++ {
			if people[j][0] >= actW {
				countK++
			}
		}
		fmt.Println("------", i, countK, actK)
		if countK > actK { // i儿童需要往前调整; 如果相等，无需调整；不会出现小于的情况
			swaped := 0
			for j := i - 1; j >= 0; j-- {
				if people[j][0] >= actW {
					swaped++
				}
				people[j], people[j+1] = people[j+1], people[j]
				if swaped == countK-actK {
					break
				}
			}
		}
	}
	return people
}
