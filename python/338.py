#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        r = [0, 1]
        while len(r) < num + 1:
            r = r + [x + 1 for x in r]
        return r[:num + 1]
