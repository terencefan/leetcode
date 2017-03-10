#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class RBTree:
    pass


class Solution(object):

    def countRangeSum(self, nums, lower, upper):
        """
        :type nums: List[int]
        :type lower: int
        :type upper: int
        :rtype: int
        """

        r, s, v = 0, 0, RBTree()
        for num in nums:
            s += num
            v.add(s)
            r += v.count_between(s - upper - 1, s - lower + 1)
        return r


# Solution().countRangeSum([-2, 5, -1], -2, 2)
# Solution().countRangeSum([0, 0], 0, 0)
Solution().countRangeSum([2147483647, -2147483648, -1, 0], -1, 0)
