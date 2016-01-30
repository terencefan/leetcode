#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def solveNQueens(self, n):
        """
        :type n: int
        :rtype: List[List[str]]
        """

        def _valid(y, x, qs):
            x1 = x2 = x
            while y > 0:
                y, x1, x2 = y - 1, x1 - 1, x2 + 1
                if (y, x1) in qs or (y, x2) in qs:
                    return False
            return True

        res, st = 0, [(0, set([]))]
        while st:
            num, qs = st.pop()

            if num == n:
                res += 1
                continue

            columns = set(range(n)) - set(x[1] for x in qs)
            for i in columns:
                if not _valid(num, i, qs):
                    continue
                t = qs | set([(num, i)])
                st.append((num + 1, t))

        return res


print Solution().solveNQueens(11)
