#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        """

        if not s:
            return 0

        n = len(s)
        f, s = [0] * (n + 1), map(int, s)
        f[0], f[1] = 1, int(s[0] > 0)
        for i in range(1, len(s)):
            if s[i]:
                f[i + 1] = f[i]
            if s[i - 1] and 0 < s[i - 1] * 10 + s[i] < 27:
                f[i + 1] += f[i - 1]
            if not f[i + 1]:
                return 0
        return f[n]


print Solution().numDecodings('120120')
