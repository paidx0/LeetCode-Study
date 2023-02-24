/*
 * @lc app=leetcode.cn id=1342 lang=java
 *
 * [1342] 将数字变成 0 的操作次数
 */

// @lc code=start
class Solution {
    public int numberOfSteps(int num) {
        if(num == 0) return 0;
        num = num % 2 == 0 ? num / 2 : num - 1;
        return 1 + numberOfSteps(num);
    }
}
// @lc code=end

