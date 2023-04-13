package main

func maxAreaOfIsland(grid [][]int) int {
	dx := [4]int{0, 0, 1, -1}
	dy := [4]int{1, -1, 0, 0}
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	max := 0
	count := 0
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		visited[x][y] = true
		count += 1
		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && !visited[xx][yy] && grid[xx][yy] == 1 {
				dfs(xx, yy)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count = 0
				dfs(i, j)
				if count > max {
					max = count
				}
			}
		}
	}
	return max
}
