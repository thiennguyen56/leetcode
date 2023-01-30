class Solution:
    n = 0
    m = 0
    dx = [0, 0, 1, -1]
    dy = [1, -1, 0, 0]
    visited = []
    count = 0
    grid = []
    def DFS(self, start):
        self.visited[start["x"]][start["y"]] = True
        for i in range(4):
            xx = start["x"] + self.dx[i]
            yy = start["y"] + self.dy[i]
            if (xx >= 0 and xx < self.m and yy >= 0 and yy < self.n and not self.visited[xx][yy] and self.grid[xx][yy] == '1'):
                self.DFS({
                    "x": xx, 
                    "y": yy
                })
    def numIslands(self, grid: List[List[str]]) -> int:
        self.grid = grid
        self.m = len(grid)
        if self.m == 0:
            return 0

        self.n = len(grid[0])
        self.visited = [[False for _ in range(self.n)] for _ in range(self.m)]
        self.count = 0
        for i in range(self.m):
            for j in range(self.n):
                if (self.grid[i][j] == "1" and self.visited[i][j] == False):
                    self.DFS({
                        "x": i,
                        "y": j
                    })
                    self.count += 1
        return self.count

