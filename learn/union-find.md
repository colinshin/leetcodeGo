# UnionFind
## 背景
为了解决类似[最低成本联通所有城市](solutions/connecting-cities-with-minimum-cost/d.go)这样的问题，引入一个非常有用的数据结构UnionFind，即并查集<br>
还有些问题如[[1061] 按字典序排列最小的等效字符串](../solutions/lexicographically-smallest-equivalent-string/readme.md)也用它，用了都说好~
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
func (uf UnionFind) Find(x int) int { // 路径压缩
	for x != uf[x] {
		uf[x] = uf.Find(uf[x])
	}
	return uf[x]
}
func (uf UnionFind) Union(x, y int) {
    rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY   // 可以按秩合并，即高度较小的根插入高度较大的根下面，进一步减少整个Union、Find操作的复杂度
}
```
find函数还可以循环实现：
```
func (uf UnionFind) Find(x int) int {
    for uf[x] != x {
        x, uf[x] = uf[x], uf[uf[x]]
    }
    reutn x
}
```
或者
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