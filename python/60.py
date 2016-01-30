#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """

        m = [1]
        for i in range(2, n):
            m.append(m[-1] * i)

        res, k = [], k - 1
        for i in reversed(m):
            p, k = divmod(k, i)
            res.append(p)
        if len(res) != n:
            res.append(0)

        k = range(n)
        return ''.join(map(str, [k.pop(i) + 1 for i in res]))

print Solution().getPermutation(1, 1)
