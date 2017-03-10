#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

MIND = 123123


def guess(x):
    if MIND > x:
        return 1
    elif MIND < x:
        return -1
    return 0


class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        m = 1
        while guess(m) > 0:
            m <<= 1

        c = m
        while m > 0:
            m >>= 1
            r = guess(c)
            if r > 0:
                c += m
            elif r < 0:
                c -= m
            else:
                return c
        return 0


c = Solution().guessNumber(2)
print(c)
