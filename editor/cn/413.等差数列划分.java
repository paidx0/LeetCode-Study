/*
 * @lc app=leetcode.cn id=413 lang=java
 *
 * [413] 等差数列划分
 */

// @lc code=start
class Solution {
    public int numberOfArithmeticSlices(int[] nums) {
        int n = nums.length;
        if(n < 3) return 0;
        int[] dp = new int[n];
        dp[0] = dp[1] = 0;
        int sum = 0;
        for(int i = 2; i < n; i++){
            if(nums[i] - nums[i - 1] == nums[i - 1] - nums[i - 2]){
                dp[i] = dp[i - 1] + 1; 
            }else{
                dp[i] = 0;
            }
            sum += dp[i];
        }
        return sum;
    }
}
// @lc code=end

