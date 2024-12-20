func pacificAtlantic(heights [][]int) [][]int {
    ROWS, COLS := len(heights), len(heights[0])
    pac, atl := make(map[int]bool), make(map[int] bool)
    
    var dfs func(int, int, map[int]bool, int)
    dfs = func(r, c int, visit map[int]bool, prevHeight int) {
        if (
            visit[r*COLS + c] ||
            r < 0 ||
            c < 0 ||
            r == ROWS ||
            c == COLS ||
            heights[r][c] < prevHeight) {
            return;
        }
        visit[r*COLS + c] = true
        dfs(r + 1, c, visit, heights[r][c])
        dfs(r - 1, c, visit, heights[r][c])
        dfs(r, c + 1, visit, heights[r][c])
        dfs(r, c - 1, visit, heights[r][c])
    }
}
