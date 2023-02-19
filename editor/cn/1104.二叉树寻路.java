/*
 * @lc app=leetcode.cn id=1104 lang=java
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
class Solution {
    public List<Integer> pathInZigZagTree(int label) {
        int count = 0;
        int tmp = label;
        while(tmp > 0){
            tmp /= 2;
            count++;
        }
        int base = (int) Math.pow(2, count - 1);
        if(count % 2 == 0)
            tmp = base * 3 - label - 1;
        else
            tmp = label;
        List<Integer> ans = new ArrayList();
        while(count > 0){
            if(count % 2 == 0){
                ans.add(0, base * 3 - tmp - 1);
            }else{
                ans.add(0, tmp);
            }
            tmp /= 2;
            base /= 2;
            count--;
        }
        return ans;
    }
}
// @lc code=end

