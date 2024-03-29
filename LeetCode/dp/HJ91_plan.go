package dp

func plan1(n, m int) int {
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n][m]
}
