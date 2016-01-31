#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def exist(self, board, word):
        """
        :type board: List[List[str]]
        :type word: str
        :rtype: bool
        """
        if not word:
            return False

        ds = [
            (1, 0),
            (0, 1),
            (-1, 0),
            (0, -1),
        ]

        m, n = len(board), len(board[0])

        st = []
        for i in range(m):
            for j in range(n):
                if board[i][j] == word[0]:
                    st.append((1, i, j, set([i * n + j])))

        while st:
            pos, x, y, steps = st.pop()
            print pos, x, y, steps
            if pos == len(word):
                return True
            for d in ds:
                x1, y1 = x + d[0], y + d[1]
                if x1 < 0 or y1 < 0 or x1 >= m or y1 >= n:
                    continue

                step = x1 * n + y1
                if board[x1][y1] == word[pos] and step not in steps:
                    st.append((pos + 1, x1, y1, steps | set([step])))
        return False


print Solution().exist([
    ['A', 'B', 'C', 'E'],
    ['S', 'F', 'E', 'S'],
    ['A', 'D', 'E', 'E'],
], 'ABCESEEEFS')
