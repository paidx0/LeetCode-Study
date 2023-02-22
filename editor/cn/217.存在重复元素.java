import java.util.HashSet;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=217 lang=java
 *
 * [217] 存在重复元素
 */

// @lc code=start
class Solution {
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> numSet = new HashSet<>();
        for(int num : nums){
            if(numSet.contains(num)) return true;
            numSet.add(num);
        }
        return false;
    }
}
// @lc code=end

