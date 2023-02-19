/*
 * @lc app=leetcode.cn id=63 lang=java
 *
 * [63] 不同路径 II
 */

// @lc code=start
class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int m = obstacleGrid.length;
        int n = obstacleGrid[0].length;
        if(obstacleGrid[0][0] == 1 || obstacleGrid[m - 1][n - 1] == 1) return 0;
        int[][] dp = new int[m][n];
        dp[0][0] = 1;
        for(int i = 1; i < Math.max(m, n); i++){
            if(i < m && obstacleGrid[i][0] != 1) dp[i][0] = dp[i - 1][0];
            if(i < n && obstacleGrid[0][i] != 1) dp[0][i] = dp[0][i - 1];
        }
        for(int i = 1; i < m; i++){
            for(int j = 1; j < n; j++){
                if(obstacleGrid[i - 1][j] != 1)
                    dp[i][j] += dp[i - 1][j]; 
                if(obstacleGrid[i][j - 1] != 1)    
                    dp[i][j] += dp[i][j - 1];
            }
        }
        return dp[m - 1][n - 1];
    }
}
// @lc code=end

