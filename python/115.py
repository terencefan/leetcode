#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def numDistinct(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        m, n = len(s), len(t)
        f = [1] + [0] * n
        for i in range(m):
            for j in reversed(range(n)):
                if s[i] == t[j]:
                    f[j + 1] += f[j]
        return f[n]


if __name__ == '__main__':
    print Solution().numDistinct('rabbbit', 'rabbit')
