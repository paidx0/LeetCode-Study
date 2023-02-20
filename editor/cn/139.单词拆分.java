import java.util.Set;

/*
 * @lc app=leetcode.cn id=139 lang=java
 *
 * [139] 单词拆分
 */

// @lc code=start
class Solution {
    Set<String> set = new HashSet<>();
    public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> set = new HashSet<>();
        for (String word : wordDict) set.add(word);
        int n = s.length();
        boolean[] f = new boolean[n + 1];
        f[0] = true;
        for(int i = 1; i <= n; i++){
            for(int j = 1; j <= i && !f[i]; j++){
                String sub = s.substring(j - 1, i);
                if(set.contains(sub)) f[i] = f[j - 1];
            }
        }
        return f[n];
    }
}
// @lc code=end

