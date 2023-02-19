/*
 * @lc app=leetcode.cn id=931 lang=java
 *
 * [931] 下降路径最小和
 */

// @lc code=start
class Solution {
    public int minFallingPathSum(int[][] matrix) {
        int n = matrix.length;
        int[][] dp = new int[n][n];
        for(int i = 0; i < n; i++)
            dp[0][i] = matrix[0][i];
        for(int i = 1; i < n; i++){
            dp[i][0] = Math.min(dp[i - 1][0], dp[i - 1][1]) + matrix[i][0];
            for(int j = 1; j < n - 1; j++){
                dp[i][j] = min(dp[i - 1][j - 1],
                                dp[i - 1][j],
                                dp[i - 1][j + 1]) + matrix[i][j];
            }
            dp[i][n - 1] = Math.min(dp[i - 1][n - 1], dp[i - 1][n - 2]) + matrix[i][n - 1];
        }
        int min = dp[n - 1][0];
        for(int i = 1; i < n; i++)
            min = Math.min(min, dp[n - 1][i]);
        return min;
    }
    public int min(int a, int b, int c){
        if(a < c) return Math.min(a, b);
        else return Math.min(b, c);
    }
}
// @lc code=end

