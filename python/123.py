#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def maxProfit(self, prices, k=2):
        """
        :type prices: List[int]
        :rtype: int
        """

        if len(prices) < 2:
            return 0

        n = len(prices)
        f = [0] * n

        s = prices[0]
        for i in range(1, n):
            s = min(prices[i], s)
            f[i] = max(f[i - 1], prices[i] - s)

        res, max_price = f[-1], 0
        for j in reversed(range(1, n)):
            max_price = max(max_price, prices[j])
            res = max(res, f[j - 1] + max_price - prices[j])
        return res


print Solution().maxProfit([3, 2, 6, 5, 0, 3])
