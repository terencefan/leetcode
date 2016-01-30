#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def generateMatrix(self, n):
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

        i, j, k = 0, 0, 0
        matrix = [[0] * n for x in range(n)]

        def _valid(i, j, d):
            i += d[0]
            j += d[1]
            if 0 <= i < n and 0 <= j < n:
                if matrix[i][j]:
                    return False
                return True
            return False

        for p in range(n * n):
            matrix[i][j] = p + 1
            if not _valid(i, j, directions[k]):
                k = (k + 1) % 4
                if not _valid(i, j, directions[k]):
                    break
            i, j = i + directions[k][0], j + directions[k][1]

        return matrix


print Solution().generateMatrix(3)
