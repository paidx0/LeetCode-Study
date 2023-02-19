/*
 * @lc app=leetcode.cn id=1223 lang=java
 *
 * [1223] 掷骰子模拟
 */

// @lc code=start
class Solution {
    public int dieSimulator(int n, int[] rollMax) {
        int[][][] dp  = new int[n][6][16];
        for(int i = 0; i < 6; i++){
            dp[0][i][1] = 1;
        }
        for(int i = 1; i < n; i++){
            for(int j = 0; j < 6; j++){
                
            }
        }
        return 0;
    }
}
// @lc code=end

