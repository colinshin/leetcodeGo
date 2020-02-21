/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

import "math"

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
如果定义买入算一笔交易而卖出不算
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i])      // 第i天不持有股票的情况有两种：前一天没有股票，今天不买； 或前一天有股票，今天卖了；选择收益最大的做法即可
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])   // 第i天持有股票的情况，与上边类似: 前一天有股票， 今天不卖，或前一天没有股票，今天买入； 买入则收益减少prices[i]

初始情况：
dp[-1][k][0] = 0
dp[i][0][0] = 0
dp[-1][k][1] = -infinity
dp[i][0][1] = -infinity
但注意k比较大的情况，实际k应该最大为数组长度的一半， 如果比数组长度的一半大, 用k无限大的解法，不然会使得dp数组太大

时空复杂度都是O(n*k)，其中n为数组长度
*/
func maxProfitKk3(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		return kNotLimitedMaxProfit(prices)
	}
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for t := range dp[i] {
			dp[i][t] = make([]int, 2)
			if i == 0 || t == 0 {
				dp[i][t][1] = math.MinInt32
			}
		}
	}
	for i := 1; i <= n; i++ {
		for t := 1; t <= k; t++ {
			dp[i][t][0] = max(dp[i-1][t][0], dp[i-1][t][1]+prices[i-1])
			dp[i][t][1] = max(dp[i-1][t][1], dp[i-1][t-1][0]-prices[i-1])
		}
	}
	return dp[n][k][0]
}

/*
如果定义买入后卖出算一笔交易：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k-1][1] + price[i])	// 第i天不持有股票的情况有两种：前一天没有股票，今天不买； 或前一天有股票，今天卖了；选择收益最大的做法即可
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])		// 第i天持有股票的情况，与上边类似: 前一天有股票， 今天不卖，或前一天没有股票，今天买入； 买入则收益减少prices[i]
初始情况：
dp[-1][k][0] = 0
dp[i][0][0] = 0
dp[-1][k][1] = -infinity
对于k为0的情况：
dp[i][0][1] = max(dp[i-1][0][1], -prices[i-1]) (特别地，dp[-1][0][1] = -prices[0])
*/
func maxProfitKk0(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k > n/2 {
		return kNotLimitedMaxProfit(prices)
	}
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for t := range dp[i] {
			dp[i][t] = make([]int, 2)
			if i == 0 {
				dp[i][t][1] = math.MinInt32
			}
			if t == 0 {
				if i == 0 {
					dp[i][t][1] = -prices[0]
				} else {
					dp[i][t][1] = max(dp[i-1][t][1], -prices[i-1])
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		for t := 1; t <= k; t++ {
			dp[i][t][0] = max(dp[i-1][t][0], dp[i-1][t-1][1]+prices[i-1])
			dp[i][t][1] = max(dp[i-1][t][1], dp[i-1][t][0]-prices[i-1])
		}
	}
	return dp[n][k][0]
}

/*
空间优化：用两个数组记录当天每次交易后持有股票或不持有股票的利润，即手里的钱
时间复杂度为O(n*k)，空间复杂度O(k)
*/
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}
	if k >= n/2 {
		return kNotLimitedMaxProfit(prices)
	}
	hold := make([]int, k+1)
	release := make([]int, k+1)
	for i := range hold {
		hold[i] = math.MinInt32
	}
	for _, price := range prices {
		for j := 1; j <= k; j++ { // j=0没有意义
			hold[j] = max(hold[j], release[j-1]-price)
			release[j] = max(release[j], hold[j]+price) // 要卖出必须先持有才行
		}
	}
	return release[k]
}
