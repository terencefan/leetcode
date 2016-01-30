#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        n, k, res = min(len(a), len(b)), 0, []
        for i in range(1, n + 1):
            k, p = divmod(int(a[-i]) + int(b[-i]) + k, 2)
            res.append(p)

        i, s = 1, max(a[:-n], b[:-n])
        while i < len(s) + 1:
            k, p = divmod(int(s[-i]) + k, 2)
            res.append(p)
            i += 1
        if k:
            res.append(k)
        return ''.join(map(str, reversed(res)))


print Solution().addBinary('100', '110010')
