#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

import heapq


class Solution(object):
    def nthSuperUglyNumber(self, n, primes):
        """
        :type n: int
        :type primes: List[int]
        :rtype: int
        """
        r, index = [1], []
        for p in primes:
            heapq.heappush(index, (p, p, 0))

        i = 1
        while i < n:
            x, p, j = heapq.heappop(index)
            if x > r[-1]:
                r.append(x)
                i += 1
            heapq.heappush(index, (p * r[j + 1], p, j + 1))
        return r[n - 1]


x = Solution().nthSuperUglyNumber(25, [2, 3, 5])
print(x)
