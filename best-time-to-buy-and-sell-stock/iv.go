/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
动态规划：详见readme.md
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1])

但注意k比较大的情况，实际k应该最大为数组长度的一半， 如果比数组长度大， 用k无限大的解法，不然会使得dp数组太大
*/
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		return maxProfitKInfinity(prices)
	}
	// k为1或2的情况也可以不必单独解决，而是同下边的动态规划，这里直接引用会大大减少空间复杂度
	if k == 1 {
		return maxProfitK1(prices)
	}
	if k == 2 {
		return maxProfitK2(prices)
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 1; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 1; i <= k; i++ {
		dp[0][i] = []int{0, -prices[0]}
	}
	for i := 1; i < n; i++ {
		for ki := k; ki >= 1; ki-- {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k-1][0]-prices[i], dp[i-1][k][1])
		}
	}
	return dp[n-1][k][0]
}
