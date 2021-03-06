## 背包问题
一种特殊的动态规划
### 问题雏形
```text
有 N 件物品，一个容量为 V 的背包。
第i件物品的体积是 C[i], 价值是 W[i]。
在不超过背包容量的情况下，怎么装物品使得装入的总价值最大？求出最大总价值。
```
```text
假设背包的大小是 10，有 4 个物品，
体积分别是 [2,3,5,7]，
价值分别是 [2,5,2,5]。
```

1、如果仅考虑将前一个物品放入背包，只要背包容量大于2，都可以获取价值为2的最大价值

`表格首列为物品编号+(物品体积，物品价值）`

|物品\容量| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
|:------:|---|---|---|---|---|---|---|---|---|---|----|
|1 (2,2) |0  |0  |2  |2  |2  |2  |2  |2  |2  |2  |2   |
|2 (3,5) |
|3 (5,2) |
|4 (7,5) |

2、现在考虑仅将前两个物品放入背包, 如果背包容量不小于5就可以把两个物品都放入，获得价值`2+5=7`;
如果不能全部放入，就需要选择体积不超，价值最大的物品

|物品\容量| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
|:------:|---|---|---|---|---|---|---|---|---|---|----|
|1 (2,2) |0  |0  |2  |2  |2  |2  |2  |2  |2  |2  |2   |
|2 (3,5) |0  |0  |2  |5  |5  |7  |7  |7  |7  |7  |7   |
|3 (5,2) |
|4 (7,5) |

3、同理，考虑仅将前三个物品放入背包

|物品\容量| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
|:------:|---|---|---|---|---|---|---|---|---|---|----|
|1 (2,2) |0  |0  |2  |2  |2  |2  |2  |2  |2  |2  |2   |
|2 (3,5) |0  |0  |2  |5  |5  |7  |7  |7  |7  |7  |7   |
|3 (5,2) |0  |0  |2  |5  |5  |7  |7  |7  |7  |7  |9   |   
|4 (7,5) |

4、现在考虑将全部四个物品放入背包，可以依据前三个物品放入的结果来制定方案

|物品\容量| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
|:------:|---|---|---|---|---|---|---|---|---|---|----|
|1 (2,2) |0  |0  |2  |2  |2  |2  |2  |2  |2  |2  |2   |
|2 (3,5) |0  |0  |2  |5  |5  |7  |7  |7  |7  |7  |7   |
|3 (5,2) |0  |0  |2  |5  |5  |7  |7  |7  |7  |7  |9   |   
|4 (7,5) |0  |0  |2  |5  |5  |7  |7  |7  |7  |7  |10  |

梳理发现，求前`n`个物品在背包容量为v的情况下能得到的最大价值`f(n,v)`，可以依赖前`n-1`个物品的情况来得到，不难发现：

```text
1. 如果背包总容量小于 C[n]，不能装第n个物品，得到的价值是 f(n-1, v)
2. 装第 n 个物品，得到的价值是 W[n] + f(n-1, v-C[n])
即使第 n 个物品可以装，得到的结果也不一定比不装好，
所以 f(n,v) = max(f(n-1, v), W[n] +f(n-1, v-C[n])
```
我们梳理出物品数量和背包容量两个状态，在这两个状态上实施动态规划即可
```text
定义dp[
```
初始状态：


