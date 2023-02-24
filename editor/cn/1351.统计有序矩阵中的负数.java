/*
 * @lc app=leetcode.cn id=1351 lang=java
 *
 * [1351] 统计有序矩阵中的负数
 */

// @lc code=start
class Solution {
    public int countNegatives(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int count = 0;
        for(int i = 0; i < m; i++){
            for(int j = n - 1; j >= 0; j--){
                if(grid[i][j] < 0) count++;
                else break;
            }
        }
        return count;
    }
}
// @lc code=end

