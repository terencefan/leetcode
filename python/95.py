#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from utils import (
    TreeNode,
)


def cache(func):
    s = {}

    def wraps(self, x, y):
        if (x, y) not in s:
            s[(x, y)] = func(self, x, y)
        return s[(x, y)]

    return wraps


class Solution(object):

    def generateTrees(self, n):
        """
        :type n: int
        :rtype: List[TreeNode]
        """
        if n == 0:
            return []
        return list(self.generate(1, n))

    @cache
    def generate(self, l, r):
        if l > r:
            return [None]

        res = []
        for i in range(l, r + 1):
            for left in self.generate(l, i - 1):
                for right in self.generate(i + 1, r):
                    node = TreeNode(i)
                    node.left = left
                    node.right = right
                    res.append(node)
        return res


print Solution().generateTrees(2)
