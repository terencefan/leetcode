#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>

from collections import deque


class Solution(object):

    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """

        length, reach = len(nums), 0

        i = 0
        while i <= reach and i < length - 1:
            reach = max(reach, i + nums[i])
            i += 1
        return reach >= length - 1


print Solution().canJump([3, 2, 1, 1, 4])
