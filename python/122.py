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

        prices.append(0)

        s, res = -1, 0
        for i in range(len(prices) - 1):
            if s < 0:
                if prices[i + 1] > prices[i]:
                    s = prices[i]
            elif prices[i + 1] <= prices[i]:
                res += prices[i] - s
                s = -1
        return res


print Solution().maxProfit([1, 6, 2, 4, 9])
