#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def rotate(self, a):
        """
        :type matrix: List[List[int]]
        :rtype: void Do not return anything, modify matrix in-place instead.
        """

        m, n = len(a) - 1, len(a)

        def _rotate(i, j):
            a[i][j], a[j][m - i], a[m - i][m - j], a[m - j][i] = \
                a[m - j][i], a[i][j], a[j][m - i], a[m - i][m - j]

        for i in range(n / 2 + n % 2):
            for j in range(n / 2):
                _rotate(i, j)


print Solution().rotate([
    [1,  2,  3,  4,  5],
    [6,  7,  8,  9,  10],
    [11, 12, 13, 14, 15],
    [16, 17, 18, 19, 20],
    [21, 22, 23, 24, 25],
])
