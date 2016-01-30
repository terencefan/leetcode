#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import deque


class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        length = len(nums)
        res, st = [0] * length, deque([(0, 0)])

        while st:
            pos, step = st.popleft()
            for i in range(min(nums[pos], length - pos - 1), 0, -1):
                if res[pos + i] and res[pos + i] <= step + 1:
                    continue
                if pos + i == length - 1:
                    return step + 1
                res[pos + i] = step + 1
                st.append((pos + i, step + 1))
        return res[length - 1]
