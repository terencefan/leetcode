#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        # free, have, cool
        a = (0, -1000000, -1000000)
        for i, p in enumerate(prices):
            a = (
                max(a[0], a[2]),
                max(a[0] - p, a[1]),
                a[1] + p
            )
        return max(a)


r = Solution().maxProfit([1, 2, 3, 0, 2])
print(r)
