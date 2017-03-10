#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def canPartition(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        s = sum(nums)
        if s % 2:
            return False
        s = s / 2

        x = [0] * (s + 1)
        for i, num in enumerate(nums):
            for j in range(s, num - 1, -1):
                x[j] = max(x[j], x[j - num] + num)
                if x[j] == s:
                    return True
        return False
