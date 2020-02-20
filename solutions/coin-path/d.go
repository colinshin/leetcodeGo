/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package coin_path

import "math"

/*
给定一个数组 A（下标从 1 开始）包含 N 个整数：A1，A2，……，AN 和一个整数 B。
你可以从数组 A 中的任何一个位置（下标为 i）跳到下标 i+1，i+2，……，i+B 的任意一个可以跳到的位置上。
如果你在下标为 i 的位置上，你需要支付 Ai 个金币。如果 Ai 是 -1，意味着下标为 i 的位置是不可以跳到的。

现在，你希望花费最少的金币从数组 A 的 1 位置跳到 N 位置，你需要输出花费最少的路径，依次输出所有经过的下标（从 1 到 N）。

如果有多种花费最少的方案，输出字典顺序最小的路径。

如果无法到达 N 位置，请返回一个空数组。


样例 1 :

输入: [1,2,4,-1,2], 2
输出: [1,3,5]


样例 2 :

输入: [1,2,4,-1,2], 1
输出: []


注释 :

路径 Pa1，Pa2，……，Pan 是字典序小于 Pb1，Pb2，……，Pbm 的，
当且仅当第一个 Pai 和 Pbi 不同的 i 满足 Pai < Pbi，如果不存在这样的 i 那么满足 n < m。
A1 >= 0。 A2, ..., AN （如果存在） 的范围是 [-1, 100]。
A 数组的长度范围 [1, 1000].
B 的范围 [1, 100].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/* [未通过]
一个更容易理解的实现如下，但是有用例失败
如下用例失败：
输入:
[0,0,0,0,0,0]
3
输出
[1,3,6]
预期结果
[1,2,3,4,5,6]

原来是没有满足字典序要求！
*/
func cheapestJump(A []int, B int) []int {
	n := len(A)
	if n == 0 || -1 == A[n-1] {
		return []int{}
	}
	dp := make([]int, n)
	dp[n-1] = A[n-1]
	path := make([][]int, n)
	for i := 1; i < n; i++ {
		if A[i] == -1 {
			dp[i] = math.MaxInt32
			continue
		}
		minCost := math.MaxInt32
		lastIndex := -1
		for j := int(math.Max(float64(i-B), 0)); j < i; j++ {
			if dp[j] == math.MaxInt32 {
				continue
			}
			if dp[j] < minCost {
				minCost = dp[j]
				lastIndex = j
			}
		}
		if lastIndex != -1 {
			dp[i] = A[i] + minCost
			path[i] = append(path[lastIndex], lastIndex+1)
		} else {
			dp[i] = math.MaxInt32
		}
	}
	if dp[n-1] == math.MaxInt32 {
		return []int{}
	}
	return append(path[n-1], n)
}

/*
逆向思考：
从1走到N，跟从N走到1，花费最小的计算是等价的
从N走到1的同时记录路径，这个路径反过来就是方案之一，且是字典序最小的方案
参考上一解法没通过的用例分析就能明白，为何逆向走能得到字典序最小的路径

注意path定义为[]int，与上一解法不同；这样节省空间；只需最后将所有路径串起来即可
*/
func cheapestJump1(A []int, B int) []int {
	n := len(A)
	if n == 0 || A[n-1] == -1 {
		return []int{}
	}
	dp, path := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i], path[i] = math.MaxInt32, -1
	}
	dp[n-1] = A[n-1]

	for i := n - 2; i >= 0; i-- {
		if A[i] == -1 {
			continue
		}
		last := int(math.Min(float64(i+B), float64(n-1)))
		for j := i + 1; j <= last; j++ {
			if dp[j] == math.MaxInt32 {
				continue
			}
			if A[i]+dp[j] < dp[i] {
				dp[i], path[i] = A[i]+dp[j], j
			}
		}
	}
	if dp[0] == math.MaxInt32 {
		return []int{}
	}
	result := []int{}
	for i := 0; i != -1; i = path[i] {
		result = append(result, i+1)
	}
	return result
}
