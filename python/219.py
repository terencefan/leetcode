#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>


class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        m = [(x, i) for i, x in enumerate(nums)]
        m.sort()
        pre = (None, 0)
        for x, i in m:
            if x == pre[0] and i - pre[1] <= k:
                return True
            pre = (x, i)
        return False
