#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return reduce(lambda x, y: x ^ y, nums)

    def singleNumber2(self, nums):
        x = reduce(lambda x, y: x ^ y, nums)

        # 找到右起第一个1，也就是a和b不同的那一位
        t = x & (-x)

        # 根据这一位将数组分成两部分，分别求唯一的数
        a, b = 0, 0
        for num in nums:
            if num & t:
                a = a ^ num
            else:
                b = b ^ num
        return a, b

    def singleNumber3(self, nums):
        xor = lambda x, y: x ^ y
        f = lambda x: x & (-x)

        x = reduce(xor, nums)
        y = reduce(xor, [f(x ^ num) for num in nums])
        y = f(y)

        a = 0
        for num in nums:
            if num & y:
                a = a ^ num

        t = f(x ^ a)
        b, c = 0, 0
        for num in nums:
            if num == a:
                continue
            elif num & t:
                b = b ^ num
            else:
                c = c ^ num

        return a, b, c


print Solution().singleNumber3([1, 1, 2, 2, 3, 7, 15])
