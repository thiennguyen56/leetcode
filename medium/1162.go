package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	x int
	y int
}

var (
	dx = [4]int{1, -1, 0, 0}
	dy = [4]int{0, 0, 1, -1}
)

func BFS(queue *list.List, n int, dist [100][100]int) int {
	max := 0
	for queue.Len() > 0 {
		curAny := queue.Front()
		queue.Remove(curAny)
		cur := curAny.Value.(Point)
		x := cur.x
		y := cur.y
		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < n && dist[xx][yy] == -1 {
				dist[xx][yy] = dist[x][y] + 1
				if dist[xx][yy] > max {
					max = dist[xx][yy]
				}
				queue.PushBack(Point{
					x: xx,
					y: yy,
				})
			}
		}
	}
	return max
}
func maxDistance(grid [][]int) int {

	if len(grid) == 0 {
		return -1
	}
	n := len(grid)
	queue := list.New()
	dist := [100][100]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dist[i][j] = -1
			} else {
				dist[i][j] = 0
				queue.PushBack(Point{
					x: i,
					y: j,
				})
			}
		}
	}
	res := BFS(queue, n, dist)
	return res
}

func main() {
	fmt.Println(maxDistance([][]int{
		[]int{1, 0, 1},
		[]int{0, 0, 0},
		[]int{1, 0, 1},
	}))
	fmt.Println("vim-go")
}
