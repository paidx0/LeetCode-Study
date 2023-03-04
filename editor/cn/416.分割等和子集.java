/*
 * @lc app=leetcode.cn id=416 lang=java
 *
 * [416] 分割等和子集
 */

// @lc code=start
class Solution {
    public boolean canPartition(int[] nums) {
        int sum = 0;
        for (int i : nums) {
            sum += i;
        }
        if(sum % 2 != 0) return false;
        int n = nums.length;
        int v = sum / 2;
        int[][] dp = new int[n][v];
        dp[0][];

    }
}
// @lc code=end

