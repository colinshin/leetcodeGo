# 01背包问题说明：

## 定义：
```
有一个背包，容量是c；现在有n个物品，编号为0,...,n-1
数组w表示每个物品的重量，数组v表示每个物品的价值，即：
对于物品i，重量为w[i]， 价值为v[i]
向背包里装物品，怎么装可以使物品的总价值最大，当然，不能超过背包的容量
```
## 思路：
```这类组合问题，都可以使用递归来完成。如果能在其中找到重叠子问题(最优子结构), 则可以转化成记忆化搜索或动态规划来解决。```

## 状态定义：
```
对于这个问题，定义f(n,c)，表示n个物品，背包容量为c时得到的结果
可以减小物品数量和背包容量来简化
有两个约束条件：
1.物品数量介于0到n之间
2.容量介于0到c之间
```
## 状态转移：
```
对于i，c；有两种情况：将物品i加入和忽略物品i
则f(i, c) = max(f(i-1, c), f(i-1, c-w[i]+v[i])
```
我们先用记忆化搜索来实现
```
func zeroOnePack0(w, v []int, c int) int {
	if len(w) != len(v) || c < 0 {
		return -1
	}
	if len(w) == 0 || c == 0 {
		return 0
	}
	n := len(w)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, c+1)
		for j := 0; j <= c; j++ {
			dp[i][j] = -1
		}
	}
	return help(n-1, c, w, v, dp)
}

func help(id, c int, w, v []int, dp [][]int) int {
	if c <= 0 || id < 0 {
		return 0
	}
	if dp[id][c] != -1 {
		return dp[id][c]
	}
	max := help(id-1, c, w, v, dp)
	if c >= w[id] {
		max = maxInt(max, help(id-1, c-w[id], w, v, dp)+v[id])
	}
	dp[id][c] = max
	return max
}
```

可以用动态规划自底向上解决
```
func zeroOnePack1(w, v []int, c int) int {
	if len(w) != len(v) || c < 0 {
		return -1
	}
	if len(w) == 0 || c == 0 {
		return 0
	}
	n := len(w)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, c+1)
	}
	for j := 0; j <= c; j++ {
		if j >= w[0] {
			dp[0][j] = v[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= c; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= w[i] {
				dp[i][j] = maxInt(dp[i][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	return dp[n-1][c]
}
```
上边的时间、空间复杂度都是O(n*c)
其实根据状态转移方程，f(i, c) = max(f(i-1, c), f(i-1, c-w[i])
第i行元素只依赖于第i-1行元素。理论上，只需要保持两行元素。空间复杂度：O(2*C)=O(C)
```
func zeroOnePack2(w, v []int, c int) int {
	if len(w) != len(v) || c < 0 {
		return -1
	}
	if len(w) == 0 || c == 0 {
		return 0
	}
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, c+1), make([]int, c+1)
	for j := 0; j <= c; j++ {
		if j >= w[0] {
			dp[0][j] = v[0]
		}
	}
	for i := 1; i < len(w); i++ {
		for j := 0; j <= c; j++ {
			dp[i%2][j] = dp[(i-1)%2][j]
			if j >= w[i] {
				dp[i%2][j] = maxInt(dp[i%2][j], dp[(i-1)%2][j-w[i]]+v[i])
			}
		}
	}
	return dp[(len(w)-1)%2][c]
}
```
继续优化，dp只用一行解决， 需要注意从后向前遍历
看代码理解：
```
func pack01(w, v []int, c int) int {
	if len(w) != len(v) || c < 0 {
		return -1
	}
	if len(w) == 0 || c == 0 {
		return 0
	}
	dp := make([]int, c+1)
	for j := 0; j <= c; j++ {//只用一个物品
		if j >= w[0] {
			dp[j] = v[0]
		}
	}
	for i := 1; i < len(w); i++ {//物品不断增加
		for j := c; j >= w[i]; j-- {//容量不断减小
			dp[j] = maxInt(dp[j], dp[j-w[i]]) + v[i]
		}
	}
	return dp[c]
}
```
附：maxInt实现
```
func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
```