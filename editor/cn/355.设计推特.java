import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=355 lang=java
 *
 * [355] 设计推特
 */

// @lc code=start
class Twitter {
    Map<Integer, Integer> twMap;
    Map<Integer, Set<Integer>> flMap;
    List<Integer> twTime;
    public Twitter() {
        twMap = new HashMap<>();
        flMap = new HashMap<>();
        twTime = new ArrayList<>();
    }
    
    public void postTweet(int userId, int tweetId) {
        if(!flMap.containsKey(userId)){
            Set<Integer> flSet = new HashSet<>();
            flSet.add(userId);
            flMap.put(userId, flSet);
        }
        twMap.put(tweetId, userId);
        twTime.add(tweetId);
    }
    
    public List<Integer> getNewsFeed(int userId) {
        List<Integer> ans = new ArrayList<>();
        Set<Integer> followSet = flMap.get(userId);
        int time = twTime.size() - 1;
        while(ans.size() < 10 && time >= 0){
            int tweetId = twTime.get(time--);
            if(followSet.contains(twMap.get(tweetId)))
                ans.add(tweetId);
        }
        return ans;
    }
    
    public void follow(int followerId, int followeeId) {
        if(!flMap.containsKey(followerId)){
            Set<Integer> flSet = new HashSet<>();
            flSet.add(followerId);
            flMap.put(followerId, flSet);
        }
        Set<Integer> flSet = flMap.get(followerId);
        flSet.add(followeeId);
    }
    
    public void unfollow(int followerId, int followeeId) {
        Set<Integer> flSet = flMap.get(followerId);
        flSet.remove(followeeId);
    }
}

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * List<Integer> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */
// @lc code=end

