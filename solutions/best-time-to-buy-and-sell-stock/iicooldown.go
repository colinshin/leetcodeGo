/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

import "math"

/*
动态规划，详见readme.md；这个问题中k无限大，可以不考虑k的影响
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i]) // 因为冷冻期的原因，今天买入，则前天必须不持有股票

时间复杂度O(n), 空间复杂度O(1)
*/
func kNotLimitedMaxProfitWithCoolDown(prices []int) int {
	noStockProfit, hasStockProfit := 0, math.MinInt32
	lastNoStockProfit := noStockProfit
	for _, price := range prices {
		noStockProfit, hasStockProfit, lastNoStockProfit =
			max(noStockProfit, hasStockProfit+price),
			max(hasStockProfit, lastNoStockProfit-price),
			noStockProfit
	}
	return noStockProfit
}
