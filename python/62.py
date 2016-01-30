#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def uniquePathsWithObstacles(self, a):
        m, n = len(a), len(a[0])

        if a[0][0] == 1:
            return 0
        a[0][0] = 1

        flag = 1
        for i in range(1, m):
            if a[i][0] == 1:
                flag = 0
            a[i][0] = flag

        flag = 1
        for i in range(1, n):
            if a[0][i] == 1:
                flag = 0
            a[0][i] = flag

        for i in range(1, m):
            for j in range(1, n):
                if a[i][j] == 1:
                    a[i][j] = 0
                    continue
                a[i][j] = a[i - 1][j] + a[i][j - 1]
        return a[m - 1][n - 1]


print Solution().uniquePathsWithObstacles([
    [0, 0, 0],
    [0, 1, 0],
    [0, 0, 0],
])
