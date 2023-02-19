/*
 * @lc app=leetcode.cn id=576 lang=java
 *
 * [576] 出界的路径数
 */

// @lc code=start
class Solution {
    public int findPaths(int m, int n, int maxMove, int startRow, int startColumn) {
        if(maxMove == 0) return 0;
        int[][][] dp = new int[m][n][maxMove];
        for(int i = 0; i < Math.max(m, n); i++){
            if(i < m){
                dp[i][0][0] += 1;
                dp[i][n - 1][0] += 1;
            }
            if(i < n){
                dp[0][i][0] += 1;
                dp[m - 1][i][0] += 1;
            }
        }
        for(int k = 1; k < maxMove; k++){
            for(int i = 0; i < m; i++){
                for(int j = 0; j < n; j++){
                    dp[i][j][k] += (i == 0 ? 1 : dp[i - 1][j][k - 1]);
                    dp[i][j][k] %= 1e9 + 7;
                    dp[i][j][k] += (j == 0 ? 1 : dp[i][j - 1][k - 1]);
                    dp[i][j][k] %= 1e9 + 7;
                    dp[i][j][k] += (i == m - 1 ? 1 : dp[i + 1][j][k - 1]);
                    dp[i][j][k] %= 1e9 + 7;
                    dp[i][j][k] += (j == n - 1 ? 1 : dp[i][j + 1][k - 1]);
                    dp[i][j][k] %= 1e9 + 7;
                }
            }
        }
        //return dp[1][1][1];
        return dp[startRow][startColumn][maxMove - 1];
    }
}
// @lc code=end

