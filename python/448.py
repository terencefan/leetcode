#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

nums = [4, 3, 2, 7, 7, 2, 3, 1]


class Solution(object):

    def findDisappearedNumbers1(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        for i in range(n):
            k = nums[i] - 1
            while i != k and k >= 0:
                if nums[i] == nums[k]:
                    nums[i] = 0
                    break
                nums[i], nums[k] = nums[k], nums[i]
                k = nums[i] - 1
        return [i + 1 for i in range(n) if nums[i] == 0]

    def findDisappearedNumbers(self, nums):
        n = len(nums)
        for i in range(n):
            k = nums[i] - 1
            while k >= 0:
                t = k
                k = nums[k] - 1
                nums[t] = 0
        return [i + 1 for i in range(n) if nums[i] != 0]

print(Solution().findDisappearedNumbers(nums))
