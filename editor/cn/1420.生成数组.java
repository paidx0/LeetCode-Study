/*
 * @lc app=leetcode.cn id=1420 lang=java
 *
 * [1420] 生成数组
 */

// @lc code=start
class Solution {
    public int numOfArrays(int n, int m, int k) {
        if(k == 0) return 0;
        int[][][] dp = new int[n + 1][k + 1][m + 1];
        int mod = (int)1e9 + 7;
        for(int i = 1; i <= m; i++) dp[1][1][i] = 1;
        for(int i = 2; i <= n; i++)
            for(int j = 1; j <= k; j++)
                for(int l = 1; l <= m; l++){
                    dp[i][j][l] = (int)(((long)dp[i - 1][j][l] * l) % mod);
                    for(int a = 1; a < l; a++) dp[i][j][l] = (dp[i][j][l] + dp[i - 1][j - 1][a]) % mod;
                }
        long sum = 0;
        for(int i = 1; i <= m; i++) sum = (sum + dp[n][k][i]) % mod;
        return (int)sum % mod;
    }
}
// @lc code=end

