/*
 * @lc app=leetcode.cn id=1206 lang=java
 *
 * [1206] 设计跳表
 */

// @lc code=start
class Skiplist {
    private static final double P = 0.25;
    private static final int MAX_LEVEL = 32;
    private Node head = new Node(-1, MAX_LEVEL);
    public Skiplist() {}
    public boolean search(int target) {
        Node search = head;
        for (int i = MAX_LEVEL - 1; i >= 0; i--) {// 从最高等级找最接近自己的节点
            while (search.next[i] != null && target >= search.next[i].val) search = search.next[i];
            if (search.val == target) return true;
        }
        return false;
    }
    public void add(int num) {
        int level = this.randomLevel(); // 随机等级
        Node node = new Node(num, level);// 创建节点
        Node add = head;
        for (int i = MAX_LEVEL - 1; i >= 0; i--) { // 从最高等级开始向下寻找 找到属于自己的等级时添加
            while (add.next[i] != null && num > add.next[i].val) add = add.next[i];
            if (i >= level) continue;
            node.next[i] = add.next[i];
            add.next[i] = node;
        }
    }
    public boolean erase(int num) {
        boolean flag = false;
        Node erase = head;
        for (int i = MAX_LEVEL - 1; i >= 0; i--) {// 从最高等级找 小于自己的节点
            while (erase.next[i] != null && num > erase.next[i].val) erase = erase.next[i];
            if (erase.next[i] == null || erase.next[i].val != num) continue;
            erase.next[i] = erase.next[i].next[i];
            flag = true;
        }
        return flag;
    }
    private int randomLevel() {
        int level = 1;
        while (Math.random() < P && level < MAX_LEVEL) level++;
        return level;
    }
    private static class Node {
        private final int val;// 存储的值
        private final Node[] next;// 每个等级的下一个节点
        public Node(int val, int level) {
            this.val = val;
            this.next = new Node[level];
        }
    }
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * Skiplist obj = new Skiplist();
 * boolean param_1 = obj.search(target);
 * obj.add(num);
 * boolean param_3 = obj.erase(num);
 */
// @lc code=end

