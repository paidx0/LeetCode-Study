/*
 * @lc app=leetcode.cn id=1155 lang=java
 *
 * [1155] 掷骰子的N种方法
 */

// @lc code=start
class Solution {
    // public int numRollsToTarget(int n, int k, int target) {
    //     if(n * k < target) return 0;
    //     long[][] dp = new long[n + 1][n * k + 1];
    //     int mod = (int)1e9 + 7;
    //     for(int i = 1; i <= k; i++){
    //         dp[1][i] = 1;
    //     }
    //     for(int i = 1; i < n; i++){
    //         for(int j = i; j <= i * k; j++){
    //             for(int l = 1; l <= k; l++){
    //                 dp[i + 1][j + l] += dp[i][j];
    //                 dp[i + 1][j + l] %= mod;
    //             }
    //         }
    //     }
    //     return (int) dp[n][target] % mod;
    // }
    public int numRollsToTarget(int n, int k, int target) {
        if(n * k < target || n > target) return 0;
        long[] dp_pre = new long[n * k + 1];
        long[] dp = new long[n * k + 1];
        int mod = (int)1e9 + 7;
        for(int i = 1; i <= k; i++){
            dp_pre[i] = 1;
        }
        for(int i = 1; i < n; i++){
            for(int j = i; j <= i * k; j++){
                for(int l = 1; l <= k; l++){
                    dp[j + l] += dp_pre[j];
                    dp[j + l] %= mod;
                }
            }
            dp_pre = dp;
            dp = new long[n * k + 1];
        }
        return (int) dp_pre[target] % mod;
    }
}
// @lc code=end

