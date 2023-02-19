/*
 * @lc app=leetcode.cn id=304 lang=java
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
class NumMatrix {
    int[][] preSumMatrix;
    public NumMatrix(int[][] matrix) {
        preSumMatrix = new int[matrix.length][matrix[0].length];
        for(int i = 0; i < matrix.length; i++){
            for(int j = 0; j < matrix[0].length; j++){
                preSumMatrix[i][j] = matrix[i][j];
                if(i > 0) preSumMatrix[i][j] += preSumMatrix[i - 1][j];
                if(j > 0) preSumMatrix[i][j] += preSumMatrix[i][j - 1];
                if(i > 0 && j > 0) preSumMatrix[i][j] -= preSumMatrix[i - 1][j - 1];
            }
        }
    }
    
    public int sumRegion(int row1, int col1, int row2, int col2) {
        int res = preSumMatrix[row2][col2];
        if(col1 > 0)
            res -= preSumMatrix[row2][col1 - 1];
        if(row1 > 0)
            res -= preSumMatrix[row1 - 1][col2];
        if(row1 > 0 && col1 > 0)
            res += preSumMatrix[row1 - 1][col1 - 1];
        return res;
    }
}
/**
 * Your NumMatrix object will be instantiated and called as such:
 * NumMatrix obj = new NumMatrix(matrix);
 * int param_1 = obj.sumRegion(row1,col1,row2,col2);
 */
// @lc code=end

