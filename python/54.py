#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """

        directions = [
            (0, 1),
            (1, 0),
            (0, -1),
            (-1, 0),
        ]

        if not matrix:
            return []

        i, j, k, m, n = 0, 0, 0, len(matrix), len(matrix[0])

        def _valid(i, j, d):
            i += d[0]
            j += d[1]
            if 0 <= i < m and 0 <= j < n:
                if matrix[i][j] is None:
                    return False
                return True
            return False

        res = []
        while True:
            res.append(matrix[i][j])
            matrix[i][j] = None
            if not _valid(i, j, directions[k]):
                k = (k + 1) % 4
                if not _valid(i, j, directions[k]):
                    break
            i, j = i + directions[k][0], j + directions[k][1]

        return res


print Solution().spiralOrder([
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
])
