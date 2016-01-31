#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def combine(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """

        if n < k:
            return []

        result, st = [], [([], 1, k)]
        while st:
            res, s, m = st.pop()
            if m == 0:
                result.append(res)
                continue
            for i in xrange(s, n - m + 2):
                st.append((res + [i], i + 1, m - 1))
        return result


print Solution().combine(4, 2)
