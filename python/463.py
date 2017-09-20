#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):

    def findLand(self, grid, m, n):
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    return i, j

    def islandPerimeter(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(grid), len(grid[0])
        for i in range(m):
            grid[i].append(0)
            grid[i] = [0] + grid[i]
        grid.append([0] * (n + 2))
        grid = [[0] * (n + 2)] + grid

        x, y = self.findLand(grid, m + 2, n + 2)

        grid[x][y] == 2
        points, s = [(x, y)], 0

        while points:
            x, y = points.pop()
            for i, j in [(x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1)]:
                if grid[i][j] == 0:
                    s += 1
                elif grid[i][j] == 1:
                    points.append((i, j))
                    grid[x][y] = 2
        return s


if __name__ == '__main__':
    r = Solution().islandPerimeter([
        [1, 1],
        [1, 1],
    ])
    print(r)
