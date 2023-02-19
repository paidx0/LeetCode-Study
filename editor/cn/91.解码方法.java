/*
 * @lc app=leetcode.cn id=91 lang=java
 *
 * [91] 解码方法
 */

// @lc code=start
class Solution {
    public int numDecodings(String s) {
        int n = s.length();
        int[] dp = new int[n];
        if(s.charAt(0) != '0') dp[0] = 1;
        else return 0;
        if(n == 1) return 1;
        char tmp0 = s.charAt(0);
        char tmp1 = s.charAt(1);
        if(tmp0 >= '3') dp[1] = tmp1 > '0' ? 1 : 0;
        else if(tmp0 == '2' && tmp1 > '6') dp[1] = 1;
        else if(tmp1 == '0') dp[1] = 1;
        else dp[1] = 2;
        for(int i = 2; i < n; i++){
            if(dp[i - 1] == 0) continue;
            char preNum = s.charAt(i - 1), num = s.charAt(i);
            if(num != '0') dp[i] += dp[i - 1];
            if(preNum == '1' || (preNum =='2' && num <= '6'))
                dp[i] += dp[i - 2];
        }
        return dp[n - 1];
    }
}
// @lc code=end

