/*
 * @lc app=leetcode.cn id=264 lang=java
 *
 * [264] 丑数 II
 */

// @lc code=start
class Solution {
    public int nthUglyNumber(int n) {
        int[] dp = new int[n + 1];
        dp[1] = 1;
        int p2 = 1, p3 = 1, p5 = 1;
        for(int i = 2; i <= n; i++){
            dp[i] = min(dp[p2] * 2, dp[p3] * 3, dp[p5] * 5);
            if(dp[i] == 2 * dp[p2]) p2++;
            if(dp[i] == 3 * dp[p3]) p3++;
            if(dp[i] == 5 * dp[p5]) p5++;
        }
        return dp[n];
    }
    public int min(int a, int b, int c){
        if(Math.min(a, b) == a){
            return Math.min(a, c);
        }else{
            return Math.min(b, c);
        }
    }
}
// @lc code=end

