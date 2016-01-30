#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def minPathSum(self, a):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m, n = len(a), len(a[0])
        for i in range(m):
            for j in range(n):
                if i + j == 0:
                    continue
                elif i == 0:
                    a[i][j] += a[i][j - 1]
                elif j == 0:
                    a[i][j] += a[i - 1][j]
                else:
                    a[i][j] += min(a[i][j - 1], a[i - 1][j])
        return a[m - 1][n - 1]
