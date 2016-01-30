#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """

        m, n = len(matrix), len(matrix[0])

        def _get(x):
            p, q = divmod(x, n)
            return matrix[p][q]

        i, j = 0, m * n - 1
        while i <= j:
            mid = (i + j) / 2
            val = _get(mid)
            if val == target:
                return True
            elif target < val:
                j = mid - 1
            else:
                i = mid + 1
        return False


print Solution().searchMatrix([
    [1, 3, 5, 7],
    [10, 11, 16, 20],
    [23, 30, 34, 50],
], 3)
