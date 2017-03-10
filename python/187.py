#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):

    def findRepeatedDnaSequences(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        z = {}
        for i in range(len(s) - 9):
            l = s[i:i + 10]
            z[l] = z.setdefault(l, 0) + 1
        return [x for x, num in z.items() if num > 1]
