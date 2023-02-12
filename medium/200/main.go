package main

import "fmt"

func numIsLands(grid [][]byte) int {
	var m = len(grid)
	if m == 0 {
		return 0
	}
	var n = len(grid[0])
	var dx = [4]int{0, 0, 1, -1}
	var dy = [4]int{1, -1, 0, 0}
	var visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var count = 0
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && !visited[xx][yy] && grid[xx][yy] == '1' {
				dfs(xx, yy)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(i, j)
				count += 1
			}
		}
	}
	return count
}

func main() {
	fmt.Println("vim-go")
}
