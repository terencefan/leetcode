#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import (
    defaultdict,
)


class Solution(object):

    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """

        chars, missing = defaultdict(int), len(t)
        for c in t:
            chars[c] += 1

        i, p, q = 0, 0, 0
        for j, c in enumerate(s, 1):
            missing -= chars[c] > 0
            chars[c] -= 1

            if missing:
                continue

            while i < j and chars[s[i]] < 0:
                chars[s[i]] += 1
                i += 1

            if not q or j - i < q - p:
                p, q = i, j

        return s[p:q]


if __name__ == '__main__':
    print Solution().minWindow("cabbwefgewcwaefgcf", "cae")
