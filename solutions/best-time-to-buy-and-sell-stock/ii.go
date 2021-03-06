/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

import "math"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
动态规划，详见readme.md；这个问题中k无限大，可以不考虑k的影响
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][0] - prices[i], dp[i-1][1])

时间复杂度O(n), 空间复杂度O(1)
*/
func kNotLimitedMaxProfit(prices []int) int {
	noStockProfit, hasStockProfit := 0, math.MinInt32
	for _, v := range prices {
		noStockProfit = max(noStockProfit, hasStockProfit+v)
		hasStockProfit = max(noStockProfit-v, hasStockProfit)
	}
	return noStockProfit
}

/*
实际就是只要当前价格高于前一天，就可以在前一天卖出，今天买入
时间复杂度O(n), 空间复杂度O(1)
*/
func maxProfit21(prices []int) int {
	n := len(prices)
	maxProfit := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
