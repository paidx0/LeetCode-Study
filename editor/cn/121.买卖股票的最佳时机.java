/*
 * @lc app=leetcode.cn id=121 lang=java
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
class Solution {
    public int maxProfit(int[] prices) {
        int[] dp = new int[prices.length];
        dp[0] = 0;
        int minPrice = prices[0];
        int max = 0;
        for(int i = 1; i < prices.length; i++){
            dp[i] = Math.max(dp[i-1], prices[i] - minPrice);
            minPrice = Math.min(minPrice, prices[i]);
            if(dp[i] > max) max = dp[i];
        }
        return max;
    }
}
// @lc code=end

