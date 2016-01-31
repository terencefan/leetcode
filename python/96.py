#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


def cache(func):
    s = {}

    def wraps(self, x, y):
        if (x, y) not in s:
            s[(x, y)] = func(self, x, y)
        return s[(x, y)]

    return wraps


class Solution(object):

    def numTrees(self, n):
        """
        :type n: int
        :rtype: List[TreeNode]
        """
        if n == 0:
            return 0
        return self.num(1, n)

    @cache
    def num(self, l, r):
        if l > r:
            return 1

        res = 0
        for i in range(l, r + 1):
            res += self.num(l, i - 1) * self.num(i + 1, r)
        return res


print Solution().numTrees(3)
