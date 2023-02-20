/*
 * @lc app=leetcode.cn id=552 lang=java
 *
 * [552] 学生出勤记录 II
 */

// @lc code=start
class Solution {
    public int checkRecord(int n) {
        long[][][] dp = new long[n][3][6];
        int mod = 1000000007;
        dp[0][0][0] = 1;
        dp[0][1][4] = 1;
        dp[0][2][3] = 1;
        for(int i = 1; i < n; i++){
            dp[i][0][0] = (dp[i][0][0] 
                           + dp[i - 1][2][3]
                           + dp[i - 1][1][4]
                           + dp[i - 1][1][5]) % mod;
            dp[i][1][1] = (dp[i][1][1]
                           + dp[i - 1][0][0]
                           + dp[i - 1][2][0]) % mod;
            dp[i][1][2] = (dp[i][1][2]
                           + dp[i - 1][1][1]) % mod;
            dp[i][1][4] = (dp[i][1][4]
                           + dp[i - 1][2][3]) % mod;
            dp[i][1][5] = (dp[i][1][5]
                           + dp[i - 1][1][4]) % mod;
            dp[i][2][0] = (dp[i][2][0]
                           + dp[i - 1][2][0]
                           + dp[i - 1][0][0]
                           + dp[i - 1][1][1]
                           + dp[i - 1][1][2]) % mod;
            dp[i][2][3] = (dp[i][2][3]
                           + dp[i - 1][2][3]
                           + dp[i - 1][1][4]
                           + dp[i - 1][1][5]) % mod;
        }
        long ans = dp[n - 1][0][0] + dp[n - 1][1][1] 
                    + dp[n - 1][1][2] + dp[n - 1][1][4]
                    + dp[n - 1][1][5] + dp[n - 1][2][0]
                    + dp[n - 1][2][3];
        ans %= mod;
        return (int)ans;
        /*
         * P0: A = 1 L = 0
         * P1: A = 1 L = 1
         * P2: A = 1 L = 2
         * P3: A = 0 L = 0
         * P4: A = 0 L = 1
         * P5: A = 0 L = 2
         * S0: P0
         * S1: P1 P2 P4 P5
         * S2: P0 P3
         * S0P0 = S2P3 + S1P4 + S1P5
         * S1P1 = S0P0 + S2P0
         * S1P2 = S1P1
         * S1P4 = S2P3
         * S1P5 = S1P4
         * S2P0 = S2P0 + S0P0 + S1P1 + S1P2
         * S2P3 = S2P3 + S1P4 + S1P5
         */
        
    }
}
// @lc code=end

