#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class NumArray(object):

    def __init__(self, nums):
        """
        :type nums: List[int]
        """

    def sumRange(self, i, j):
        """
        :type i: int
        :type j: int
        :rtype: int
        """
        x = 1
        while x < i:
            x <<= 1


if __name__ == '__main__':
    a = NumArray([-2, 0, 3, -5, 2, -1])
    print(a.sumRange(0, 2))
    print(a.sumRange(2, 5))
    print(a.sumRange(0, 5))
