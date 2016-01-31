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

print Solution().largestRectangleArea([2, 1, 5, 6, 2, 3])
