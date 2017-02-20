#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def candy(self, ratings):
        """
        :type ratings: List[int]
        :rtype: int
        """
        n = len(ratings)
        a = [1] * n

        for i in range(1, n):
            if ratings[i] > ratings[i - 1]:
                a[i] = a[i - 1] + 1
        for i in range(n - 1, 0, -1):
            if ratings[i] < ratings[i - 1]:
                a[i - 1] = max(a[i] + 1, a[i - 1])

        return sum(a)


if __name__ == '__main__':
    Solution().candy([1, 3, 4, 3, 2, 1])
