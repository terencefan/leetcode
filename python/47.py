#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result, length = [], len(nums)

        def _change(pos):
            if pos == length - 1:
                result.append(list(nums))
            for i in range(pos, length):
                if nums[i] in nums[pos:i] and i != pos:
                    continue
                nums[i], nums[pos] = nums[pos], nums[i]
                _change(pos + 1)
                nums[pos], nums[i] = nums[i], nums[pos]

        _change(0)
        return result


print Solution().permute([3, 3, 0, 0, 2, 3, 2])
