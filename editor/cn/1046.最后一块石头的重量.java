/*
 * @lc app=leetcode.cn id=1046 lang=java
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>((a, b) -> b - a);
        for(int i : stones) maxHeap.offer(i);
        while(maxHeap.size() > 1){
            int a = maxHeap.poll(), b = maxHeap.poll();
            if(a > b) maxHeap.offer(a - b);
        }
        return maxHeap.size() == 0 ? 0 : maxHeap.poll();
    }
}
// @lc code=end

