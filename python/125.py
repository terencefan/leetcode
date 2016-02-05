#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """

        def isC(c):
            return 'a' <= c <= 'z' or 'A' <= c <= 'Z' or '0' <= c <= '9'

        a = [c.lower() for c in s if isC(c)]
        i, j = 0, len(a) - 1,
        print a

        while i < j:
            if a[i] != a[j]:
                return False
            i, j = i + 1, j - 1
        return True


print Solution().isPalindrome('`l;`` 1o1 ??;l`')
