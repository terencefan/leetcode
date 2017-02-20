#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):

    def solve(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        if not board:
            return

        m = len(board)
        n = len(board[0])

        def _fill(i, j):
            s = [(i, j)]
            while s:
                x, y = s.pop()
                board[x][y] = 'F'
                if x > 0 and board[x - 1][y] == 'O':
                    s.append((x - 1, y))
                if x < m - 1 and board[x + 1][y] == 'O':
                    s.append((x + 1, y))
                if y > 0 and board[x][y - 1] == 'O':
                    s.append((x, y - 1))
                if y < n - 1 and board[x][y + 1] == 'O':
                    s.append((x, y + 1))

        def _reverse():
            for i in range(m):
                for j in range(n):
                    if board[i][j] == 'F':
                        board[i][j] = 'O'
                    elif board[i][j] == 'O':
                        board[i][j] = 'X'

        for i in range(m):
            for j in range(n):
                if board[i][j] != 'O':
                    continue
                if i == 0 or j == 0 or i == m - 1 or j == n - 1:
                    print(i, j)
                    _fill(i, j)
                    print(board)
        _reverse()
        return board


if __name__ == '__main__':
    board = [
        'OOOOXX',
        'OOOOOO',
        'OXOXOO',
        'OXOOXO',
        'OXOXOO',
        'OXOOOO',
    ]
    board = list(map(lambda x: list(x), board))
    res = Solution().solve(board)
    print(res)
