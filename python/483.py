#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

import math


class Solution(object):

    def smallestGoodBase(self, n):
        """
        :type n: str
        :rtype: str
        """

        def calculate(a, n):
            res = 1
            for i in range(n - 1):
                res *= a
                res += 1
            return res

        n = int(n)
        m = int(math.ceil(math.log(n, 2)))
        a = [n] * (m + 1)
        k = n - 1
        for i in range(3, m + 1):
            l, r = 2, a[i - 1]
            while l <= r:
                p = int((l + r) / 2)
                q = calculate(p, i)
                if q == n:
                    k = p
                    break
                elif q > n:
                    r = p - 1
                else:
                    l = p + 1
            a[i] = p
        print(a)
        return str(k)


Solution().smallestGoodBase("2251799813685247")
