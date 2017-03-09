#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def findDiagonalOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        if not matrix:
            return []
        m, n, r = len(matrix), len(matrix[0]), []
        k = abs(m - n) + 2 * min(m, n) - 1
        for i in range(k):
            s = max(i - n + 1, 0)
            t = min(m - 1, i)
            p = range(s, t + 1) if i % 2 else range(t, s - 1, -1)
            for k in p:
                r.append(matrix[k][i - k])
        return r
