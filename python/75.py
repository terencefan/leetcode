#!/usr/bin/env python
# -*- coding: utf-8 -*-

__author__ = "stdrickforce"  # Tengyuan Fan
# Email: <stdrickforce@gmail.com> <tfan@xingin.com>


class Solution(object):

    def sortColors(self, a):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        i, j, k = 0, len(a) - 1, 0
        while k <= j:
            if a[k] == 0 and i != k:
                a[k], a[i] = a[i], a[k]
                i += 1
            elif a[k] == 2 and j != k:
                a[k], a[j] = a[j], a[k]
                j -= 1
            else:
                k += 1


Solution().sortColors([0])
