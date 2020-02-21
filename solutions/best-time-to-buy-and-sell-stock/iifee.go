/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

import "math"

/*
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每次交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

示例 1:

输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
输出: 8
解释: 能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
注意:

0 < prices.length <= 50000.
0 < prices[i] < 50000.
0 <= fee < 50000.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
动态规划，详见readme.md
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee
*/
func kNotLimitedMaxProfitWithFee(prices []int, fee int) int {
	noStockProfit, hasStockProfit := 0, math.MinInt32
	for _, price := range prices {
		noStockProfit = max(noStockProfit, hasStockProfit+price)
		hasStockProfit = max(hasStockProfit, noStockProfit-price-fee)
	}
	return noStockProfit
}
