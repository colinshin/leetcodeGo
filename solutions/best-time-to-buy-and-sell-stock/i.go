/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package best_time_to_buy_and_sell_stock

import "math"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
实际就是最高的波峰减去最低的波谷的值，但需注意符合题意的波峰必须在波谷右边出现
时间复杂度O(n), 空间复杂度O(1)
*/
func maxProfit10(prices []int) int {
	minPrice, maxProfit := math.MaxInt32, 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}
	return maxProfit
}

/*
动态规划：详见readme.md

dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][0][0] - prices[i], dp[i-1][1][1])
			= max(-prices[i], dp[i-1][1][1])
发现与k无关，化简
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(-prices[i], dp[i-1][1])
时间复杂度O(n), 空间复杂度O(2n)=O(n)
*/
func maxProfit11(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][]int, n)
	dp[0] = make([]int, 2)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

/*
实际上每次确定结果只和上一次的结果有关，用两个变量来替代dp数组来优化空间复杂度
时间复杂度O(n), 空间复杂度O(1)
*/
func maxProfit12(prices []int) int {
	noStockProfit, hasStockProfit := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		noStockProfit, hasStockProfit =
			max(noStockProfit, hasStockProfit+prices[i]), max(-prices[i], hasStockProfit)
	}
	return noStockProfit
}

//或者根据readme的说明写初始条件
func maxProfitK1(prices []int) int {
	noStockProfit, hasStockProfit := 0, math.MinInt32
	for _, v := range prices {
		noStockProfit, hasStockProfit =
			max(noStockProfit, hasStockProfit+v), max(-v, hasStockProfit)
	}
	return noStockProfit
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
