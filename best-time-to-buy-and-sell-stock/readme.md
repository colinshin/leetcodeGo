# 股票收益问题

参考：[一个通用的方法团灭6道股票问题](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-wen/)

## dp穷举
定义dp[i][k][s]表示第i天，还有最多k次交易次数，持有或不持有股票所能得到的最大收益<br>
假设共n天，最大交易次数为m，则所有的状态共（n * m * 2）种
```
for 0 <= i < n {
    for 1 <= k <= m {
        for 0 <= s <= 1 {
            dp[i][k][s] = max(buy, sell, rest) // 买入、卖出或不买卖
```
最终的答案就是dp[n-1][m][0]，即最后一天且不持有股票<br>
注意到一直都是同一支股票，dp的递推关系为：
```
// 第i天不持有股票的情况有两种：前一天没有股票，今天不买卖； 或前一天有股票，今天卖了；选择收益最大的做法即可， 即
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i])
// 第i天持有股票的情况，与上边类似: 前一天没有股票，今天买入； 或前一天有股票， 今天不买卖
dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1]) // 买入则收益减少prices[i]
```
注意买入我们会将k-1，卖出则不会；改成卖出时k-1，而买入时不变也是可以的<br>
那么初始情况如何确定？
```
// i从0开始，-1代表还没开始
dp[-1][k][0] = 0 
dp[-1][k][1] = -infinity
// k从1开始，k==0意味着不允许交易
dp[i][0][0] = 0
dp[i][0][1] = -infinity
```
数组索引-1及-infinity如何表示？<br>
可以对dp数组优化为有限的变量，消除下表-1的问题； 用math.MinInt32表示-infinity