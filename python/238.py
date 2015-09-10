#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Created At: Thu Sep 10 15:06:04 2015
# Updated At: Thu Sep 10 15:21:06 2015

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        product, noz = 1, 0
        for num in nums:
            noz += 0 if num else 1
            product = product * (num or 1)

        result = []
        for num in nums:
            if noz > 1:
                result.append(0)
            elif noz == 1:
                result.append(0 if num else product)
            else:
                result.append(product / num)
        return result


if __name__ == '__main__':
    print Solution().productExceptSelf([0, 1])
    print Solution().productExceptSelf([0, 0])
    print Solution().productExceptSelf([1, 2, 3, 4])
