#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def minDistance(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: int
        """
        m, n = len(word1), len(word2)
        f = [[0] * (n + 1) for i in xrange(m + 1)]
        for i in xrange(m + 1):
            f[i][0] = i
        for j in xrange(n + 1):
            f[0][j] = j

        for i in xrange(1, m + 1):
            for j in xrange(1, n + 1):
                f[i][j] = min(
                    f[i - 1][j] + 1,
                    f[i][j - 1] + 1,
                    f[i - 1][j - 1] + (word1[i - 1] != word2[j - 1])
                )
        return f[m][n]


print Solution().minDistance('ab', 'accb')
