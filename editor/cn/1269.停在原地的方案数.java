/*
 * @lc app=leetcode.cn id=1269 lang=java
 *
 * [1269] 停在原地的方案数
 */

// @lc code=start
class Solution {
    public int numWays(int steps, int arrLen) {
        arrLen = Math.min(steps, arrLen); //优化arrLen, 不分配一定无法到达的空间
        int[][] dp = new int[2][arrLen];
        int mod = (int)1e9 + 7;
        dp[0][0] = 1;
        for(int i = 1; i <= steps; i++){
            for(int j = 0; j < arrLen; j++){
                dp[1][j] += j == 0 ? 0 : dp[0][j - 1];
                dp[1][j] %= mod;
                dp[1][j] += dp[0][j];
                dp[1][j] %= mod;
                dp[1][j] += j == arrLen - 1 ? 0 : dp[0][j + 1];
                dp[1][j] %= mod;
            }
            dp[0] = dp[1];
            dp[1] = new int[arrLen];
        }
        return dp[0][0];
    }
}
// @lc code=end

