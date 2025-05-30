func isMatch(s string, p string) bool {
    dp := make([][]bool, len(p) + 1)
    for i := 0; i < len(p) + 1; i++ {
        dp[i] = make([]bool, len(s) + 1)
        for j := 0; j < len(s) + 1; j++ {
            dp[i][j] = false
        }
    }

    dp[0][0] = true
    for i := 1; i < len(p); i++ {
        if p[i] == '*' {
            dp[i + 1][0] = dp[i - 1][0]
        }
    }

    for i := 0; i < len(p); i++ {
        for j := 0; j < len(s); j++ {
            if p[i] == '.' || p[i] == s[j] {

                dp[i + 1][j + 1] = dp[i][j]
            } else if p[i] == '*' {
                dp[i + 1][j + 1] = dp[i - 1][j + 1] || dp[i][j + 1]
                if p[i - 1] == '.' || p[i - 1] == s[j] {
                    if dp[i + 1][j] {
                        dp[i + 1][j + 1] = dp[i + 1][j]
                    }
                }
            }
        }
    }

    return dp[len(p)][len(s)]
}
