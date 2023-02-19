/*
 * @lc app=leetcode.cn id=1641 lang=java
 *
 * [1641] 统计字典序元音字符串的数目
 */

// @lc code=start
class Solution {
    public int countVowelStrings(int n) {
        int[][] dp  = new int[n][5];
        for(int i = 0; i < 5; i++){
            dp[0][i] = 1;
        }
        for(int i = 1; i < n; i++){
            dp[i][0] = dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4];
            dp[i][1] = dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4];
            dp[i][2] = dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4];
            dp[i][3] = dp[i - 1][3] + dp[i - 1][4];
            dp[i][4] = dp[i - 1][4];
        }
        return dp[n - 1][0] + dp[n - 1][1] + dp[n - 1][2] + dp[n - 1][3] + dp[n - 1][4];
    }
}
// @lc code=end

