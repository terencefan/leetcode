#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """

        s, t = 0, len(nums) - 1
        print nums
        while s < t:
            mid = (s + t) / 2
            print s, t, mid
            if nums[mid] == target:
                return mid

            if nums[s] <= nums[mid]:
                if nums[s] <= target < nums[mid]:
                    t = mid - 1
                else:
                    s = mid + 1
            else:
                if nums[t] >= target > nums[mid]:
                    s = mid + 1
                else:
                    t = mid - 1
        return s if nums[s] == target else -1


if __name__ == '__main__':
    print Solution().search([5, 6, 7, -1, 0, 1, 2, 3, 4], -1)
