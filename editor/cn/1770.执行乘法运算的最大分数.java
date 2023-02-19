import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;
import java.util.stream.Collectors;

/*
 * @lc app=leetcode.cn id=1770 lang=java
 *
 * [1770] 执行乘法运算的最大分数
 */

// @lc code=start
class Solution {
    public int maximumScore(int[] nums, int[] multipliers) {
        int m = multipliers.length;
        int n = nums.length;
        int[][] dp = new int[m + 1][m + 1];
        Deque<Integer> que = new ArrayDeque(Arrays.stream(nums).boxed().collect(Collectors.toList()));
        dp[0][0] = 0;
        for(int i = 1; i <= m; i++){
            dp[i][0] = dp[i - 1][0] + nums[i - 1] * multipliers[i - 1];
            dp[0][i] = dp[0][i - 1] + nums[n - i] * multipliers[i - 1];
        }
        for(int i = 1; i <= m; i++){
            for(int j = 1; i + j <= m; j++){
                int tmp1 = dp[i - 1][j] + nums[i - 1] * multipliers[i + j - 1];
                int tmp2 = dp[i][j - 1] + nums[n - j] * multipliers[i + j - 1];
                dp[i][j] = Math.max(tmp1, tmp2);
            }
        }
        int max = dp[0][m];
        for(int i = 0; i <= m; i++){
            if(dp[i][m - i] > max) max = dp[i][m - i];
        }
        return max;
    }
}
// @lc code=end

