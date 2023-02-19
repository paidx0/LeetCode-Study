/*
 * @lc app=leetcode.cn id=790 lang=java
 *
 * [790] 多米诺和托米诺平铺
 */

// @lc code=start
class Solution {
    public int numTilings(int n) {
        // int[] ans = new int[n < 3 ? 3 : n];
        // ans[0] = 1;
        // ans[1] = 2;
        // ans[2] = 5;
        // for(int i = 3; i < n; i++){
        //     long tmp = (long)ans[i - 1] * 2 + ans[i - 3];
        //     tmp %= (1e9 + 7);
        //     ans[i] = (int)tmp;
        // }
        // return ans[n - 1];
        int[][] dp = new int[n + 1][4];
        int mod = (int)(1e9 + 7);
        dp[0][0] = 0;
        dp[0][1] = 0;
        dp[0][2] = 0;
        dp[0][3] = 1;
        for(int i = 1; i <= n; i++){
            dp[i][0] = dp[i - 1][3];
            dp[i][1] = dp[i - 1][0] + dp[i - 1][2] % mod;
            dp[i][2] = dp[i - 1][0] + dp[i - 1][1] % mod;
            dp[i][3] = (int)(((long)dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3]) % mod);
        }
        return dp[n][3];
    }
}
// @lc code=end

