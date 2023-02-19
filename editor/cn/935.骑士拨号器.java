/*
 * @lc app=leetcode.cn id=935 lang=java
 *
 * [935] 骑士拨号器
 */

// @lc code=start
class Solution {
    public int knightDialer(int n) {
        long[][][] dp = new long[4][3][n];
        int mod = (int)1e9 + 7;
        for(int i = 0; i < 4; i++){
            for(int j = 0; j < 3; j++){
                if(i < 3 || j == 1){
                    dp[i][j][0] = 1;
                }
            }
        }
        for(int k = 1; k < n; k++){
            for(int i = 0; i < 4; i++){
                for(int j = 0; j < 3; j++){
                    if(i < 3 || j == 1){
                        if(j - 2 >= 0){
                            if(i - 1 >= 0){
                                dp[i][j][k] += dp[i - 1][j - 2][k - 1];
                            }
                            if(i + 1 <= 3){
                                dp[i][j][k] += dp[i + 1][j - 2][k - 1];
                            }
                        }
                        if(j + 2 <= 2){
                            if(i - 1 >= 0){
                                dp[i][j][k] += dp[i - 1][j + 2][k - 1];
                            }
                            if(i + 1 <= 3){
                                dp[i][j][k] += dp[i + 1][j + 2][k - 1];
                            }
                        }
                        if(i - 2 >= 0){
                            if(j - 1 >= 0){
                                dp[i][j][k] += dp[i - 2][j - 1][k - 1];
                            }
                            if(j + 1 <= 2){
                                dp[i][j][k] += dp[i - 2][j + 1][k - 1];
                            }
                        }
                        if(i + 2 <= 3){
                            if(j - 1 >= 0){
                                dp[i][j][k] += dp[i + 2][j - 1][k - 1];
                            }
                            if(j + 1 <= 2){
                                dp[i][j][k] += dp[i + 2][j + 1][k - 1];
                            }
                        }
                        dp[i][j][k] %= mod;
                    }
                }
            }
        }
        int sum = 0;
        for(int i = 0; i < 4; i++){
            for(int j = 0; j < 3; j++){
                if(i < 3 || j == 1){
                    sum += (int)dp[i][j][n - 1];
                    sum %= mod;
                }
            }
        }
        return sum;
    }
}
// @lc code=end

