/*
 * @lc app=leetcode.cn id=140 lang=java
 *
 * [140] 单词拆分 II
 */

// @lc code=start
class Solution {
    Set<String> set = new HashSet<>();
    Map<Integer, List<String>> cache = new HashMap();
    String str;
    public List<String> wordBreak(String s, List<String> wordDict) {
        for(String word : wordDict) set.add(word);
        str = s;
        return dfs(0);
    }
    public List<String> dfs(int index){
        List<String> res = new ArrayList<>();
        if(cache.containsKey(index)){
            return cache.get(index);
        }
        if(index >= str.length()){
            res.add("");
            return res;
        }
        for(int i = index; i <= str.length(); i++){
            String subStr = str.substring(index, i);
            if(set.contains(subStr)){
                List<String> tmp = dfs(i);
                for(String word : tmp){
                    res.add(subStr + (word == "" ? "" : " ") + word);
                }
            }
        }
        cache.put(index, res);
        return res;
    }
}
// @lc code=end

