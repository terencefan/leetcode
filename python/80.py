#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i, j, k = 0, 0, 0
        while j < len(nums):
            nums[i] = nums[j]
            if i > 0 and nums[i] == nums[i - 1]:
                if k < 1:
                    k += 1
                    i += 1
            else:
                k = 0
                i += 1
            j += 1

        for i in range(i, j):
            nums.pop()


print Solution().removeDuplicates([1, 1, 1])
