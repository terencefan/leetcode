#!/usr/bin/env python
# -*- coding: utf-8 -*-

# Author: stdrickforce (Tengyuan Fan)
# Email: <stdrickforce@gmail.com> <fantengyuan@baixing.com>

import heapq


class Solution(object):
    def kSmallestPairs(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[List[int]]
        """
        if not nums1 or not nums2 or not k:
            return []

        h = []
        for i in range(min(k, len(nums1))):
            heapq.heappush(h, (nums1[i] + nums2[0], (i, 0)))

        r = []
        while k > 0 and h:
            v, t = heapq.heappop(h)
            i, j = t
            if j + 1 < len(nums2):
                heapq.heappush(h, (nums1[i] + nums2[j + 1], (i, j + 1)))
            r.append([nums1[i], nums2[j]])
            k -= 1
        return r
