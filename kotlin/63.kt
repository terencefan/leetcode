package leetcode.p63

class Solution {
    fun uniquePathsWithObstacles(obstacleGrid: Array<IntArray>): Int {
        if (obstacleGrid.isEmpty() || obstacleGrid[0].isEmpty() || obstacleGrid[0][0] == 1) return 0

        var m = obstacleGrid.size
        var n = obstacleGrid[0].size

        obstacleGrid[0][0] = 1
        for (i in 0 until m) {
            for (j in 0 until n) {
                if (i == 0 && j == 0) continue
                if (obstacleGrid[i][j] == 1) {
                    obstacleGrid[i][j] = 0
                    continue
                } 
                if (i > 0) {
                    obstacleGrid[i][j] += obstacleGrid[i - 1][j]
                }
                if (j > 0) {
                    obstacleGrid[i][j] += obstacleGrid[i][j - 1]
                }
            }
        }
        return obstacleGrid[m - 1][n - 1]
    }
}