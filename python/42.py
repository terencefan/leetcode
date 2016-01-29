#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """

        length = len(height)
        if not length:
            return 0

        top = (0, height[0])
        for i, h in enumerate(height):
            if h > top[0]:
                top = (h, i)

        def _stack(iterator):
            res = [top]
            for i in iterator:
                while height[i] > res[-1][0]:
                    res.pop()
                res.append((height[i], i))
            return res

        def _fill(st):
            res = 0
            for i in range(len(st) - 1):
                res += st[i + 1][0] * abs(st[i + 1][1] - st[i][1])
            return res

        liter = reversed(range(0, top[1]))
        riter = range(top[1] + 1, length)

        lst = _stack(liter)
        rst = _stack(riter)

        return _fill(lst) + _fill(rst) - sum(height) + top[0]


print Solution().trap([0, 1, 0, 3, 1, 0, 1, 3, 2, 1, 2, 1])
