/*
 * @lc app=leetcode.cn id=1220 lang=java
 *
 * [1220] 统计元音字母序列的数目
 */

// @lc code=start
class Solution {
    public int countVowelPermutation(int n) {
        long[][] dp = new long[n][5];
        int mod = 1000000007;
        for(int i = 0; i < 5; i++){
            dp[0][i] = 1;
        }
        for(int i = 1; i < n; i++){
            dp[i][0] = (dp[i - 1][1] + dp[i - 1][2] + dp[i-1][4]) % mod;
            dp[i][1] = (dp[i - 1][0] + dp[i - 1][2]) % mod;
            dp[i][2] = (dp[i - 1][1] + dp[i - 1][3]) % mod;
            dp[i][3] = (dp[i - 1][2]) % mod;
            dp[i][4] = (dp[i - 1][2] + dp[i - 1][3]) % mod;
        }
        long ans = 0;
        for(int i = 0; i < 5; i++){
            ans =  (ans + dp[n - 1][i]) % mod;
        }
        return (int)ans % mod;
    }
    /* 
        e -> a i -> a u -> a
        a -> e i -> e
        e -> i o -> i
        i -> o
        i -> u o -> u
    */
}
// @lc code=end

