/*
 * @lc app=leetcode.cn id=1301 lang=java
 *
 * [1301] 最大得分的路径数目
 */

// @lc code=start
class Solution {
    public int[] pathsWithMaxScore(List<String> board) {
        int n = board.size();
        int[][][] dp = new int[n][n][2];
        int[] ans = new int[2];
        dp[0][0][0] = 0;
        dp[0][0][1] = 0;
        dp[n - 1][n - 1][0] = 0;
        dp[n - 1][n - 1][1] = 1;
        for(int i = n - 2; i >= 0; i--){
            char preScore = board.get(n - 1).charAt(i + 1);
            int score = board.get(n - 1).charAt(i) - 48;
            if(preScore != 'X'){
                if(preScore == 'E') dp[n - 1][i][0] = dp[n - 1][i + 1][0];
                else dp[n - 1][i][0] = dp[n - 1][i + 1][0] + score;
                dp[n - 1][i][1] += dp[n - 1][i + 1][1];
            }
            preScore = board.get(i + 1).charAt(n - 1);
            score = board.get(i).charAt(n - 1) - 48;
            if(preScore != 'X'){
                if(preScore == 'E') dp[i][n - 1][0] = dp[i + 1][n - 1][0];
                else dp[i][n - 1][0] = dp[i + 1][n - 1][0] + score;
                dp[i][n - 1][1] += dp[i + 1][n - 1][1];
            }
        }
        for(int i = n - 2; i >= 0; i--){
            for(int j = n - 2; j >= 0; j--){
                int num1 = board.get(i).charAt(j + 1);
                int num2 = board.get(i + 1).charAt(j + 1);
                int num3 = board.get(i + 1).charAt(j);
                int score = board.get(i).charAt(j) - 48;
                score = (score == 'E' - 48 ? 0 : score);
                int maxNum = max(num1 != 'X' ? dp[i][j + 1][0] : 0,
                                 num2 != 'X' && num2 != 'E' ? dp[i + 1][j + 1][0] : 0,
                                 num3 != 'X' ? dp[i + 1][j][0] : 0);
                if(maxNum != 0 && dp[i][j + 1][0] == maxNum){
                    dp[i][j][0] = dp[i][j + 1][0] + score;
                    dp[i][j][1] += dp[i][j + 1][1];
                }
                if((maxNum == 0 && num2 == 'S') || (maxNum != 0 && dp[i + 1][j + 1][0] == maxNum)){
                    dp[i][j][0] = dp[i + 1][j + 1][0] + score;
                    dp[i][j][1] += dp[i + 1][j + 1][1];
                }
                if(maxNum != 0 && dp[i + 1][j][0] == maxNum){
                    dp[i][j][0] = dp[i + 1][j][0] + score;
                    dp[i][j][1] += dp[i + 1][j][1];
                }
                dp[i][j][1] %=(int)1e9 + 7;
            }
        }
        return dp[0][0];
    }
    int max(int a, int b, int c){
        if(a > b) return Math.max(a, c);
        else return Math.max(b, c);
    }
}
// @lc code=end

