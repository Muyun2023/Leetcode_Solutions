func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    
    dp := make([][]int, m+1)
    
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    
    dp[0][1] = 1

    return dp[m][n];
}
