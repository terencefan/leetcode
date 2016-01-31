#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def largestRectangleArea(self, height):
        """
        :type heights: List[int]
        :rtype: int
        """
        st, res = [], 0
        height.append(0)
        for i, h in enumerate(height):
            while st and height[i] < height[st[-1]]:
                h = height[st.pop()]
                w = i - st[-1] - 1 if st else i
                res = max(res, h * w)
            st.append(i)
        return res

    def maximalRectangle(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix:
            return 0
        m, n = len(matrix), len(matrix[0])
        f, res = [0] * n, 0
        for i in range(m):
            for j in range(n):
                f[j] = f[j] + 1 if matrix[i][j] == '1' else 0
            res = max(res, self.largestRectangleArea(f))
        return res


print Solution().maximalRectangle([
    '0010',
    '0111',
    '1110',
    '0100',
])
