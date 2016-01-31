#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    # DP solution.
    # def grayCode(self, n):
    #     """
    #     :type n: int
    #     :rtype: List[int]
    #     """
    #     res = [0]
    #     for i in range(n):
    #         res += [x + pow(2, i) for x in reversed(res)]
    #     return res

    def grayCode(self, n):
        res = []
        for i in range(pow(2, n)):
            res.append(i ^ (i >> 1))
        return res


print Solution().grayCode(3)
