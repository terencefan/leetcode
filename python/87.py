#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import defaultdict


class Solution(object):
    def isScramble(self, s1, s2):
        """
        :type s1: str
        :type s2: str
        :rtype: bool
        """

        if s1 == s2:
            return True

        count = defaultdict(int)
        for c1, c2 in zip(s1, s2):
            count[c1] += 1
            count[c2] -= 1
        for k, v in count.items():
            if v != 0:
                return False

        for i in range(1, len(s1)):
            if self.isScramble(s1[:i], s2[:i]) and \
                    self.isScramble(s1[i:], s2[i:]):
                return True
            if self.isScramble(s1[:i], s2[-i:]) and \
                    self.isScramble(s1[i:], s2[:-i]):
                return True
        return False


print Solution().isScramble('bca', 'abd')
