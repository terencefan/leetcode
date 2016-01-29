#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        length = len(nums)
        for i in range(length):
            num = nums[i]
            while num > 0 and num <= length and nums[num - 1] != nums[i]:
                nums[num - 1], nums[i] = nums[i], nums[num - 1]
                num = nums[i]

        for i in range(length):
            if nums[i] != i + 1:
                return i + 1
        return length + 1

print Solution().firstMissingPositive([1, 1])
