/*
 * @lc app=leetcode.cn id=53 lang=java
 *
 * [53] 最大子数组和
 */

// @lc code=start
class Solution {
    public int maxSubArray(int[] nums) {
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        int max = dp[0];
        for(int i = 1; i < nums.length; i++){
            if(dp[i - 1] < 0){
                dp[i] = nums[i];
            }else{
                dp[i] = nums[i] + dp[i - 1];
            }
            max = Math.max(max, dp[i]);
        }
        return max;
    }
}
// @lc code=end

