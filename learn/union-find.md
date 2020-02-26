# UnionFind
## 背静
为了解决类似[最低成本联通所有城市](solutions/connecting-cities-with-minimum-cost/d.go)这样的问题，引入一个非常有用的数据结构UnionFind，即并查集
## 一个有意思的说明
https://blog.csdn.net/niushuai666/article/details/6662911
## 实现
```
type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i, _ := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}
func (uf UnionFind) Find(x int) int {
	for x != uf[x] {
		uf[x] = uf.Find(uf[x])
	}
	return uf[x]
}
func (uf UnionFind) Union(x, y int) {
	uf[x] = y
}
```
find函数还可以循环实现：
```
func (uf UnionFind) Find(x int) int {
	root := x
	for root != uf[root] {
		root = uf[root]
	}
	for root != x {
		uf[x], x = root, uf[x]
	}
	return root
}
```