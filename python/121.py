#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """

        if len(prices) < 2:
            return 0

        s, res = prices[0], 0
        for price in prices[1:]:
            res = max(res, price - s)
            s = min(price, s)
        return res

print Solution().maxProfit([2, 1, 3, 5, 2])
