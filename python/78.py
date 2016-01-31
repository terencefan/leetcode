#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def subsets(self, nums):
        nums.sort()
        return self._subsets(nums)

    def _subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        if len(nums) == 0:
            return [[]]

        index, last = 1, nums[-1]
        while index < len(nums):
            if nums[-index - 1] != last:
                break
            index += 1

        res, s = [], self.subsets(nums[:-index])
        for i in range(index + 1):
            for t in s:
                res.append(t + [last] * i)
        return res


print Solution().subsets([1, 2, 2, 3])
