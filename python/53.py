#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        length = len(nums)
        tmp = [i for i in nums]
        for i in range(1, length):
            tmp[i] = max(tmp[i], tmp[i - 1] + nums[i])
        return max(tmp)


print Solution().maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])
