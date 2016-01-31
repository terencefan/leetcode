#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    # dp solution.
    def isInterleave(self, s1, s2, s3):
        if len(s3) != len(s1) + len(s2):
            return False
        elif not s1:
            return s2 == s3
        elif not s2:
            return s1 == s3
        m, n = len(s1), len(s2)

        f = [[False] * (n + 1) for i in range(m + 1)]
        f[0][0] = True
        for i in range(m):
            f[i + 1][0] = f[i][0] and s1[i] == s3[i]
        for j in range(n):
            f[0][j + 1] = f[0][j] and s2[j] == s3[j]

        for i in range(m):
            for j in range(n):
                res = False
                if s1[i] == s3[i + j + 1]:
                    res = res or f[i][j + 1]
                if s2[j] == s3[i + j + 1]:
                    res = res or f[i + 1][j]
                f[i + 1][j + 1] = res
        return f[m][n]

    # recursive solution, Timeout
    def isInterleave1(self, s1, s2, s3):
        if len(s1) + len(s2) != len(s3):
            return False
        elif not s1:
            return s2 == s3
        elif not s2:
            return s1 == s3

        res = False
        if s1[0] == s3[0]:
            res = res or self.isInterleave(s1[1:], s2, s3[1:])
        if s2[0] == s3[0]:
            res = res or self.isInterleave(s1, s2[1:], s3[1:])
        return res


if __name__ == '__main__':
    print Solution().isInterleave('aabcc', 'dbbca', 'aadbbcbcac')
