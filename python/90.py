#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Created At: Thu Sep 10 13:52:12 2015
# Updated At: Thu Sep 10 14:12:14 2015

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result, nums = [], sorted(nums)
        for index, i in enumerate(nums):
            if i == nums[index - 1] and index > 0:
                continue
            for ans in self.subsetsWithDup(nums[index + 1:]):
                result.append([i] + ans)
        result.append([])
        return result


if __name__ == '__main__':
    print Solution().subsetsWithDup([4, 1, 0])
