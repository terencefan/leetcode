#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        if n == 0:
            return 1
        elif n == 1:
            return x
        elif n < 0:
            return 1.0 / self.myPow(x, -n)

        res = self.myPow(x, n / 2)
        if n % 2:
            return res * res * x
        else:
            return res * res


print Solution().myPow(0.00001, 2147483547)
