/*
 * @lc app=leetcode.cn id=120 lang=java
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        int n = triangle.size();
        int[][] dp = new int[n][n];
        dp[0][0] = triangle.get(0).get(0);
        for(int i = 1; i < n; i++){
            dp[i][0] = triangle.get(i).get(0) + dp[i - 1][0];
            dp[i][i] = triangle.get(i).get(i) + dp[i - 1][i - 1];
        }
        for(int i = 1; i < n; i++){
            for(int j = 1; j < i; j++){
                dp[i][j] = Math.min(dp[i - 1][j], dp[i - 1][j - 1]) + triangle.get(i).get(j);
            }
        }
        int ans = dp[n - 1][0];
        for(int i = 1; i < n; i++){
            ans = Math.min(ans, dp[n - 1][i]);
        }
        return ans;
    }
}
// @lc code=end

