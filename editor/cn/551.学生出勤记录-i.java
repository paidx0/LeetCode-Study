/*
 * @lc app=leetcode.cn id=551 lang=java
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
class Solution {
    public boolean checkRecord(String s) {
        int absent = 0, late = 0;
        int i, j;
        for(i = 0; i < s.length(); i++){
            if(s.charAt(i) == 'A') absent += 1;
            if(s.charAt(i) == 'L') late += 1;
            else late = 0;
            if(absent == 2 || late == 3) return false;
        }
        return true;
    }
}
// @lc code=end

