package connecting_cities_with_minimum_cost

import "sort"

func minimumCost(n int, connections [][]int) int {
	if len(connections) < n-1 { // 要有每个城市的联接信息，最终才能将所有城市联通，否则总有落单的
		return -1
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	unionFind := NewUnionFind(n)
	connected, result, i := 0, 0, 0
	for connected < n-1 {
		connection := connections[i]
		i++
		city1, city2 := connection[0]-1, connection[1]-1
		city1, city2 = unionFind.find(city1), unionFind.find(city2)
		if city1 != city2 {
			unionFind.join(city1, city2)
			connected++
			result += connection[2]
		}
	}
	return result
}

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}
func (uf UnionFind) find(x int) int {
	root := x
	for root != uf[root] {
		root = uf[root]
	}
	for root != x {
		uf[x], x = root, uf[x]
	}
	return root
}
func (uf UnionFind) join(x, y int) {
	uf[x] = y
}
