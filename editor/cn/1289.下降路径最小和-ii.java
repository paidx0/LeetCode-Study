/*
 * @lc app=leetcode.cn id=1289 lang=java
 *
 * [1289] 下降路径最小和  II
 */

// @lc code=start
class Solution {
    int MAX = Integer.MAX_VALUE;
    public int minFallingPathSum(int[][] nums) {
        int n = nums.length;
        int[][] dp = new int[n][n];
        int min = -1, secMin = -1;
        for (int i = 0; i < n; i++) {
            int val = nums[0][i];
            dp[0][i] = val;
            // 更新 i1 和 i2
            if (val < (min == -1 ? MAX : dp[0][min])) {
                secMin = min;
                min = i;
            } else if (val < (secMin == -1 ? MAX : dp[0][secMin])) {
                secMin = i;
            }
        }
        for (int i = 1; i < n; i++) {
            int tmpMin = -1, tmpSecMin = -1;
            for (int j = 0; j < n; j++) {
                dp[i][j] = MAX;
                int val = nums[i][j];
                if (j != min) {
                    dp[i][j] = dp[i - 1][min] + val;
                } else {
                    dp[i][j] = dp[i - 1][secMin] + val;
                }
                if (dp[i][j] < (tmpMin == -1 ? MAX : dp[i][tmpMin])) {
                    tmpSecMin = tmpMin;
                    tmpMin = j;
                } else if (dp[i][j] < (tmpSecMin == -1 ? MAX : dp[i][tmpSecMin])) {
                    tmpSecMin = j;
                }
            }
            min = tmpMin; secMin = tmpSecMin;
        }
        return dp[n - 1][min];
    }
}
//[[-73,61,43,-48,-36],[3,30,27,57,10],[96,-76,84,59,-15],[5,-49,76,31,-7],[97,91,61,-46,67]]
// @lc code=end

