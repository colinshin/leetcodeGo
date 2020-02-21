/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

import "math"

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [3,3,5,0,0,3,1,4]
输出: 6
解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
动态规划，详见readme.md

如果定义买入算一笔交易而卖出不算
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i])      // 第i天不持有股票的情况有两种：前一天没有股票，今天不买； 或前一天有股票，今天卖了；选择收益最大的做法即可
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])   // 第i天持有股票的情况，与上边类似: 前一天有股票， 今天不卖，或前一天没有股票，今天买入； 买入则收益减少prices[i]
在这里k只有1或2两种情况
dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], dp[i][0][0] - prices[i])
			= max(dp[i-1][1][1], -prices[i])
*/
func maxProfitK2(prices []int) int {
	// 定义买入算一笔交易而卖出不算
	noStockDone2Profit, hasStockDone2Profit := 0, math.MinInt32
	noStockDone1Profit, hasStockDone1Profit := 0, math.MinInt32
	for _, price := range prices {
		// 当天不持有股票： 要么前一天也不持有，要么前一天持有，当天卖了
		// 当天持有股票： 要么前一天持有，要么前一天不持有，当天买了
		noStockDone2Profit = max(noStockDone2Profit, hasStockDone2Profit+price)
		hasStockDone2Profit = max(hasStockDone2Profit, noStockDone1Profit-price)
		noStockDone1Profit = max(noStockDone1Profit, hasStockDone1Profit+price)
		hasStockDone1Profit = max(hasStockDone1Profit, -price)
	}
	return noStockDone2Profit
}

/*
或者这样理解:

只要知道前一个时间点买卖第一第二笔股票的最大收益信息，就可以算出当前最大的收益了，这样可以省去额外空间。
遍历prices数组的时候，按照持有股票或不持有股票，维护四个变量:
hold1是在该价格点买入第一笔股票后手里剩的钱
release1是在该价格点卖出第一笔股票后手里剩的钱
hold2是在该价格点买入第二笔股票后手里剩的钱
release2是在该价格点卖出第二笔股票后手里剩的钱
因为卖是要后于买的，而第二次交易也是后于第一次交易的，为了用这些变量自身来记录上次的值，计算顺序为release2 -> hold2 -> release1 -> hold1
或者用go的多值交换技巧来计算，就不用太考虑顺序
*/
func maxProfitK20(prices []int) int {
	hold1, hold2 := math.MinInt32, math.MinInt32
	release1, release2 := 0, 0
	for _, price := range prices {
		release2 = max(release2, hold2+price)
		hold2 = max(hold2, release1-price)
		release1 = max(release1, hold1+price)
		hold1 = max(hold1, -price)
	}
	return release2
}
