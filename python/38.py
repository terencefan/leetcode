#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        if n == 1:
            return '1'

        s = self.countAndSay(n - 1)
        x, num = None, 0
        res = ''
        for c in s:
            if c != x:
                res += '' if x is None else '%s%s' % (num, x)
                x, num = c, 1
            else:
                num += 1
        res += '%s%s' % (num, x)
        return res

print Solution().countAndSay(1)
print Solution().countAndSay(2)
print Solution().countAndSay(3)
print Solution().countAndSay(4)
print Solution().countAndSay(5)
print Solution().countAndSay(6)
print Solution().countAndSay(7)
