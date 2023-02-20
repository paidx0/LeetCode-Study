/*
 * @lc app=leetcode.cn id=1575 lang=java
 *
 * [1575] 统计所有可行路径
 */

// @lc code=start
class Solution {
    public int countRoutes(int[] locations, int start, int finish, int fuel) {
        int n = locations.length;
        int[][] dp = new int[fuel + 1][n];
        int mod = 1000000007;
        dp[0][start] = 1;
        for(int i = 0; i <= fuel; i++){
            for(int j = 0; j < n; j++){
                for(int k = 0; k < n; k++){
                    if(j == k) continue;
                    int dis = Math.abs(locations[j] - locations[k]);
                    if(dis + i <= fuel){
                        dp[dis + i][k] = (dp[dis + i][k] + dp[i][j]) % mod;
                    }
                }
            }
        }
        int ans = 0;
        for (int i = 0; i <= fuel; ++i) {
            ans += dp[i][finish];
            ans %= mod;
        }
        return ans;
    }
}
// @lc code=end

